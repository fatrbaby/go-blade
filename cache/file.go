package cache

import (
	"time"
)

type fileDriver struct {
}

func (fd *fileDriver) Set(key, value string, duration time.Duration) bool {
	panic("implement me")
}

func (fd *fileDriver) Get(key string) string {
	panic("implement me")
}
