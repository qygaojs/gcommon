package core

import "testing"

func TestCaller(t *testing.T) {
	file, line := Caller(0)
	t.Logf("file:%s line:%d\n", file, line)
	file, line = Caller(1)
	t.Logf("file:%s line:%d", file, line)
	file, line = Caller(2)
	t.Logf("file:%s line:%d", file, line)
	file, line = Caller(-1)
	t.Logf("file:%s line:%d", file, line)
	file, line = Caller(10000)
	t.Logf("file:%s line:%d", file, line)
}

func TestSelfCaller(t *testing.T) {
	file, line := SelfCaller()
	t.Logf("file:%s line:%d", file, line)
}

func TestGid(t *testing.T) {
	gid := GoroutineID()
	t.Logf("gid:%d", gid)
}
