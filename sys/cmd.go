package sys

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	log "github.com/qygaojs/gcommon/logger"
)

/**
 * 执行命令并获取结果
 */
func CmdOut(name string, args ...string) (string, error) {
	startTm := time.Now()
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		endTm := time.Now()
		interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
		log.Warn("server=system|func=%s|time=%d|args=%v|err=[[%s]]", name, interval, args, err)
		// 有可能err不为空，但out依然有值
		return out.String(), err
	}
	endTm := time.Now()
	interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
	respStr := out.String()
	if len(respStr) > 1024 {
		respStr = respStr[:1024]
	}
	log.Info("server=system|func=%s|time=%d|args=%v|out=[[%s]]", name, interval, args, respStr)
	return out.String(), nil
}

/**
 * 执行命令并获取结果
 */
func CmdOutBytes(name string, args ...string) ([]byte, error) {
	startTm := time.Now()
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		endTm := time.Now()
		interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
		log.Warn("server=system|func=%s|time=%d|args=%v|err=[[%s]]", name, interval, args, err)
		return out.Bytes(), err
	}
	endTm := time.Now()
	interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
	respStr := out.String()
	if len(respStr) > 1024 {
		respStr = respStr[:1024]
	}
	log.Info("server=system|func=%s|time=%d|args=%v|out=[[%s]]", name, interval, args, respStr)
	return out.Bytes(), nil
}

/**
 * 执行命令,获取结果并将结果去掉空白符
 */
func CmdOutTrim(name string, args ...string) (string, error) {
	ret, err := CmdOut(name, args...)
	if err != nil {
		return ret, err
	}
	return strings.TrimSpace(ret), nil
}

/**
 * 执行命令,设置命令执行超时，并获取结果
 */
func CmdOutWithTimeout(timeout int, name string, args ...string) (string, error) {
	startTm := time.Now()
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Start(); err != nil {
		endTm := time.Now()
		interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
		log.Warn("server=system|func=%s|time=%d|timeout=%d|args=%v|err=[[%s]]",
			name, interval, timeout, args, err)
		return out.String(), err
	}
	if err := waitTimeout(cmd, time.Duration(timeout)*time.Second); err != nil {
		endTm := time.Now()
		interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
		log.Warn("server=system|func=%s|time=%d|timeout=%d|args=%v|err=[[%s]]",
			name, interval, timeout, args, err)
		return out.String(), err
	}
	endTm := time.Now()
	interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
	respStr := out.String()
	if len(respStr) > 1024 {
		respStr = respStr[:1024]
	}
	log.Info("server=system|func=%s|time=%d|timeout=%d|args=%v|out=[[%s]]",
		name, interval, timeout, args, respStr)
	return out.String(), nil
}

/**
 * 执行命令,设置命令执行超时，获取结果,将结果去掉首尾空格
 */
func CmdOutTrimWithTimeout(timeout int, name string, args ...string) (string, error) {
	ret, err := CmdOutWithTimeout(timeout, name, args...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(ret), nil
}

func waitTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	var err error

	done := make(chan error)

	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		//log.Print("timeout process: %s will be killed", cmd.Path)
		go func() {
			<-done
		}()
		err = cmd.Process.Kill() // 命令执行超时，杀掉
		if err != nil {
			return err
		}
		return fmt.Errorf("execute command timeout")
	case err = <-done:
		return err
	}
}
