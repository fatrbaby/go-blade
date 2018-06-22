package cache

import "time"

type File struct {

}

func (File) Set(key, value string, duration time.Duration) bool {
	panic("implement me")
}

func (File) Get(key string) string {
	panic("implement me")
}



