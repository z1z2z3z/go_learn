/*
 * @Description: 自定义的map(边读边写)
 */
package definetype

import "sync"

type RWMutexMap struct {
	lock sync.RWMutex
	m    map[string]int
}

func (receiver *RWMutexMap) Get(key string) (int, bool) {
	receiver.lock.RLock()
	value, ok := receiver.m[key]
	receiver.lock.RUnlock()
	return value, ok
}

func (receiver *RWMutexMap) Set(key string, value int) {
	receiver.lock.Lock()
	receiver.m[key] = value
	receiver.lock.Unlock()
}
func (receiver *RWMutexMap) Del(key string) {
	receiver.lock.Lock()
	delete(receiver.m, key)
	receiver.lock.Unlock()
}
