package requests

import (
	"encoding/json"
	"testing"
)

func TestReqData(t *testing.T) {
	reqData := map[string]interface{}{
		"aa": 1,
		"bb": []int{1, 2, 3},
	}
	resp := make(map[string]interface{})
	err := GetJson(&resp, Url("http://127.0.0.1:8080/test?aa=10&cc=100"), Query(reqData))
	if err != nil {
		t.Errorf("request failed:%s", err)
	} else {
		respBytes, _ := json.Marshal(resp)
		t.Logf("request ret:%s", string(respBytes))
	}
}

func TestBody(t *testing.T) {
	reqData := map[string]interface{}{
		"aa": 1,
		"bb": []int{1, 2, 3},
	}
	reqBytes, err := json.Marshal(reqData)
	if err != nil {
		t.Errorf("reqData convert failed:%s", err)
		return
	}

	resp := make(map[string]interface{})
	err = GetJson(&resp, Url("http://127.0.0.1:8080/test?aa=10&cc=100"), Body(reqBytes))
	if err != nil {
		t.Errorf("request failed:%s", err)
	} else {
		respBytes, _ := json.Marshal(resp)
		t.Logf("request ret:%s", string(respBytes))
	}
}

func TestHead(t *testing.T) {
	reqData := map[string]interface{}{
		"aa": 1,
		"bb": []int{1, 2, 3},
	}

	resp := make(map[string]interface{})
	err := GetJson(&resp, Url("http://127.0.0.1:8080/test?aa=10&cc=100"), Header(reqData))
	if err != nil {
		t.Errorf("request failed:%s", err)
	} else {
		respBytes, _ := json.Marshal(resp)
		t.Logf("request ret:%s", string(respBytes))
	}
}

func TestBytes(t *testing.T) {
	reqData := map[string]interface{}{
		"aa": 1,
		"bb": []int{1, 2, 3},
	}

	resp, err := GetBytes(Url("http://127.0.0.1:8080/test?aa=10&cc=100"), Header(reqData))
	if err != nil {
		t.Errorf("request failed:%s", err)
	} else {
		t.Logf("request ret:%s", string(resp))
	}
}

func TestResponse(t *testing.T) {
	resp, err := Get(Url("http://127.0.0.1:8080/test"))
	if err != nil {
		t.Errorf("request failed:%s", err)
		return
	}
	defer resp.Close()
	if code := resp.Code(); code != 200 {
		t.Errorf("http response code:%d", code)
		return
	}
	if content, err := resp.Content(); err != nil {
		t.Errorf("get http response content failed:%s", err)
	} else {
		t.Logf("get http response content :%s", content)
	}
}

func TestTimeout(t *testing.T) {
	resp, err := GetBytesTimeout(1, Url("http://127.0.0.1:8080/test_timeout"))
	if err != nil {
		t.Logf("request failed:%s", err)
	} else {
		t.Errorf("request ret:%s", string(resp))
	}
}

func TestRedirect(t *testing.T) {
	resp, err := Get(Url("http://127.0.0.1:8080/test_redirect"))
	if err != nil {
		t.Errorf("request failed:%s", err)
		return
	}
	defer resp.Close()
	if code := resp.Code(); code != 200 {
		t.Errorf("http response code:%d", code)
		return
	}
	if content, err := resp.Content(); err != nil {
		t.Errorf("get http response content failed:%s", err)
	} else {
		t.Logf("get http response content :%s", content)
	}
}
