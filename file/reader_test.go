package file

import (
	"testing"
)

func writeInt(fname string) {
	data := `1234`
	WriteString(fname, data)
}

func writeFile(fname string) {
	data := `
	aaaa
	bbbb
	cccc


	`
	WriteString(fname, data)
}

func TestReadBytes(t *testing.T) {
	fname := "./test.content"
	writeFile(fname)
	data, err := ReadBytes(fname)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%s]]", string(data))
	}
	fname = "./test.content.xxx"
	data, err = ReadBytes(fname)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%s]]", string(data))
	}
}

func TestReadString(t *testing.T) {
	fname := "./test.content"
	writeFile(fname)
	data, err := ReadString(fname)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%s]]", data)
	}
	fname = "./test.content.xxx"
	data, err = ReadString(fname)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%s]]", data)
	}
}

func TestReadTrimString(t *testing.T) {
	fname := "./test.content"
	writeFile(fname)
	data, err := ReadTrimString(fname)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%s]]", data)
	}
	fname = "./test.content.xxx"
	data, err = ReadTrimString(fname)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%s]]", data)
	}
}

func TestReadToUint(t *testing.T) {
	fname := "./test.content"
	writeInt(fname)
	data, err := ReadUint64(fname)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%d]]", data)
	}
	fname = "./test.content.xxx"
	data, err = ReadUint64(fname)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%d]]", data)
	}
}

func TestReadToInt(t *testing.T) {
	fname := "./test.content"
	writeInt(fname)
	data, err := ReadInt64(fname)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%d]]", data)
	}
	fname = "./test.content.xxx"
	data, err = ReadInt64(fname)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%d]]", data)
	}
}

func TestReadAllLine(t *testing.T) {
	fname := "./test.content"
	writeFile(fname)
	data, err := ReadAllLine(fname)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%s]]", data)
	}
	fname = "./test.content.xxx"
	data, err = ReadAllLine(fname)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%s]]", data)
	}
}

func TestReadLineFetch(t *testing.T) {
	f := func(line string) bool {
		t.Logf("read line:[[%s]]", line)
		return true
	}
	fname := "./test.content"
	writeFile(fname)
	err := ReadLineFetch(fname, f)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	}
	fname = "./test.content.xxx"
	err = ReadLineFetch(fname, f)
	if err != nil {
		t.Logf("read file failed:%s", err)
	}
}

func TestReadLineOffsetCount(t *testing.T) {
	fname := "./test.content"
	writeFile(fname)
	data, err := ReadLineOffsetCount(fname, 2, 5)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%s]]", data)
	}
	fname = "./test.content.xxx"
	data, err = ReadLineOffsetCount(fname, 10, 50)
	if err != nil {
		t.Logf("read file failed:%s", err)
	} else {
		t.Errorf("read file success:[[%s]]", data)
	}
}

func TestReadStringAt(t *testing.T) {
	fname := "./test.content"
	writeFile(fname)
	data, err := ReadStringAt(fname, 12)
	if err != nil {
		t.Errorf("read file failed:%s", err)
	} else {
		t.Logf("read file success:[[%s]]", string(data))
	}
}
