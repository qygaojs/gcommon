package cache

import (
	"errors"
	"time"
)

/**
 * key不存在的错误
 */
var ErrKeyNotExist = errors.New("key is not exist")

/**
 * key已存在的错误
 */
var ErrKeyExist = errors.New("key is exist")

/**
 * 缓存表初始化的可选参数
 */
type TableOptions struct {
	f func(*CacheTable)
}

/**
 * loadValueFunc参数初始化函数
 */
func LoadValue(loadValue LoadFunc) TableOptions {
	return TableOptions{func(obj *CacheTable) {
		obj.loadValue = loadValue
	}}
}

/**
 * 设置缓存的参数
 */
type SetParamType struct {
	nx bool          // 缓存不存在时，设置缓存
	xx bool          // 缓存存在时，设置缓存
	ex time.Duration // 缓存过期时间
}

/**
 * 设置缓存的可选参数
 */
type SetOptions struct {
	f func(*SetParamType)
}

/**
 * 缓存不存在时，设置缓存
 */
func SetNX() SetOptions {
	return SetOptions{func(obj *SetParamType) {
		obj.nx = true
	}}
}

/**
 * 缓存存在时，设置缓存
 */
func SetXX() SetOptions {
	return SetOptions{func(obj *SetParamType) {
		obj.xx = true
	}}
}

/**
 * 缓存过期时间
 * 单位: 秒
 */
func SetEX(ex int) SetOptions {
	return SetOptions{func(obj *SetParamType) {
		expire := time.Duration(ex) * time.Second
		obj.ex = expire
	}}
}
