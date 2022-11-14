package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"crypto/md5"
)

/**
 * 获取文件大小
 */
func Size(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	fileSize := fileInfo.Size() //获取size
	return fileSize, nil
}

/**
 * 获取文件修改时间
 */
func ModTime(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.ModTime().Unix(), nil
}

/**
 * 判断文件或目录是否存在
 */
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

/**
 * 判断是否是目录
 */
func IsDir(dirPath string) bool {
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		//fmt.Println(err)
		return false
	}
	return fileInfo.IsDir()
}

/**
 * 获取目录下所有的文件
 */
func DirFiles(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, fmt.Errorf("%s is not exist", dirPath)
	}
	if !IsDir(dirPath) {
		return []string{}, fmt.Errorf("%s is not path", dirPath)
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	ret := []string{}
	for _, fi := range fs {
		if !fi.IsDir() {
			ret = append(ret, fi.Name())
		}
	}

	return ret, nil
}

/**
 * 获取目录下所有的子目录
 */
func SubDirs(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, fmt.Errorf("%s is not exist", dirPath)
	}
	if !IsDir(dirPath) {
		return []string{}, fmt.Errorf("%s is not path", dirPath)
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	ret := []string{}
	for _, fi := range fs {
		if !fi.IsDir() {
			continue
		}
		name := fi.Name()
		if name == "." || name == ".." {
			continue
		}
		ret = append(ret, name)
	}

	return ret, nil
}

/**
 * 获取目录的大小
 * 所有文件及子目录中文件大小之和
 */
func DirSize(dirPath string) (int64, error) {
	if !IsExist(dirPath) {
		return 0, fmt.Errorf("%s is not exist", dirPath)
	}
	if !IsDir(dirPath) {
		return 0, fmt.Errorf("%s is not path", dirPath)
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return 0, err
	}

	var totalSize int64 = 0
	for _, fi := range fs {
		if !fi.IsDir() {
			size := fi.Size()
			totalSize += size
		} else {
			subDir := filepath.Join(dirPath, fi.Name())
			size, err := DirSize(subDir)
			if err != nil {
				return 0, err
			}
			totalSize += size
		}
	}

	return totalSize, nil
}

/**
 * 从目录中查找文件，并返回文件的完整路径
 */
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = fmt.Errorf("%s not found in paths", filename)
	return
}

/**
 * 删除文件
 */
func Delete(filename string) error {
	return os.Remove(filename)
}

/**
 * 创建文件夹
 */
func CreatePath(path string) error {
	return os.Mkdir(path, 0755)
}

/**
 * 删除文件夹
 */
func DeletePath(path string) error {
	return os.RemoveAll(path)
}

/**
 * 获取程序自己的全路径文件名
 */
func SelfPath() (string, error) {
	return filepath.Abs(os.Args[0])
}

/**
 * 获取程序自己的全路径目录
 */
func SelfDir() (string, error) {
	filename, err := SelfPath()
	if err != nil {
		return "", nil
	}
	return filepath.Dir(filename), nil
}

/**
 * 获取文件的全路径文件名
 * 如果文件本身是全路径文件，则直接返回
 * 否则将文件拼上当前目录
 */
func RealPath(filename string) (string, error) {
	if filepath.IsAbs(filename) {
		return filename, nil
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(pwd, filename), nil
}

/**
 * 用md5算法计算文件签名
 */
func Signature(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("open file failed:%s", err)
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("read file failed:%s", err)
	}
	return fmt.Sprintf("%x\n", md5.Sum(body)), nil
}
