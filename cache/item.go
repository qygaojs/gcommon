package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex

	key   interface{}
	value interface{}

	createTime time.Time // 创建时间
	accessTime time.Time // 最后一次访问时间
	expireTime time.Time // 过期时间
	isExpire   bool      // 是否有过期时间

	accessCount int // 访问次数
}

/**
 * 初始化
 * ttl单位为秒 <=0 则代表永不过期
 */
func NewItem(key interface{}, value interface{}, ex time.Duration) *CacheItem {
	obj := new(CacheItem)
	obj.key = key
	obj.value = value

	currTime := time.Now()
	obj.createTime = currTime
	obj.accessTime = currTime
	if ex <= 0 {
		obj.isExpire = false
	} else {
		obj.isExpire = true
		obj.expireTime = currTime.Add(ex)
	}

	obj.accessCount = 0
	return obj
}

/**
 * 更新访问时间和访问次数
 */
func (this *CacheItem) KeepAlive() {
	this.Lock()
	defer this.Unlock()
	this.access()
}

/**
 * 访问缓存，更新缓存的访问次数与最后一次访问时间
 */
func (this *CacheItem) access() {
	this.accessCount += 1
	this.accessTime = time.Now()
}

/**
 * 缓存的key
 */
func (this *CacheItem) Key() interface{} {
	this.KeepAlive()
	return this.key
}

/**
 * 缓存的value
 */
func (this *CacheItem) Value() interface{} {
	this.KeepAlive()

	return this.value
}

/**
 * 更新value
 */
func (this *CacheItem) SetValue(value interface{}) {
	this.Lock()
	defer this.Unlock()
	this.access()

	this.value = value
}

/**
 * 缓存的创建时间
 */
func (this *CacheItem) CreateTime() time.Time {
	return this.createTime
}

/**
 * 缓存的访问时间
 */
func (this *CacheItem) AccessTime() time.Time {
	this.RLock()
	defer this.RUnlock()
	return this.accessTime
}

/**
 * 缓存是否过期
 */
func (this *CacheItem) IsExpire() bool {
	this.RLock()
	defer this.RUnlock()
	if !this.isExpire {
		return false
	}
	currTime := time.Now()
	return this.expireTime.Before(currTime) || this.expireTime.Equal(currTime)
}

/**
 * 缓存的存活时间
 * 如果没有过期时间，则为-1
 */
func (this *CacheItem) TTL() time.Duration {
	this.KeepAlive()

	this.RLock()
	defer this.RUnlock()
	if !this.isExpire {
		return time.Duration(-1 * time.Second)
	}
	currTime := time.Now()
	return this.expireTime.Sub(currTime)
}

/**
 * 设置缓存的过期时间
 * 如果没有过期时间，则为<=0
 */
func (this *CacheItem) SetExpireTime(ex time.Duration) {
	this.Lock()
	defer this.Unlock()
	this.access()

	if ex <= 0 {
		this.isExpire = false
	} else {
		this.isExpire = true
		currTime := time.Now()
		this.expireTime = currTime.Add(ex)
	}
}

/**
 * 缓存的访问次数
 */
func (this *CacheItem) AccessCount() int {
	this.RLock()
	defer this.RUnlock()
	return this.accessCount
}
