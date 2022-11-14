package file

import (
	"os"
	"testing"
)

func TestFileSize(t *testing.T) {
	fname := "./file_test.go"
	size, err := Size(fname)
	if err != nil {
		t.Errorf("get %s size failed:%s", fname, err)
	} else {
		t.Logf("%s size:%d", fname, size)
	}

	fname = "/Users/gaojiansheng/source/gcommon/file"
	size, err = Size(fname)
	if err != nil {
		t.Errorf("get dir path size failed:%s", err)
	} else {
		t.Logf("%s size:%d", fname, size)
	}
}

func TestFileModTime(t *testing.T) {
	fname := "./file_test.go"
	modTime, err := ModTime(fname)
	if err != nil {
		t.Errorf("get %s mod time failed:%s", fname, err)
	} else {
		t.Logf("%s modify time:%d", fname, modTime)
	}

	fname = "./xxxx_test.go"
	modTime, err = ModTime(fname)
	if err != nil {
		t.Logf("get %s mod time failed:%s", fname, err)
	} else {
		t.Errorf("%s modify time:%d", fname, modTime)
	}

	fname = "/Users/gaojiansheng"
	modTime, err = ModTime(fname)
	if err != nil {
		t.Errorf("get %s mod time failed:%s", fname, err)
	} else {
		t.Logf("%s modify time:%d", fname, modTime)
	}
}

func TestFileIsExist(t *testing.T) {
	t.Logf("./file_test.go is exist:%t", IsExist("./file_test.go"))
	t.Logf("./xxxx_test.go is exist:%t", IsExist("./xxxx_test.go"))
}

func TestIsDir(t *testing.T) {
	t.Logf("/Users/gaojiansheng/ isDir:%t", IsDir("/Users/gaojiansheng"))
	t.Logf("./file_test.go isDir:%t", IsDir("./file_test.go"))
	t.Logf("./xxxx_test.go isDir:%t", IsDir("./xxxx_test.go"))
}

func TestDirFiles(t *testing.T) {
	fname := "/Users/gaojiansheng/"
	files, err := DirFiles(fname)
	if err != nil {
		t.Errorf("get %s files failed:%s", fname, err)
	} else {
		t.Logf("%s files:%s", fname, files)
	}

	fname = "./file_test.go"
	files, err = DirFiles(fname)
	if err != nil {
		t.Logf("get %s files failed:%s", fname, err)
	} else {
		t.Errorf("%s files:%s", fname, files)
	}
}

func TestSubDirs(t *testing.T) {
	fname := "/Users/gaojiansheng/"
	dirs, err := SubDirs(fname)
	if err != nil {
		t.Errorf("get %s sub dir failed:%s", fname, err)
	} else {
		t.Logf("%s sub dir:%s", fname, dirs)
	}

	fname = "./file_test.go"
	dirs, err = SubDirs(fname)
	if err != nil {
		t.Logf("get %s sub dir failed:%s", fname, err)
	} else {
		t.Errorf("%s sub dir:%s", fname, dirs)
	}
}

func TestDirSize(t *testing.T) {
	fname := "/Users/gaojiansheng/source/gcommon"
	size, err := DirSize(fname)
	if err != nil {
		t.Errorf("get %s size failed:%s", fname, err)
	} else {
		t.Logf("%s size:%d", fname, size)
	}

	fname = "./file_test.go"
	size, err = DirSize(fname)
	if err != nil {
		t.Logf("get %s size failed:%s", fname, err)
	} else {
		t.Errorf("%s size:%d", fname, size)
	}
}

func TestSearchFile(t *testing.T) {
	filename, err := SearchFile("file_test.go", "/Users/gaojiansheng/source/gcommon/file", "/Users/gaojiansheng")
	if err != nil {
		t.Errorf("search file failed:%s", err)
	} else {
		t.Logf("file_test.go file:%s", filename)
	}

	filename, err = SearchFile("xxxx_test.go", "/Users/gaojiansheng/source/gcommon/file", "/Users/gaojiansheng")
	if err != nil {
		t.Logf("search file failed:%s", err)
	} else {
		t.Errorf("xxxx_test.go file:%s", filename)
	}
}

func TestDeleteFile(t *testing.T) {
	fname := "test.log"
	fp, err := os.Create(fname)
	if err != nil {
		t.Errorf("create file failed:%s", err)
		return
	}
	fp.Close()

	err = Delete(fname)
	if err != nil {
		t.Errorf("delete file failed. file:%s err:%s", fname, err)
	} else {
		t.Logf("delete file success:%s", fname)
	}

	err = Delete(fname)
	if err != nil {
		t.Logf("second delete file failed. file:%s err:%s", fname, err)
	} else {
		t.Errorf("second delete file success:%s", fname)
	}
}

func TestCreateAndDelDir(t *testing.T) {
	fname := "test"
	err := CreatePath(fname)
	if err != nil {
		t.Errorf("create dir failed. dir:%s err:%s", fname, err)
	} else {
		t.Logf("create dir success:%s", fname)
	}

	err = CreatePath(fname)
	if err != nil {
		t.Logf("second create dir failed. dir:%s err:%s", fname, err)
	} else {
		t.Errorf("second create dir success:%s", fname)
	}

	err = DeletePath(fname)
	if err != nil {
		t.Errorf("del dir failed. dir:%s err:%s", fname, err)
	} else {
		t.Logf("del dir success:%s", fname)
	}

	err = DeletePath(fname)
	if err != nil {
		t.Errorf("second del dir failed. dir:%s err:%s", fname, err)
	} else {
		t.Logf("second del dir success:%s", fname)
	}
}

func TestSelfPath(t *testing.T) {
	filename, err := SelfPath()
	if err != nil {
		t.Errorf("get self path failed. err:%s", err)
	} else {
		t.Logf("get self path success:%s", filename)
	}
}

func TestSelfDir(t *testing.T) {
	dir, err := SelfDir()
	if err != nil {
		t.Errorf("get self dir failed. err:%s", err)
	} else {
		t.Logf("get self dir success:%s", dir)
	}
}

func TestRealPath(t *testing.T) {
	fname := "file_test.go"
	filename, err := RealPath(fname)
	if err != nil {
		t.Errorf("get %s real path failed:%s", fname, err)
	} else {
		t.Logf("%s real path:%s", fname, filename)
	}

	fname = "xxxx_test.go"
	filename, err = RealPath(fname)
	if err != nil {
		t.Errorf("get %s real path failed:%s", fname, err)
	} else {
		t.Logf("%s real path:%s", fname, filename)
	}

	fname = "/User/gaojs/aa.c"
	filename, err = RealPath(fname)
	if err != nil {
		t.Errorf("get %s real path failed:%s", fname, err)
	} else {
		t.Logf("%s real path:%s", fname, filename)
	}
}

func TestSignature(t *testing.T) {
	fname := "file_test.go"
	filename, err := Signature(fname)
	if err != nil {
		t.Errorf("get %s file signature failed:%s", fname, err)
	} else {
		t.Logf("%s file signature:%s", fname, filename)
	}

	fname = "xxxx_test.go"
	filename, err = Signature(fname)
	if err != nil {
		t.Logf("get %s file signature failed:%s", fname, err)
	} else {
		t.Errorf("%s file signature:%s", fname, filename)
	}

	fname = "/Users/gaojiansheng/source/gcommon/file"
	filename, err = Signature(fname)
	if err != nil {
		t.Logf("get %s file signature failed:%s", fname, err)
	} else {
		t.Errorf("%s file signature:%s", fname, filename)
	}
}
