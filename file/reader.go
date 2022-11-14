package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/**
 * 读取文件所有内容，并将文件内容返回
 */
func ReadBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

/**
 * 从指定位置开始读文件
 */
func ReadBytesAt(filename string, offset int64) ([]byte, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	defer fp.Close()

	ret := []byte{}
	for {
		b := make([]byte, 1024, 1024)
		n, err := fp.ReadAt(b, offset)
		if n > 0 {
			b = b[:n] // 截断，以防不满1024
			ret = append(ret, b...)
			offset += int64(n)
		}
		if n == 0 || err == io.EOF {
			break
		}
	}

	return ret, nil
}

/**
 * 读取文件所有内容，并将文件内容转字符串返回
 */
func ReadString(filename string) (string, error) {
	data, err := ReadBytes(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/**
 * 从指定位置开始读文件
 */
func ReadStringAt(filename string, offset int64) (string, error) {
	data, err := ReadBytesAt(filename, offset)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/**
 * 读取文件所有内容，并将文件内容转字符串返回
 * 文件内容去掉首尾空白字符
 */
func ReadTrimString(filename string) (string, error) {
	data, err := ReadString(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(data), nil
}

/**
 * 读取文件所有内容，并将文件内容转uint64
 */
func ReadUint64(filename string) (uint64, error) {
	data, err := ReadTrimString(filename)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

/**
 * 读取文件所有内容，并将文件内容转int64
 */
func ReadInt64(filename string) (int64, error) {
	data, err := ReadTrimString(filename)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

/**
 * 读取文件所有内容，返回当前行
 */
func ReadLine(r *bufio.Reader) (string, error) {
	line, err := r.ReadString(byte('\n'))
	if err != nil {
		return "", err
	}
	return strings.Trim(line, "\n"), nil
}

/**
 * 读取文件所有内容，并将文件内容按行返回
 */
func ReadAllLine(filename string) ([]string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()

	var ret []string

	r := bufio.NewReader(fp)

	for {
		line, err := ReadLine(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return []string{}, err //非eof的错误，直接返回
		}
		line = strings.Trim(line, "\n")
		ret = append(ret, line)
	}
	return ret, nil
}

/**
 * 迭代器方式读取文件
 * 如果f返回false，则停止读取
 */
func ReadLineFetch(filename string, f func(string) bool) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	r := bufio.NewReader(fp)

	for {
		line, err := ReadLine(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err //非eof的错误，直接返回
		}
		if !f(line) {
			break // 不处理了，直接退出
		}
	}
	return nil
}

/**
 * 读取文件所有内容，并将文件指定行
 */
func ReadLineOffsetCount(filename string, offset uint, cnt uint) ([]string, error) {
	if cnt == 0 {
		return []string{}, nil
	}
	fp, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()

	var ret []string

	r := bufio.NewReader(fp)
	for i := 0; i < int(offset+cnt); i++ {
		line, err := ReadLine(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return []string{}, err
		}
		if i < int(offset) {
			continue
		}
		line = strings.Trim(line, "\n")
		ret = append(ret, line)
	}
	return ret, nil
}
