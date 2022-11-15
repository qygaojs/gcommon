/**
 * 生产环境尽量不要使用DEBUG级别
 * 因为debug级别会将协程号信息打印出来
 * 打印协程号的性能不高
 * 在调试时，根据协程号查询一个请求还是比较有用的
 * 所以除非生产环境并发确实很低，否则不要使用debug打印日志
 */
package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/qygaojs/gcommon/core"
	"github.com/qygaojs/gcommon/file"
)

// 日志级别
type LEVEL int8

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	NOTICE
	WARNING
	ERROR
	FATAL
	OFF
)

var LEVEL_COLOR = map[LEVEL]string{
	DEBUG:   "%c[2;39m%s%c[0m",
	INFO:    "%c[0;36m%s%c[0m",
	NOTICE:  "%c[0;34m%s%c[0m",
	WARNING: "%c[0;33m%s%c[0m",
	ERROR:   "%c[0;35m%s%c[0m",
	FATAL:   "%c[1;31m%s%c[0m",
	OFF:     "",
}

// 日志大小的单位
type UNIT int64

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

// 日志切分方式
type ROTATETYPE int

const (
	ROTATE_OFF  ROTATETYPE = iota
	ROTATE_DAY             // 按天切分
	ROTATE_SIZE            // 按大小切分
)

// 按日期切分日志后缀格式
const DATEFORMAT = "2006-01-02"

type _LOGITEM struct {
	filename    string
	level       LEVEL
	fp          *os.File
	lg          *log.Logger
	lastCheckTs int64
	isColor     bool
}

type _LogItemSlice []*_LOGITEM

func (l _LogItemSlice) Len() int {
	return len(l)
}

func (l _LogItemSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l _LogItemSlice) Less(i, j int) bool {
	return l[i].level > l[j].level // 从大到小排序
}

func (l _LogItemSlice) String() {
	for _, item := range l {
		fmt.Printf("filename:%s level:%d\n", item.filename, item.level)
	}
}

type _LOGFILE struct {
	mu *sync.RWMutex

	items []*_LOGITEM // 日志文件

	isPropagate  bool       // 高级别日志是否写到低级别中
	rotateType   ROTATETYPE // 日志切分方式
	maxFileSize  UNIT       // 每个日志文件最大的大小
	maxFileCount int64      // 保留的最大文件数
}

var logObj *_LOGFILE = nil

func GetLogger() *_LOGFILE {
	return logObj
}

type LogOption struct {
	f func(*_LOGFILE)
}

func IsPropagate(isPropagate bool) LogOption {
	return LogOption{func(lg *_LOGFILE) {
		lg.isPropagate = isPropagate
	}}
}

func RotateType(rotateType ROTATETYPE) LogOption {
	return LogOption{func(lg *_LOGFILE) {
		lg.rotateType = rotateType
	}}
}

func RotateTypeStr(rotateType string) LogOption {
	return LogOption{func(lg *_LOGFILE) {
		if rotateType == "day" {
			lg.rotateType = ROTATE_DAY
		} else if rotateType == "size" {
			lg.rotateType = ROTATE_SIZE
		} else {
			fmt.Println("rotateType is invalid:", rotateType)
			os.Exit(-1)
		}
	}}
}

func MaxFileSize(maxFileSize UNIT) LogOption {
	return LogOption{func(lg *_LOGFILE) {
		lg.maxFileSize = maxFileSize
	}}
}

func MaxFileCount(maxFileCount int64) LogOption {
	return LogOption{func(lg *_LOGFILE) {
		lg.maxFileCount = maxFileCount
	}}
}

/**
 * config: 日志文件配置, 格式:{"日志级别": "日志文件名"},
 *         例: {"DEBUG": "/export/log/test.log", "WARN": "/export/log/test_warn.log"}
 *         日志最多只支持2级
 * option: 日志配置选项, 格式:{"配置选项名": 值}
 *         配置选项:
 *         ---------------------------------------------------------------
 *         | rotateType   | 日志切分方式，day按天切分 size按文件大小切分 |
 *         | maxFileSize  | 日志文件大小，单位: B                        |
 *         | maxFileCount | 日志文件数                                   |
 *         | isPropagate  | 是否将高级别的日志写到低级别中               |
 *         ---------------------------------------------------------------
 */
func Install(config map[string]string, logOptions map[string]interface{}) error {
	options := []LogOption{}
	if isPropagate, ok := logOptions["isPropagate"]; ok {
		if propagate, err := core.Bool(isPropagate); err != nil {
			fmt.Println("isPropagate is invalid, ignore:", isPropagate)
		} else {
			options = append(options, IsPropagate(propagate))
		}
	}
	if rotateType, ok := logOptions["rotateType"]; ok {
		if rotate, err := core.String(rotateType); err != nil {
			fmt.Println("rotateType is invalid, ignore:", rotateType)
		} else {
			options = append(options, RotateTypeStr(rotate))
		}
	}
	if maxFileSize, ok := logOptions["maxFileSize"]; ok {
		if fileSize, err := core.Int(maxFileSize); err != nil {
			fmt.Println("maxFileSize is invalid, ignore:", maxFileSize)
		} else {
			options = append(options, MaxFileSize(UNIT(fileSize)))
		}
	}
	if maxFileCount, ok := logOptions["maxFileCount"]; ok {
		if fileCount, err := core.Int(maxFileCount); err != nil {
			fmt.Println("maxFileCount is invalid, ignore:", maxFileCount)
		} else {
			options = append(options, MaxFileCount(fileCount))
		}
	}
	return New(config, options...)
}

func New(config map[string]string, options ...LogOption) error {
	if logObj != nil {
		fmt.Println("log object is exist")
		return nil
	}
	logObj = new(_LOGFILE)
	logObj.items = getFileLevel(config)
	if len(logObj.items) == 0 {
		return fmt.Errorf("log config is invalid")
	}
	for _, option := range options {
		option.f(logObj)
	}
	logObj.mu = new(sync.RWMutex)
	return nil
}

func string2level(s string) LEVEL {
	s = strings.ToUpper(s)
	if s == "DEBUG" {
		return DEBUG
	} else if s == "INFO" {
		return INFO
	} else if s == "NOTICE" || s == "NOTE" {
		return NOTICE
	} else if s == "WARNING" || s == "WARN" {
		return WARNING
	} else if s == "ERROR" {
		return ERROR
	} else if s == "FATAL" {
		return FATAL
	}
	return OFF
}

func getFileLevel(config map[string]string) []*_LOGITEM {
	logItem := make([]*_LOGITEM, 0, len(config))
	for k, v := range config {
		level := string2level(k)
		item := getLogItem(v, level)
		logItem = append(logItem, item)
	}
	if len(logItem) == 0 {
		fmt.Println("log config is invalid")
		os.Exit(-1)
	}
	sort.Sort(_LogItemSlice(logItem))
	return logItem
}

func getLogItem(filename string, level LEVEL) *_LOGITEM {
	item := new(_LOGITEM)
	item.filename = filename
	item.level = level
	item.lastCheckTs = 0

	if level == OFF {
		return item
	}
	if filename == "stdout" {
		item.fp = os.Stdout
		item.isColor = true
	} else {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			fmt.Println("create log dir failed:", err)
			os.Exit(-1)
		}
		fp, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		item.fp = fp
		if err != nil {
			fmt.Println("open log file failed:", err)
			os.Exit(-1)
		}
		item.isColor = false
	}
	//item.lg = log.New(item.fp, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	item.lg = log.New(item.fp, "", 0)
	return item
}

func (this *_LOGITEM) flush() error {
	if this.fp != nil {
		return this.fp.Sync()
	}
	return nil
}

func (this *_LOGITEM) Close() error {
	if this.fp != nil {
		this.fp.Sync()
		this.fp.Close()
		this.fp = nil
	}
	return nil
}

func (this *_LOGITEM) rotateDay(fileCount int64) bool {
	currTs := time.Now()
	if this.lastCheckTs <= 0 {
		// 首次运行，不切分
		this.lastCheckTs = currTs.Unix()
		return true
	}
	lastCheckTm := time.Unix(this.lastCheckTs, 0)
	if lastCheckTm.Year() == currTs.Year() && lastCheckTm.Month() == currTs.Month() && lastCheckTm.Day() == currTs.Day() {
		// 没有换天，不切换
		return true
	}
	logObj.mu.Lock()
	defer logObj.mu.Unlock()
	// 两次检查锁，以处理并发问题
	lastCheckTm = time.Unix(this.lastCheckTs, 0)
	if lastCheckTm.Year() == currTs.Year() && lastCheckTm.Month() == currTs.Month() && lastCheckTm.Day() == currTs.Day() {
		// 没有换天，不切换
		return true
	}
	this.flush()
	logDir, shortName := filepath.Split(this.filename)
	//logDir := this.filename[:strings.LastIndex(this.filename, "/")]
	//shortName := this.filename[strings.LastIndex(this.filename, "/")+1:]
	rd, err := ioutil.ReadDir(logDir)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] rotate log failed:%v", err))
		return false
	}
	oldLogs := []string{}
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}
		fn := fi.Name()
		if strings.Index(fn, shortName) != 0 {
			// 不是切分的文件
			continue
		}
		fnFields := strings.Split(fn, ".")
		fnSuffix := fnFields[len(fnFields)-1]
		suffixFields := strings.Split(fnSuffix, "-")
		if len(suffixFields) == 3 && len(suffixFields[0]) == 4 && len(suffixFields[1]) == 2 && len(suffixFields[2]) == 2 {
			oldLogs = append(oldLogs, fn)
		}
	}
	moreLogNum := int64(len(oldLogs)) - fileCount
	if moreLogNum >= 0 && fileCount > 0 {
		sort.Sort(sort.StringSlice(oldLogs))
		for i := int64(0); i <= moreLogNum; i++ {
			delFn := filepath.Join(logDir, oldLogs[i])
			err := os.Remove(delFn)
			if err != nil {
				this.lg.Print(fmt.Sprintf("[ERROR] delete old log failed:%v", err))
				return false
			}
		}
	}
	newFn := this.filename + "." + currTs.Format("2006-01-02")
	err = os.Rename(this.filename, newFn)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] delete old log failed:%v", err))
		return false
	}
	this.fp.Close()
	this.fp, err = os.OpenFile(this.filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] open log failed:%v", err))
		os.Exit(-1)
	}
	this.lg = log.New(this.fp, "", 0)
	this.lastCheckTs = currTs.Unix() // 保存最新一次的日志切分时间
	return true
}

func (this *_LOGITEM) rotateSize(fileCount int64, maxFileSize UNIT) bool {
	size, err := file.Size(this.filename)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] get log size failed:%v", err))
		return false
	}
	if UNIT(size) < maxFileSize {
		return true
	}
	logObj.mu.Lock()
	defer logObj.mu.Unlock()
	// 两次检查锁，以处理并发问题
	size, err = file.Size(this.filename)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] get log size failed:%v", err))
		return false
	}
	if UNIT(size) < maxFileSize {
		return true
	}
	this.flush()
	logDir, shortName := filepath.Split(this.filename)
	rd, err := ioutil.ReadDir(logDir)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] rotate log failed:%v", err))
		return false
	}
	oldLogNums := []int{}
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}
		fn := fi.Name()
		if strings.Index(fn, shortName) != 0 {
			// 不是切分的文件
			continue
		}
		fnFields := strings.Split(fn, ".")
		fnSuffix := fnFields[len(fnFields)-1]
		num, err := strconv.Atoi(fnSuffix)
		if err == nil {
			oldLogNums = append(oldLogNums, num)
		}
	}
	newNums := 1 // 默认是1
	if len(oldLogNums) > 0 {
		if fileCount > 0 {
			sort.Sort(sort.IntSlice(oldLogNums))
			moreLogNum := int64(len(oldLogNums)) - fileCount
			for i := int64(0); i <= moreLogNum; i++ {
				// 可能之前会有多少历史日志，全删掉
				delFn := this.filename + "." + strconv.Itoa(oldLogNums[i])
				err := os.Remove(delFn)
				if err != nil {
					this.lg.Print(fmt.Sprintf("[ERROR] delete old log failed:%v", err))
					return false
				}
			}
		}
		newNums = oldLogNums[len(oldLogNums)-1] + 1
	}
	newFn := this.filename + "." + strconv.Itoa(newNums)
	err = os.Rename(this.filename, newFn)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] delete old log failed:%v", err))
		return false
	}
	this.fp.Close()
	this.fp, err = os.OpenFile(this.filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		this.lg.Print(fmt.Sprintf("[ERROR] open log failed:%v", err))
		os.Exit(-1)
	}
	this.lg = log.New(this.fp, "", 0)
	return true
}

func (this *_LOGFILE) Close() {
	for _, item := range this.items {
		item.Close()
	}
	return
}

func (this *_LOGFILE) write(item *_LOGITEM, level LEVEL, data string) bool {
	if item.isColor {
		color_format, ok := LEVEL_COLOR[level]
		if !ok {
			item.lg.Print(data)
		}
		//data = fmt.Sprintf(color_format, 0x1B, data, 0x1B)
		data = fmt.Sprintf(color_format, 0x1B, data, 0x1B)
		item.lg.Print(data)
	} else {
		item.lg.Print(data)
	}
	if item.filename == "stdout" {
		// 标准输出，不切日志
		return true
	}
	if this.rotateType == ROTATE_DAY {
		item.rotateDay(this.maxFileCount)
	} else if this.rotateType == ROTATE_SIZE {
		item.rotateSize(this.maxFileCount, this.maxFileSize)
	}
	return true
}

func catchError() {
	if err := recover(); err != nil {
		log.Print(fmt.Sprintf("[ERROR] print log exception: [[%s]]", err))
	}
}

func caller() (string, int) {
	for stackNum := 3; stackNum < 6; stackNum++ {
		// 默认调用层级是3层，有可能会更多
		file, line := core.Caller(stackNum)
		if file == "logger.go" {
			continue
		}
		if file == "" {
			break // 获取失败
		}
		return file, line
	}
	return "", 0
}

func dataFormat(levelStr, format string, v ...interface{}) string {
	file, line := caller()
	currTm := time.Now()
	strDate := fmt.Sprintf("%s,%03d", currTm.Format("2006-01-02 15:04:05"), currTm.Nanosecond()/1e6)
	info := fmt.Sprintf(format, v...)
	var data string
	var level LEVEL = OFF
	if len(logObj.items) > 0 {
		// 获取配置的日志最低的级别
		level = logObj.items[len(logObj.items)-1].level
	}
	if level <= DEBUG {
		data = fmt.Sprintf("%s %d.%d %s:%d [%s] %s", strDate, os.Getpid(), core.GoroutineID(), file, line, levelStr, info)
	} else {
		data = fmt.Sprintf("%s %s:%d [%s] %s", strDate, file, line, levelStr, info)
	}
	return data
}

func Debug(format string, v ...interface{}) {
	defer catchError()
	if logObj == nil {
		return
	}
	data := dataFormat("DEBUG", format, v...)
	for _, item := range logObj.items {
		if item.lg != nil && item.level <= DEBUG {
			logObj.write(item, DEBUG, data)
			if !logObj.isPropagate {
				// 不复制日志，则只将日志写到允许写入的最高级别中
				return
			}
		}
	}
}

func Info(format string, v ...interface{}) {
	defer catchError()
	if logObj == nil {
		return
	}
	data := dataFormat("INFO", format, v...)
	for _, item := range logObj.items {
		if item.lg != nil && item.level <= INFO {
			logObj.write(item, INFO, data)
			if !logObj.isPropagate {
				// 不复制日志，则只将日志写到允许写入的最高级别中
				return
			}
		}
	}
}

func Notice(format string, v ...interface{}) {
	defer catchError()
	if logObj == nil {
		return
	}
	data := dataFormat("NOTICE", format, v...)
	for _, item := range logObj.items {
		if item.lg != nil && item.level <= NOTICE {
			logObj.write(item, NOTICE, data)
			if !logObj.isPropagate {
				// 不复制日志，则只将日志写到允许写入的最高级别中
				return
			}
		}
	}
}

func Note(format string, v ...interface{}) {
	Notice(format, v...)
}

func Warning(format string, v ...interface{}) {
	defer catchError()
	if logObj == nil {
		return
	}
	data := dataFormat("WARNING", format, v...)
	for _, item := range logObj.items {
		if item.lg != nil && item.level <= WARNING {
			logObj.write(item, WARNING, data)
			if !logObj.isPropagate {
				// 不复制日志，则只将日志写到允许写入的最高级别中
				return
			}
		}
	}
}

func Warn(format string, v ...interface{}) {
	Warning(format, v...)
}

func Error(format string, v ...interface{}) {
	defer catchError()
	if logObj == nil {
		return
	}
	data := dataFormat("ERROR", format, v...)
	for _, item := range logObj.items {
		if item.lg != nil && item.level <= ERROR {
			logObj.write(item, ERROR, data)
			if !logObj.isPropagate {
				// 不复制日志，则只将日志写到允许写入的最高级别中
				return
			}
		}
	}
}

func Fatal(format string, v ...interface{}) {
	defer catchError()
	if logObj == nil {
		return
	}
	data := dataFormat("FATAL", format, v...)
	for _, item := range logObj.items {
		if item.lg != nil && item.level <= FATAL {
			logObj.write(item, FATAL, data)
			if !logObj.isPropagate {
				// 不复制日志，则只将日志写到允许写入的最高级别中
				os.Exit(-1)
			}
		}
	}
	os.Exit(-1)
}

func Close() {
	if logObj == nil {
		return
	}
	logObj.Close()
	logObj = nil
}
