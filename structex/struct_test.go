package structex

import (
	"testing"
)

type Test struct {
	Id   int64
	Name string
	Age  int64
}

func TestGetAttr(t *testing.T) {
	test := new(Test)
	test.Id = 1234
	test.Name = "张三"
	test.Age = 18
	if ret := GetAttr(test, "Id"); ret != nil {
		t.Logf("id:%d", ret)
	} else {
		t.Errorf("get id failed")
	}
	if ret := GetAttr(test, "Name"); ret != nil {
		t.Logf("name:%s", ret)
	} else {
		t.Errorf("get name failed")
	}
	if ret := GetAttr(test, "Age"); ret != nil {
		t.Logf("age:%d", ret)
	} else {
		t.Errorf("get age failed")
	}
	if ret := GetAttr(test, "XXX"); ret != nil {
		t.Errorf("XXX:%d", ret)
	} else {
		t.Logf("get not exist key failed")
	}
}

func TestSetAttr(t *testing.T) {
	test := new(Test)
	if err := SetAttr(test, "Id", int64(456)); err != nil {
		t.Errorf("set id failed:%s", err)
	}
	if err := SetAttr(test, "Name", "李四"); err != nil {
		t.Errorf("set name failed:%s", err)
	}
	if err := SetAttr(test, "Age", int64(34)); err != nil {
		t.Errorf("set Age failed:%s", err)
	}
	if err := SetAttr(test, "XXX", 34); err != nil {
		t.Logf("set not exist key failed:%s", err)
	} else {
		t.Errorf("set not exist key success")
	}
	t.Logf("result:%v", test)
}

func TestKeys(t *testing.T) {
	test := new(Test)
	t.Logf("keys:%v", Keys(test))
}
