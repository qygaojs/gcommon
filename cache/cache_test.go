package cache

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSingle(t *testing.T) {
	test := Cache("test")
	err := test.Set("aa", 1234)
	if err != nil {
		t.Errorf("set aa cache failed:%s", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s", err)
		return
	} else {
		t.Logf("aa cache value:%d", value)
	}

	err = test.Set("bb", "bbbbbb")
	if err != nil {
		t.Errorf("set bb cache failed:%s", err)
		return
	}
	bValue, err := test.Get("bb")
	if err != nil {
		t.Errorf("get bb cache failed:%s", err)
		return
	} else {
		t.Logf("bb cache value:%s", bValue)
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s", err)
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestTimeout(t *testing.T) {
	test := Cache("test")
	err := test.Set("aa", 4567, SetEX(1))
	if err != nil {
		t.Errorf("set aa cache failed:%s", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s", err)
		return
	} else {
		t.Logf("aa cache value:%d", value)
	}

	err = test.Set("bb", "bbbbbb", SetEX(2))
	if err != nil {
		t.Errorf("set bb cache failed:%s", err)
		return
	}
	bValue, err := test.Get("bb")
	if err != nil {
		t.Errorf("get bb cache failed:%s", err)
		return
	} else {
		t.Logf("bb cache value:%s", bValue)
	}

	err = test.Set("cc", "ccccc")
	if err != nil {
		t.Errorf("set cc cache failed:%s", err)
		return
	}
	cValue, err := test.Get("cc")
	if err != nil {
		t.Errorf("get cc cache failed:%s", err)
		return
	} else {
		t.Logf("cc cache value:%s", cValue)
	}
	time.Sleep(time.Duration(3) * time.Second)
	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s", err)
	} else {
		t.Errorf("aa cache value:%d", value)
		return
	}
	value, err = test.Get("bb")
	if err != nil {
		t.Logf("get bb cache failed:%s", err)
	} else {
		t.Errorf("bb cache value:%s", value)
		return
	}
	value, err = test.Get("cc")
	if err != nil {
		t.Errorf("get cc cache failed:%s", err)
		return
	} else {
		t.Logf("cc cache value:%s", value)
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s", err)
		return
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestNx(t *testing.T) {
	test := Cache("test")
	err := test.Set("aa", 123)
	if err != nil {
		t.Errorf("set aa cache failed:%s", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s", err)
		return
	} else {
		t.Logf("aa cache value:%d", value)
	}

	err = test.Set("aa", 456)
	if err != nil {
		t.Errorf("set aa cache failed:%s", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s", err)
		return
	} else {
		t.Logf("aa cache value:%d", value)
	}

	err = test.Set("aa", 789, SetNX())
	if err != nil {
		t.Logf("set aa cache failed:%s", err)
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s", err)
		return
	} else {
		t.Logf("aa cache value:%d", value)
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s", err)
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestXx(t *testing.T) {
	test := Cache("testXX")
	err := test.Set("aa", 123, SetXX())
	if err != nil {
		t.Logf("set aa cache failed:%s\n", err)
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}

	err = test.Set("aa", 456)
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}

	err = test.Set("aa", 789, SetXX())
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestSet(t *testing.T) {
	test := Cache("testSet")
	err := test.Set("aa", 123, SetNX(), SetEX(1))
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	time.Sleep(time.Duration(2) * time.Second)

	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}

	err = test.Set("aa", 456, SetNX(), SetEX(1))
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}

	err = test.Set("aa", 789, SetXX(), SetEX(2))
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	time.Sleep(time.Duration(3) * time.Second)

	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}
	err = test.Set("aa", 110, SetXX(), SetNX(), SetEX(1))
	if err != nil {
		t.Logf("set aa cache failed:%s\n", err)
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
		return
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestExpire(t *testing.T) {
	test := Cache("testExpire")
	err := test.Set("aa", 123)
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	test.Expire("aa", 1)

	time.Sleep(time.Duration(2) * time.Second)

	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
		return
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestTtl(t *testing.T) {
	test := Cache("testTtl")
	err := test.Set("aa", 123)
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	ttl, err := test.TTL("aa")
	if err != nil {
		t.Errorf("get aa ttl failed:%s\n", err)
		return
	} else {
		t.Logf("aa ttl:%d\n", ttl)
	}
	test.Expire("aa", 3)

	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	ttl, err = test.TTL("aa")
	if err != nil {
		t.Errorf("get aa ttl failed:%s\n", err)
		return
	} else {
		t.Logf("aa ttl:%d\n", ttl)
	}
	time.Sleep(1 * time.Second)

	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	ttl, err = test.TTL("aa")
	if err != nil {
		t.Errorf("get aa ttl failed:%s\n", err)
		return
	} else {
		t.Logf("aa ttl:%d\n", ttl)
	}
	time.Sleep(1 * time.Second)

	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	ttl, err = test.TTL("aa")
	if err != nil {
		t.Errorf("get aa ttl failed:%s\n", err)
		return
	} else {
		t.Logf("aa ttl:%d\n", ttl)
	}
	time.Sleep(1 * time.Second)

	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
	}
	ttl, err = test.TTL("aa")
	if err != nil {
		t.Logf("get aa ttl failed:%s\n", err)
	} else {
		t.Errorf("aa ttl:%d\n", ttl)
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestDelete(t *testing.T) {
	test := Cache("testDelete")
	err := test.Set("aa", 123)
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	err = test.Delete("aa")
	if err != nil {
		t.Errorf("delete aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}

	err = test.Set("aa", 123, SetEX(1))
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}

	err = test.Delete("aa")
	if err != nil {
		t.Errorf("delete aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}

	err = test.Delete("aa")
	if err != nil {
		t.Errorf("delete aa cache failed:%s\n", err)
		return
	}
	value, err = test.Get("aa")
	if err != nil {
		t.Logf("get aa cache failed:%s\n", err)
	} else {
		t.Errorf("aa cache value:%d\n", value)
		return
	}

	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func TestMulti(t *testing.T) {
	test := Cache("testMulti")
	err := test.Set("aa", 1234)
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err := test.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%d\n", value)
	}
	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
		return
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}

	test1 := Cache("testMulti1")
	err = test1.Set("aa", "aaaaaaa")
	if err != nil {
		t.Errorf("set aa cache failed:%s\n", err)
		return
	}
	value, err = test1.Get("aa")
	if err != nil {
		t.Errorf("get aa cache failed:%s\n", err)
		return
	} else {
		t.Logf("aa cache value:%s\n", value)
	}

	test1AccessInfo, err := test1.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
		return
	} else {
		t.Logf("table access info:")
		for k, v := range test1AccessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}

func load(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("params is invalid")
	}
	a, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("params is invalid")
	}
	ret := a + "load"
	return ret, nil
}

func TestGetLoad(t *testing.T) {
	test := Cache("testLoad", LoadValue(load))
	value, err := test.GetLoad("aa", "aaaaa")
	if err != nil {
		t.Errorf("get aa failed:%s\n", err)
		return
	} else {
		t.Logf("get aa value:%s\n", value)
	}

	value, err = test.GetLoad("aa", "aaaaa11111")
	if err != nil {
		t.Errorf("get aa failed:%s\n", err)
		return
	} else {
		t.Logf("get aa value:%s\n", value)
	}

	value, err = test.GetLoad("bb", 1234)
	if err != nil {
		t.Logf("get bb failed:%s\n", err)
	} else {
		t.Errorf("get bb value:%s\n", value)
		return
	}

	accessInfo, err := test.AccessStat(nil)
	if err != nil {
		t.Errorf("get table access info failed:%s\n", err)
		return
	} else {
		t.Logf("table access info:")
		for k, v := range accessInfo {
			accessBytes, _ := json.Marshal(v)
			t.Logf("    key:%s value:%s", k, string(accessBytes))
		}
	}
}
