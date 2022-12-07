package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	response *http.Response
}

func NewResponse(resp *http.Response) *Response {
	respData := new(Response)
	respData.response = resp
	return respData
}

/**
 * 关闭reponse
 */
func (response *Response) Close() bool {
	if response.response != nil {
		response.response.Body.Close()
		response.response = nil
	}
	return true
}

/**
 * 获取http返回码
 */
func (response *Response) Code() int {
	return response.response.StatusCode
}

/**
 * 获取http请求状态
 */
func (response *Response) Status() string {
	return response.response.Status
}

/**
 * 获取http返回码
 */
func (response *Response) Content() ([]byte, error) {
	if response.response.StatusCode < 200 || response.response.StatusCode >= 300 {
		message, _ := ioutil.ReadAll(response.response.Body)
		return []byte{}, fmt.Errorf("Status:%s, Message:%s", response.response.Status, string(message))
	}
	return ioutil.ReadAll(response.response.Body)
}

/**
 * 获取原始数据
 */
func (response *Response) GetOrig() (*http.Response, error) {
	if response.response != nil {
		return response.response, nil
	} else {
		return nil, fmt.Errorf("response is nil")
	}
}

/**
 * 反序列化http请求结果
 */
func (response *Response) ToJson(ret interface{}) error {
	defer response.Close()
	if response.response.StatusCode < 200 || response.response.StatusCode >= 300 {
		message, _ := ioutil.ReadAll(response.response.Body)
		return fmt.Errorf("Status:%s, Message:%s", response.response.Status, string(message))
	}
	err := json.NewDecoder(response.response.Body).Decode(ret)
	if err != nil {
		return fmt.Errorf("json parse data failed: %s", err)
	}
	return nil
}

/**
 * 获取http请求结果
 */
func (response *Response) ToBytes() ([]byte, error) {
	defer response.Close()
	return response.Content()
}
