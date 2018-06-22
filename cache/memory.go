package cache

import "time"

type memory struct {
	pool map[string]*Cache
}

func (mem *memory) Set(key, value string, duration time.Duration) bool {
	mem.pool[key] = &Cache{
		Value:   value,
		Expired: time.Now().Add(duration),
	}

	_, has := mem.pool[key]

	return has
}

func (mem *memory) Get(key string) string {
	value, has := mem.pool[key]

	if !has {
		return ""
	}

	if value.Expired.Before(time.Now()) {
		return value.Value
	}

	return ""
}

func NewMemoryCache() Driver {
	memory := new(memory)
	memory.pool = make(map[string]*Cache)

	return memory
}
