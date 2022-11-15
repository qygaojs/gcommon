package logger

import "testing"

func TestWrite(t *testing.T) {
	logConfig := map[string]string{
		"INFO": "test.log",
	}
	New(logConfig)
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}

func TestMultiLevel(t *testing.T) {
	logConfig := map[string]string{
		"INFO":  "test.log",
		"WARN":  "test_warn.log",
		"ERROR": "test_error.log",
	}
	New(logConfig)
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}

func TestIsPropagate(t *testing.T) {
	logConfig := map[string]string{
		"INFO": "test.log",
		"WARN": "test_warn.log",
	}
	New(logConfig, IsPropagate(true))
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}

func TestStdout(t *testing.T) {
	logConfig := map[string]string{
		"DEBUG": "stdout",
		"WARN":  "test_warn.log",
	}
	New(logConfig, IsPropagate(true))
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}

func TestRotateSize(t *testing.T) {
	logConfig := map[string]string{
		"DEBUG": "./test.log",
		"WARN":  "./test_warn.log",
	}
	New(logConfig, IsPropagate(true), RotateType(ROTATE_SIZE), MaxFileSize(481), MaxFileCount(5))
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}

func TestRotateDay(t *testing.T) {
	logConfig := map[string]string{
		"DEBUG": "./test_day.log",
		"WARN":  "./test_day_warn.log",
	}
	New(logConfig, IsPropagate(true), RotateType(ROTATE_DAY), MaxFileCount(5))
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}

func TestInstall(t *testing.T) {
	logConfig := map[string]string{
		"DEBUG": "./test.log",
		"WARN":  "./test_warn.log",
	}
	logOption := map[string]interface{}{
		"rotateType":   "size",
		"maxFileSize":  300,
		"maxFileCount": 5,
		"isPropagate":  true,
	}
	Install(logConfig, logOption)
	Info("---------------")
	Debug("this is a test:%d str:%s", 1, "test")
	Info("this is a test:%d str:%s", 2, "test")
	Notice("this is a test:%d str:%s", 3, "test")
	Note("this is a test:%d str:%s", 4, "test")
	Warning("this is a test:%d str:%s", 5, "test")
	Warn("this is a test:%d str:%s", 6, "test")
	Error("this is a test:%d str:%s", 7, "test")
	//Fatal("this is a test:%d str:%s", 8, "test")
	Close()
}
