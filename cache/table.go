package cache

import (
	//"fmt"
	"sync"
	"time"
	//log "github.com/qygaojs/gcommon/logger"
)

type LoadFunc func(args ...interface{}) (interface{}, error)

type CacheTable struct {
	sync.RWMutex

	name string                     // 缓存表名
	data map[interface{}]*CacheItem // 缓存项

	clearupTimer *time.Timer // 清除过期缓存时钟

	// 访问时，value不存在，则加载value的函数
	loadValue LoadFunc
}

/**
 * 初始化
 */
func NewTable(name string, options ...TableOptions) *CacheTable {
	obj := new(CacheTable)
	obj.name = name
	obj.data = make(map[interface{}]*CacheItem)

	obj.clearupTimer = nil
	obj.loadValue = nil

	for _, option := range options {
		option.f(obj)
	}
	return obj
}

/**
 * 设置加载value的函数
 */
func (this *CacheTable) SetLoadValueFunc(f LoadFunc) {
	this.Lock()
	defer this.Unlock()
	this.loadValue = f
}

/**
 * 遍历缓存，并将每项都回调函数
 */
func (this *CacheTable) Foreach(callback func(interface{}, *CacheItem)) {
	this.RLock()
	defer this.RUnlock()
	for k, v := range this.data {
		callback(k, v)
	}
}

/**
 * 获取所有的缓存数据
 */
func (this *CacheTable) Data() map[interface{}]*CacheItem {
	this.RLock()
	defer this.RUnlock()
	return this.data
}

/**
 * 清理
 */
func (this *CacheTable) StartClearUp() {
	this.Lock()
	defer this.Unlock()

	if this.clearupTimer != nil {
		this.clearupTimer.Stop()
	}
	lessDuration := 0 * time.Second
	for key, obj := range this.data {
		if obj.IsExpire() {
			//log.Debug("clearup key:%s", key)
			delete(this.data, key)
			continue
		}
		expireTs := obj.TTL()
		if expireTs < 0 {
			continue
		}
		if lessDuration == 0 || expireTs < lessDuration {
			lessDuration = expireTs
		}
	}
	//log.Debug("next expire check:%s", lessDuration)
	if lessDuration > 0 {
		this.clearupTimer = time.AfterFunc(lessDuration, func() {
			go this.StartClearUp()
		})
	}
}

/**
 * 获取缓存信息
 */
func (this *CacheTable) Get(key interface{}) (interface{}, error) {
	this.RLock()
	value, ok := this.data[key]
	this.RUnlock()
	if ok {
		return value.Value(), nil
	} else {
		return nil, ErrKeyNotExist
	}
}

/**
 * 获取缓存信息
 * 如果缓存不存在，则加载
 */
func (this *CacheTable) GetLoad(key interface{}, args ...interface{}) (interface{}, error) {
	this.RLock()
	valObj, ok := this.data[key]
	this.RUnlock()
	if ok {
		return valObj.Value(), nil
	}
	if this.loadValue == nil {
		return nil, ErrKeyNotExist
	}

	// 加写锁，并重新判断是否key已有值，如果key还没有值，则直接去获取
	this.Lock()
	valObj, ok = this.data[key]
	if ok {
		this.Unlock()
		return valObj.Value(), nil
	}
	value, err := this.loadValue(args...)
	if err != nil {
		this.Unlock()
		return nil, err
	}
	valObj = NewItem(key, value, 0)
	this.data[key] = valObj
	this.Unlock()
	this.StartClearUp()
	return value, nil
}

/**
 * 设置缓存
 */
func (this *CacheTable) Set(key interface{}, value interface{}, options ...SetOptions) error {
	// 解析set参数
	params := new(SetParamType)
	for _, option := range options {
		option.f(params)
	}
	//fmt.Printf("params:%s\n", params)
	this.Lock()
	_, ok := this.data[key]
	if params.nx && ok {
		this.Unlock()
		return ErrKeyExist
	}
	if params.xx && !ok {
		this.Unlock()
		return ErrKeyNotExist
	}
	valueObj := NewItem(key, value, params.ex)
	this.data[key] = valueObj
	this.Unlock()
	this.StartClearUp()
	return nil
}

/**
 * 设置缓存过期时间
 */
func (this *CacheTable) Expire(key interface{}, ex int) error {
	expire := time.Duration(ex) * time.Second
	this.RLock()
	value, ok := this.data[key]
	this.RUnlock()
	if !ok {
		return ErrKeyNotExist
	}
	value.SetExpireTime(expire)
	this.StartClearUp()
	return nil
}

/**
 * 获取缓存过期时间
 * 单位为秒
 */
func (this *CacheTable) TTL(key interface{}) (int, error) {
	this.RLock()
	value, ok := this.data[key]
	this.RUnlock()
	if !ok {
		return -1, ErrKeyNotExist
	}
	ex := value.TTL()
	return int(ex.Seconds()), nil
}

/**
 * 删除key
 * 单位为秒
 */
func (this *CacheTable) Delete(key interface{}) error {
	this.Lock()
	_, ok := this.data[key]
	if !ok {
		this.Unlock()
		return nil
	}
	delete(this.data, key)
	this.Unlock()
	this.StartClearUp()
	return nil
}

type CacheAccess struct {
	Key         interface{}
	AccessCount int
	AccessTime  time.Time
}

/**
 * 删除key
 * 单位为秒
 */
func (this *CacheTable) AccessStat(key interface{}) (map[interface{}]*CacheAccess, error) {
	this.RLock()
	defer this.RUnlock()
	ret := make(map[interface{}]*CacheAccess)
	if key != nil {
		value, ok := this.data[key]
		if !ok {
			return ret, ErrKeyNotExist
		}
		info := new(CacheAccess)
		info.Key = key
		info.AccessCount = value.AccessCount()
		info.AccessTime = value.AccessTime()
		ret[key] = info
		return ret, nil
	}
	for key, value := range this.data {
		info := new(CacheAccess)
		info.Key = key
		info.AccessCount = value.AccessCount()
		info.AccessTime = value.AccessTime()
		ret[key] = info
	}
	return ret, nil
}
