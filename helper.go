package blade

import (
	"crypto/sha1"
	"fmt"
)

var (
	hasher = sha1.New()
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

func HashKey(key string) string {
	hasher.Reset()
	hasher.Write([]byte(key))

	return fmt.Sprintf("%x", hasher.Sum(nil))
}
