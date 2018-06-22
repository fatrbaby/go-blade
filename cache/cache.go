package cache

type Cache interface {
	Get(key string) []byte
	Set(key string, value []byte) bool
}
