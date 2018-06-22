package cache

import "time"

type memory struct {
	pool map[string]*Cache
}

func (mem *memory) Set(key, value string, duration time.Duration) bool {
	mem.pool[key] = &Cache{
		Value: value,
		Expired: duration,
	}

	_, has := mem.pool[key]

	return has
}

func (mem *memory) Get(key string) string {
	cache, has := mem.pool[key]

	if !has {
		return ""
	}

	if cache.Expired.Before(time.Now()) {
		return cache.Value
	}

	return ""
}

func NewMemoryCache() *memory {
	memory := new(memory)
	memory.pool = make(map[string]*Cache)

	return memory
}

