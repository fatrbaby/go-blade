package cache

import "time"

type Cache struct {
	Value string
	Expired time.Duration
}

type Driver interface {
	Set(key, value string, duration time.Duration) bool
	Get(key string) string
}
