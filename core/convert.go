package core

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var ErrNil = errors.New("value is nil")

// 类型断言方式实现
// 类型断言不太好，因为如果有自定义的类型，则识别不出来
/**
 * 转int64
 */
func Int(value interface{}) (int64, error) {
	switch value := value.(type) {
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(value).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(value).Uint()), nil
	case []byte:
		return strconv.ParseInt(string(value), 10, 64)
	case string:
		return strconv.ParseInt(value, 10, 64)
	case nil:
		return 0, ErrNil
	case error:
		return 0, value
	default:
		// 其他值不支持
		return 0, fmt.Errorf("invalid type for int, got type %T", value)
	}
}

/**
 * 转uint64
 */
func Uint(value interface{}) (uint64, error) {
	switch value := value.(type) {
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()
		if v < 0 {
			return 0, fmt.Errorf("invalid value for uint:%d", v)
		}
		return uint64(v), nil
	case uint, uint8, uint16, uint32, uint64:
		return uint64(reflect.ValueOf(value).Uint()), nil
	case []byte:
		return strconv.ParseUint(string(value), 10, 64)
	case string:
		return strconv.ParseUint(string(value), 10, 64)
	case nil:
		return 0, ErrNil
	case error:
		return 0, value
	default:
		// 其他值不支持
		return 0, fmt.Errorf("invalid type for uint, got type %T", value)
	}
}

/**
 * 转string
 */
func String(value interface{}) (string, error) {
	switch value := value.(type) {
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(value).Int(), 10), nil
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(value).Uint(), 10), nil
	case string:
		return value, nil
	case bool:
		return strconv.FormatBool(value), nil
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(value).Float(), 'f', 6, 64), nil
	case []byte:
		return string(value), nil
	case nil:
		return "", ErrNil
	case error:
		return "", value
	default:
		// 其他值不支持
		return "", fmt.Errorf("invalid type for string, got type %T", value)
	}
}

/**
 * 转string
 */
func StringQuote(value interface{}) (string, error) {
	ret, err := String(value)
	if err != nil {
		return "", err
	}
	return strconv.Quote(ret), nil
}

func Bytes(value interface{}) ([]byte, error) {
	switch value := value.(type) {
	case string:
		return []byte(value), nil
	case []byte:
		return value, nil
	case nil:
		return nil, ErrNil
	case error:
		return nil, value
	default:
		// 其他值不支持
		return nil, fmt.Errorf("invalid type for []byte, got type %T", value)
	}
}

/**
 * 转bool
 */
func Bool(value interface{}) (bool, error) {
	switch value := value.(type) {
	case bool:
		return value, nil
	case []byte:
		return strconv.ParseBool(string(value))
	case string:
		return strconv.ParseBool(string(value))
	case nil:
		return false, ErrNil
	case error:
		return false, value
	default:
		// 其他值不支持
		return false, fmt.Errorf("invalid type for uint, got type %T", value)
	}
}

/**
 * 转float
 */
func Float(value interface{}) (float64, error) {
	switch value := value.(type) {
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(value).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(value).Uint()), nil
	case float32, float64:
		return reflect.ValueOf(value).Float(), nil
	case []byte:
		return strconv.ParseFloat(string(value), 64)
	case string:
		return strconv.ParseFloat(string(value), 64)
	case nil:
		return 0, ErrNil
	case error:
		return 0, value
	default:
		// 其他值不支持
		return 0, fmt.Errorf("invalid type for uint, got type %T", value)
	}
}

func Values(value interface{}) ([]interface{}, error) {
	switch value := value.(type) {
	case []interface{}:
		return value, nil
	case nil:
		return nil, ErrNil
	case error:
		return nil, value
	default:
		// 其他值不支持
		return nil, fmt.Errorf("invalid type for []interface{}, got type %T", value)
	}
}

func Value2String(value reflect.Value) (string, error) {
	switch value.Kind() {
	case reflect.Invalid:
		if value.IsNil() {
			return "NULL", nil
		}
		return "", fmt.Errorf("invalid type, got type %s", value.Kind())
	case reflect.String:
		return value.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10), nil
	case reflect.Bool:
		return strconv.FormatBool(value.Bool()), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', 6, 64), nil
	case reflect.Interface:
		return Value2String(value.Elem())
	default:
		// 其他值不支持
		return "", fmt.Errorf("invalid type, got type %s", value.Kind())
	}
}
