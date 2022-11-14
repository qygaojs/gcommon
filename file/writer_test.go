package file

import "testing"

func TestWriteBytes(t *testing.T) {
	fname := "./test.content"
	cnt, err := WriteBytes(fname, []byte("this is a test\n"))
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
}

func TestWriteString(t *testing.T) {
	fname := "./test.content"
	cnt, err := WriteString(fname, "this is a test\n")
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
}

func TestWriteStringList(t *testing.T) {
	fname := "./test.content"
	data := []string{"1 test", "2 test", "3 test", "4 test"}
	cnt, err := WriteStringList(fname, data)
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
}

func TestAppendBytes(t *testing.T) {
	fname := "./test.content"
	data := []string{"1 test", "2 test", "3 test", "4 test"}
	cnt, err := WriteStringList(fname, data)
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
	cnt, err = AppendBytes(fname, []byte("this is a test\n"))
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
}

func TestAppendString(t *testing.T) {
	fname := "./test.content"
	data := []string{"1 test", "2 test", "3 test", "4 test"}
	cnt, err := WriteStringList(fname, data)
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
	cnt, err = AppendString(fname, "this is a test\n")
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
}

func TestAppendStringList(t *testing.T) {
	fname := "./test.content"
	data := []string{"1 test", "2 test", "3 test", "4 test"}
	cnt, err := WriteStringList(fname, data)
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
	data = []string{"1 append", "2 append", "3 append", "4 append"}
	cnt, err = AppendStringList(fname, data)
	if err != nil {
		t.Errorf("write file failed:%s", err)
	} else {
		t.Logf("write file success:%d", cnt)
	}
}
