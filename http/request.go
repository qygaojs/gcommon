package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/qygaojs/gcommon/core"
)

type Request struct {
	method string
	url    string
	header map[string]interface{}
	body   []byte                 // http请求数据，写到body中
	query  map[string]interface{} // http请求数据, 拼接到url后
	isLong bool                   // 是否使用长连接
}

type RequestOption struct {
	f func(*Request)
}

func Method(method string) RequestOption {
	return RequestOption{func(req *Request) {
		req.method = method
	}}
}

func Url(url string) RequestOption {
	return RequestOption{func(req *Request) {
		req.url = url
	}}
}

func Header(header map[string]interface{}) RequestOption {
	return RequestOption{func(req *Request) {
		req.header = header
	}}
}

func Body(body []byte) RequestOption {
	return RequestOption{func(req *Request) {
		req.body = body
	}}
}

func Query(query map[string]interface{}) RequestOption {
	return RequestOption{func(req *Request) {
		req.query = query
	}}
}

func IsLong(isLong bool) RequestOption {
	return RequestOption{func(req *Request) {
		req.isLong = isLong
	}}
}

func NewRequest(options ...RequestOption) *Request {
	request := new(Request)
	// 初始化
	request.header = make(map[string]interface{})
	request.query = make(map[string]interface{})

	for _, option := range options {
		option.f(request)
	}
	return request
}

func (request *Request) GetMethod() string {
	return request.method
}

func (request *Request) SetMethod(method string) {
	request.method = method
}

func (request *Request) GetUrl() string {
	return request.url
}

func (request *Request) SetUrl(url string) {
	request.url = url
}

func (request *Request) GetHeader() map[string]interface{} {
	return request.header
}

func (request *Request) SetHeader(header map[string]interface{}) {
	request.header = header
}

func (request *Request) GetBody() []byte {
	return request.body
}

func (request *Request) SetBody(body []byte) {
	request.body = body
}

func (request *Request) GetQuery() map[string]interface{} {
	return request.query
}

func (request *Request) SetQuery(query map[string]interface{}) {
	request.query = query
}

/**
 * 将请求数据转成value
 */
func (request *Request) Data2Values() (url.Values, error) {
	reqValues := url.Values{}
	for key, value := range request.query {
		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			for i := 0; i < val.Len(); i++ {
				item, err := core.Value2String(val.Index(i))
				if err != nil {
					return reqValues, err
				}
				reqValues.Add(key, item)
			}
		} else {
			item, err := core.Value2String(val)
			if err != nil {
				return reqValues, err
			}
			reqValues.Add(key, item)
		}
	}
	return reqValues, nil
}

/**
 * 生成url
 * url可能原来也有参数，要把query和url的参数拼一起
 */
func (request *Request) GenUrl() (string, error) {
	u, err := url.Parse(request.url)
	if err != nil {
		return "", err
	}
	reqValues, err := request.Data2Values()
	if err != nil {
		return "", err
	}
	urlValues := u.Query()
	// 有可能url中也带了参数，所以需要合并reqValues和urlValues
	for key, values := range urlValues {
		for _, val := range values {
			reqValues.Add(key, val)
		}
	}
	u.RawQuery = reqValues.Encode()
	return u.String(), nil
}

/**
 * 获取http header
 */
func (request *Request) ToHeader() (http.Header, error) {
	header := http.Header{}
	for key, value := range request.header {
		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			for i := 0; i < val.Len(); i++ {
				item, err := core.Value2String(val.Index(i))
				if err != nil {
					return header, err
				}
				header.Add(key, item)
			}
		} else {
			item, err := core.Value2String(val)
			if err != nil {
				return header, err
			}
			header.Add(key, item)
		}
	}
	return header, nil
}

/**
 * 转换http request
 */
func (request *Request) HttpRequest() (*http.Request, error) {
	// 获取url
	urlStr, err := request.GenUrl()
	if err != nil {
		return nil, err
	}
	// 创建http request
	req, err := http.NewRequest(request.method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	// 获取body
	if len(request.body) > 0 {
		reader := bytes.NewReader(request.body)
		req.Body = ioutil.NopCloser(reader)
	}
	if !request.isLong {
		req.Close = true
	}
	// 设置http header
	header, err := request.ToHeader()
	if err != nil {
		return nil, err
	}
	req.Header = header
	req.ParseForm()
	return req, nil
}

/**
 * 获取请求参数
 */
func (request *Request) Args() (string, error) {
	data := map[string]interface{}{
		"header": request.header,
		"body":   string(request.body),
		"query":  request.query,
	}
	reqBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(reqBytes), nil
}
