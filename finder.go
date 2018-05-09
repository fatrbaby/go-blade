package blade

import (
	"os"
	"time"
)

func exists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func lastModified(filename string) time.Time  {
	stat, err := os.Stat(filename)

	if err != nil {
		return time.Now()
	}

	return stat.ModTime()
}
