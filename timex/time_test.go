package timex

import (
	"testing"
)

func TestGetChTmFormat(t *testing.T) {
	ret := GetChTmFormat("2006年01月02日 15点04分05秒")
	t.Logf("get china time format:%s", ret)
}

func TestGetChTm(t *testing.T) {
	ret := GetChTm()
	t.Logf("get china time:%s", ret)
}

func TestDate(t *testing.T) {
	ret, err := Date()
	if err != nil {
		t.Errorf("get curr date failed:%s", err)
	} else {
		t.Logf("get curr date:%s", ret)
	}
}

func TestStrptimeFormat(t *testing.T) {
	ret, err := StrptimeFormat("22-12-26 20:52:23", "06-01-02 15:04:05")
	if err != nil {
		t.Errorf("get str to timestamp failed:%s", err)
	} else {
		t.Logf("get str to timestamp:%d", ret.Unix())
	}
}

func TestStrptime(t *testing.T) {
	ret, err := Strptime("2022-12-26 20:52:23")
	if err != nil {
		t.Errorf("get str to timestamp failed:%s", err)
	} else {
		t.Logf("get str to timestamp:%d", ret.Unix())
	}
}
