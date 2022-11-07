package core

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strconv"
)

/**
 * 获取调用的文件名和行号
 */
func Caller(stackNum int) (string, int) {
	// 默认调用层级是3层，有可能会更多
	_, file, line, ok := runtime.Caller(stackNum)
	if !ok {
		return "", 0
	}
	shortFn := filepath.Base(file)
	return shortFn, line
}

/**
 * 获取调用该函数的文件名和行号
 */
func SelfCaller() (string, int) {
	return Caller(2)
}

/**
 * 获取gid
 * 性能不高，谨慎使用
 * 经测试：平均时间为0.1ms左右
 */
func GoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
