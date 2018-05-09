package blade

import (
	"errors"
	"io/ioutil"
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

func lastModified(filename string) time.Time {
	stat, err := os.Stat(filename)

	if err != nil {
		return time.Unix(0, 0)
	}

	return stat.ModTime()
}

func load(filename string) ([]byte, error) {
	has, err := exists(filename)

	if err != nil {
		return []byte(""), err
	}

	if !has {
		return []byte(""), errors.New("file does not exists")
	}

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return []byte(""), err
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}
