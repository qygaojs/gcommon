package sys

import (
	"testing"
)

func TestCmdOut(t *testing.T) {
	ret, err := CmdOut("ls", "-la")
	if err != nil {
		t.Errorf("execute cmd failed:%s", err)
	} else {
		t.Logf("execute cmd success:%s", ret)
	}
}

func TestCmdOutBytes(t *testing.T) {
	ret, err := CmdOutBytes("ls", "-la")
	if err != nil {
		t.Errorf("execute cmd failed:%s", err)
	} else {
		t.Logf("execute cmd success:%s", string(ret))
	}
}

func TestCmdOutTrim(t *testing.T) {
	ret, err := CmdOutTrim("ls", "-la")
	if err != nil {
		t.Errorf("execute cmd failed:%s", err)
	} else {
		t.Logf("execute cmd success:%s", ret)
	}
}

func TestCmdOutTimeout(t *testing.T) {
	ret, err := CmdOutWithTimeout(1, "top")
	if err != nil {
		t.Logf("execute cmd failed:%s", err)
	} else {
		t.Errorf("execute cmd success:%s", ret)
	}
	ret, err = CmdOutWithTimeout(1, "ls", "-la")
	if err != nil {
		t.Errorf("execute cmd failed:%s", err)
	} else {
		t.Logf("execute cmd success:%s", ret)
	}
}

func TestCmdOutTrimTimeout(t *testing.T) {
	ret, err := CmdOutTrimWithTimeout(1, "top")
	if err != nil {
		t.Logf("execute cmd failed:%s", err)
	} else {
		t.Errorf("execute cmd success:%s", ret)
	}
	ret, err = CmdOutTrimWithTimeout(1, "ls", "-la")
	if err != nil {
		t.Errorf("execute cmd failed:%s", err)
	} else {
		t.Logf("execute cmd success:%s", ret)
	}
}
