package tools

import (
	"fmt"
	"math/big"
	"net"
	"reflect"
	"runtime"
	"strings"
)

/**
 * 将数字转ip
 */
func Int2Ip(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

/**
 * 将ip转数字
 */
func Ip2Int(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

/**
 * 获取本机ip
 */
func HostIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip, nil
}

/**
 * 获取函数名
 * 参数只能传函数名，否则会导致代码coredump
 * 返回值: 类名 函数名(类名格式为:模板名.类名)
 */
func GetFuncName(f interface{}) (clsName string, funcName string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	allFuncName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fields := strings.Split(allFuncName, ".")
	modName := strings.Replace(fields[0], "/", ".", -1)
	if len(fields) == 3 {
		clsName = fields[1]
		if sidx := strings.Index(clsName, "("); sidx >= 0 {
			sidx += 1
			if clsName[sidx:sidx+1] == "*" {
				sidx += 1
			}
			eidx := strings.Index(clsName, ")")
			clsName = clsName[sidx:eidx]
		}
	}
	funcName = fields[len(fields)-1]
	if eidx := strings.Index(funcName, "-fm"); eidx >= 0 {
		// 类方法名后自动带了个-fm，去掉
		funcName = funcName[:eidx]
	}
	if clsName == "" {
		clsName = modName
	} else {
		clsName = modName + "." + clsName
	}
	return
}
