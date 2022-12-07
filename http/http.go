package requests

import (
	"net/http"
)

// ================ get请求 ================
/**
 * 发起get请求，并返回http request
 */
func Get(options ...RequestOption) (*Response, error) {
	client := New()
	method := Method("GET")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

func GetTimeout(timeout int, options ...RequestOption) (*Response, error) {
	client := New(Timeout(timeout))
	method := Method("GET")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起get请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func GetOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := Get(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起get请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func GetOrigTimeout(timeout int, options ...RequestOption) (*http.Response, error) {
	resp, err := GetTimeout(timeout, options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起get请求，并返回json
 */
func GetJson(ret interface{}, options ...RequestOption) error {
	resp, err := Get(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起get请求，并返回json
 */
func GetJsonTimeou(ret interface{}, timeout int, options ...RequestOption) error {
	resp, err := GetTimeout(timeout, options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起get请求，并返回字节
 */
func GetBytes(options ...RequestOption) ([]byte, error) {
	resp, err := Get(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

/**
 * 发起get请求，并返回字节
 */
func GetBytesTimeout(timeout int, options ...RequestOption) ([]byte, error) {
	resp, err := GetTimeout(timeout, options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

// ================ post请求 ================
/**
 * 发起post请求，并返回http request
 */
func Post(options ...RequestOption) (*Response, error) {
	client := New()
	method := Method("POST")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起post请求，并返回http request
 */
func PostTimeout(timeout int, options ...RequestOption) (*Response, error) {
	client := New(Timeout(timeout))
	method := Method("POST")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起post请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func PostOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := Post(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起post请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func PostOrigTimeout(timeout int, options ...RequestOption) (*http.Response, error) {
	resp, err := PostTimeout(timeout, options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起post请求，并返回json
 */
func PostJson(ret interface{}, options ...RequestOption) error {
	resp, err := Post(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起post请求，并返回json
 */
func PostJsonTimeout(ret interface{}, timeout int, options ...RequestOption) error {
	resp, err := PostTimeout(timeout, options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起post请求，并返回字节
 */
func PostBytes(options ...RequestOption) ([]byte, error) {
	resp, err := Post(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

/**
 * 发起post请求，并返回字节
 */
func PostBytesTimeout(timeout int, options ...RequestOption) ([]byte, error) {
	resp, err := PostTimeout(timeout, options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

// ================ put请求 ================
/**
 * 发起put请求，并返回http request
 */
func Put(options ...RequestOption) (*Response, error) {
	client := New()
	method := Method("PUT")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起put请求，并返回http request
 */
func PutTimeout(timeout int, options ...RequestOption) (*Response, error) {
	client := New(Timeout(timeout))
	method := Method("PUT")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起put请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func PutOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := Put(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起put请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func PutOrigTimeout(timeout int, options ...RequestOption) (*http.Response, error) {
	resp, err := PutTimeout(timeout, options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起put请求，并返回json
 */
func PutJson(ret interface{}, options ...RequestOption) error {
	resp, err := Put(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起put请求，并返回json
 */
func PutJsonTimeout(ret interface{}, timeout int, options ...RequestOption) error {
	resp, err := PutTimeout(timeout, options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起put请求，并返回字节
 */
func PutBytes(options ...RequestOption) ([]byte, error) {
	resp, err := Put(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

/**
 * 发起put请求，并返回字节
 */
func PutBytesTimeout(timeout int, options ...RequestOption) ([]byte, error) {
	resp, err := PutTimeout(timeout, options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

// ================ delete请求 ================
/**
 * 发起delete请求，并返回http request
 */
func Delete(options ...RequestOption) (*Response, error) {
	client := New()
	method := Method("DELETE")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起delete请求，并返回http request
 */
func DeleteTimeout(timeout int, options ...RequestOption) (*Response, error) {
	client := New(Timeout(timeout))
	method := Method("DELETE")
	options = append(options, method)
	req := NewRequest(options...)
	return client.Do(req)
}

/**
 * 发起delete请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func DeleteOrig(options ...RequestOption) (*http.Response, error) {
	resp, err := Delete(options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起delete请求，并返回原数据
 * 有可能返回的是一个大文件，可以通过此函数取到http.Response.Body，然后处理
 */
func DeleteOrigTimeout(timeout int, options ...RequestOption) (*http.Response, error) {
	resp, err := DeleteTimeout(timeout, options...)
	if err != nil {
		return nil, err
	}
	return resp.GetOrig()
}

/**
 * 发起delete请求，并返回json
 */
func DeleteJson(ret interface{}, options ...RequestOption) error {
	resp, err := Delete(options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起delete请求，并返回json
 */
func DeleteJsonTimeout(ret interface{}, timeout int, options ...RequestOption) error {
	resp, err := DeleteTimeout(timeout, options...)
	if err != nil {
		return err
	}
	return resp.ToJson(ret)
}

/**
 * 发起delete请求，并返回字节
 */
func DeleteBytes(options ...RequestOption) ([]byte, error) {
	resp, err := Delete(options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}

/**
 * 发起delete请求，并返回字节
 */
func DeleteBytesTimeout(timeout int, options ...RequestOption) ([]byte, error) {
	resp, err := DeleteTimeout(timeout, options...)
	if err != nil {
		return []byte{}, err
	}
	return resp.ToBytes()
}
