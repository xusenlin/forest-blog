package utils

import (
	"errors"
	"os"
)

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}


func IsFile(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !IsDir(dir) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

func RemoveDir(dir string) error {

	if !IsDir(dir) {
		return  errors.New("cannot delete without directory")
	}

	return  os.RemoveAll(dir)
}
