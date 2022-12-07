package requests

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/qygaojs/gcommon/logger"
)

type HttpClient struct {
	client *http.Client

	timeout int // http请求超时
}

type HttpOption struct {
	f func(*HttpClient)
}

// http请求超时，默认10秒
func Timeout(timeout int) HttpOption {
	return HttpOption{func(req *HttpClient) {
		req.timeout = timeout
	}}
}

func New(options ...HttpOption) *HttpClient {
	client := new(HttpClient)
	for _, option := range options {
		option.f(client)
	}
	// http请求默认超时10s
	if client.timeout <= 0 {
		client.timeout = 10
	}
	// 创建http客户端
	client.client = &http.Client{
		Timeout: time.Duration(client.timeout) * time.Second,
	}
	return client
}

func operateLog(req *Request) func(ret **Response, err *error) {
	startTm := time.Now()
	return func(ret **Response, err *error) {
		reqStr, _ := req.Args()
		if len(reqStr) > 1024 {
			reqStr = reqStr[:1024]
		}
		endTm := time.Now()
		interval := (endTm.UnixNano() / 1000) - (startTm.UnixNano() / 1000)
		if panic_err := recover(); panic_err != nil {
			log.Error("server=HttpClient|url=%s|method=%s|time=%d|args=%s|exception=[[%s]]",
				req.url, req.method, interval, reqStr, panic_err)
			*err = fmt.Errorf("%s", panic_err)
			return
		}
		//redis_ret, redis_err := redis.String(*ret, *err)
		if *err != nil {
			log.Warn("server=HttpClient|url=%s|method=%s|time=%d|args=%s|err=[[%s]]",
				req.url, req.method, interval, reqStr, *err)
		} else {
			status := (*ret).Status()
			log.Info("server=HttpClient|url=%s|method=%s|time=%d|args=%s|status=%s",
				req.url, req.method, interval, reqStr, status)
		}
		return
	}
}

/**
 * 发起请求
 * 直接将请求数据返回
 * 可能存在一次返回很大的数据的问题，以及产生的攻击, 暂时不考虑
 */
func (client *HttpClient) Do(req *Request) (respData *Response, err error) {
	defer operateLog(req)(&respData, &err)
	request, err := req.HttpRequest()
	if err != nil {
		return
	}
	resp, err := client.client.Do(request)
	if err != nil || resp == nil {
		return
	}
	respData = NewResponse(resp)
	return
}

// ================ get请求 ================
/**
 * 发起get请求，并返回http request
 */
func (client *HttpClient) Get(options ...RequestOption) (*Response, error) {
	method := Method("GET")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起get请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func (client *HttpClient) GetOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := client.Get(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起get请求，并返回json
 */
func (client *HttpClient) GetJson(ret interface{}, options ...RequestOption) error {
	resp, err := client.Get(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起get请求，并返回字节
 */
func (client *HttpClient) GetBytes(options ...RequestOption) ([]byte, error) {
	resp, err := client.Get(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

// ================ post请求 ================
/**
 * 发起post请求，并返回http request
 */
func (client *HttpClient) Post(options ...RequestOption) (*Response, error) {
	method := Method("POST")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起post请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func (client *HttpClient) PostOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := client.Post(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起post请求，并返回json
 */
func (client *HttpClient) PostJson(ret interface{}, options ...RequestOption) error {
	resp, err := client.Post(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起post请求，并返回字节
 */
func (client *HttpClient) PostBytes(options ...RequestOption) ([]byte, error) {
	resp, err := client.Post(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

// ================ put请求 ================
/**
 * 发起put请求，并返回http request
 */
func (client *HttpClient) Put(options ...RequestOption) (*Response, error) {
	method := Method("PUT")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起put请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func (client *HttpClient) PutOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := client.Put(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起put请求，并返回json
 */
func (client *HttpClient) PutJson(ret interface{}, options ...RequestOption) error {
	resp, err := client.Put(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起put请求，并返回字节
 */
func (client *HttpClient) PutBytes(options ...RequestOption) ([]byte, error) {
	resp, err := client.Put(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

// ================ delete请求 ================
/**
 * 发起delete请求，并返回http request
 */
func (client *HttpClient) Delete(options ...RequestOption) (*Response, error) {
	method := Method("DELETE")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起delete请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func (client *HttpClient) DeleteOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := client.Delete(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起delete请求，并返回json
 */
func (client *HttpClient) DeleteJson(ret interface{}, options ...RequestOption) error {
	resp, err := client.Delete(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起delete请求，并返回字节
 */
func (client *HttpClient) DeleteBytes(options ...RequestOption) ([]byte, error) {
	resp, err := client.Delete(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}
