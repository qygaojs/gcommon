package cache

import "sync"

var (
	// 分类建缓存表，可能更符合业务场景
	cache = make(map[string]*CacheTable)

	mutex sync.RWMutex
)

/**
 * 获取缓存表
 */
func Cache(name string, options ...TableOptions) *CacheTable {
	mutex.RLock()
	obj, ok := cache[name]
	mutex.RUnlock()
	if ok {
		return obj
	}

	mutex.Lock()
	defer mutex.Unlock()
	obj, ok = cache[name]
	if ok {
		return obj
	}
	obj = NewTable(name, options...)
	cache[name] = obj
	return obj
}
