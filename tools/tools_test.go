package tools

import (
	"testing"
)

func TestIp2Int(t *testing.T) {
	ret := Ip2Int("10.222.16.4")
	t.Logf("10.222.16.4 to int:%d", ret)

	ret = Ip2Int("127.0.0.1")
	t.Logf("127.0.0.1 to int:%d", ret)

	ret = Ip2Int("0.0.0.0")
	t.Logf("0.0.0.0 to int:%d", ret)

	ret = Ip2Int("255.255.255.255")
	t.Logf("255.255.255.255 to int:%d", ret)
}

func TestInt2Ip(t *testing.T) {
	ret := Int2Ip(182325252)
	t.Logf("int2ip:%s", ret)

	ret = Int2Ip(0)
	t.Logf("int2ip:%s", ret)

	ret = Int2Ip(1234123412)
	t.Logf("int2ip:%s", ret)

	ret = Int2Ip(-1)
	t.Logf("int2ip:%s", ret)

	ret = Int2Ip(-100)
	t.Logf("int2ip:%s", ret)
}

func TestHostIp(t *testing.T) {
	ret, err := HostIp()
	if err != nil {
		t.Errorf("get host ip failed:%s", err)
	} else {
		t.Logf("get host ip success:%s", ret)
	}
}

func TestGetFuncName(t *testing.T) {
	clsName, funcName, err := GetFuncName(TestHostIp)
	if err != nil {
		t.Errorf("get funcName failed:%s", err)
	} else {
		t.Logf("get funcName success. cls:%s func:%s", clsName, funcName)
	}

	clsName, funcName, err = GetFuncName(t.Errorf)
	if err != nil {
		t.Errorf("get funcName failed:%s", err)
	} else {
		t.Logf("get funcName success. cls:%s func:%s", clsName, funcName)
	}

	clsName, funcName, err = GetFuncName(clsName)
	if err != nil {
		t.Errorf("get funcName failed:%s", err)
	} else {
		t.Logf("get funcName success. cls:%s func:%s", clsName, funcName)
	}
}
