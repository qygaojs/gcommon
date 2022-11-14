package file

import (
	"os"
	"path"
)

/**
 * 写文件并指定文件模式
 */
func WriteBytesMode(filename string, b []byte, mode os.FileMode) (int, error) {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, mode)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	return fp.Write(b)
}

/**
 * 写文件
 */
func WriteBytes(filename string, b []byte) (int, error) {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fp, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	return fp.Write(b)
}

/**
 * 写文件
 */
func WriteString(filename string, s string) (int, error) {
	return WriteBytes(filename, []byte(s))
}

/**
 * 写多行到文件
 */
func WriteStringList(filename string, str_list []string) (int, error) {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fp, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer fp.Close()

	total := 0
	for _, s := range str_list {
		s += "\n"
		cnt, err := fp.Write([]byte(s))
		if err != nil {
			return 0, err
		}
		total += cnt
	}
	return total, nil
}

/**
 * 在文件尾追加
 */
func AppendBytes(filename string, b []byte) (int, error) {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	return fp.Write(b)
}

/**
 * 在文件尾追加
 */
func AppendString(filename string, s string) (int, error) {
	return AppendBytes(filename, []byte(s))
}

/**
 * 在文件尾追加多行
 */
func AppendStringList(filename string, str_list []string) (int, error) {
	err := os.MkdirAll(path.Dir(filename), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	total := 0
	for _, s := range str_list {
		s += "\n"
		cnt, err := fp.Write([]byte(s))
		if err != nil {
			return 0, err
		}
		total += cnt
	}
	return total, nil
}
