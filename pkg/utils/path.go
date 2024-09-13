package utils

import "os"

func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}
