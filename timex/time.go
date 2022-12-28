package timex

import (
	"math/rand"
	"time"

	log "github.com/qygaojs/gcommon/logger"
)

/**
 * 随机休息秒
 * second: 休息的最大秒数
 */
func RSleep(second int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tmRand := r.Intn(second)
	log.Debug("sleep %d second", tmRand)
	time.Sleep(time.Duration(tmRand) * time.Second)
	return
}

/**
 * 随机休息毫秒
 * msecond: 休息的最大毫秒数
 */
func RMSleep(msecond int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tmRand := r.Intn(msecond)
	log.Debug("sleep %d millisecond", tmRand)
	time.Sleep(time.Duration(tmRand) * time.Millisecond)
	return
}

/**
 * 获取中国时区的时间
 */
func GetChTmFormat(timeFormat string) string {
	chinaZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Info("get china zone failed. use current time")
		return time.Now().Format(timeFormat)
	}
	unixTm := time.Unix(time.Now().Unix(), 0) //通过unix标准时间的秒，纳秒设置时间
	chinaTm := unixTm.In(chinaZone).Format(timeFormat)
	return chinaTm
}

/**
 * 获取中国时区的时区
 */
func GetChTm() string {
	timeFormat := "2006-01-02 15:04:05"
	return GetChTmFormat(timeFormat)
}

/**
 * 获取当前日期
 */
func Date() (time.Time, error) {
	dateFormat := "2006-01-02"
	dateStr := time.Now().Format(dateFormat)
	return time.ParseInLocation(dateFormat, dateStr, time.Local)
}

/**
 * 字符串转时间
 */
func StrptimeFormat(str, format string) (time.Time, error) {
	chinaZone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Info("get china zone failed. use current time")
		return time.Now(), err
	}
	t, err := time.ParseInLocation(format, str, chinaZone)
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}

/**
 * 字符串转时间
 * 使用YYYY-mm-dd HH:MM:SS格式转换
 */
func Strptime(str string) (time.Time, error) {
	format := "2006-01-02 15:04:05"
	return StrptimeFormat(str, format)
}
