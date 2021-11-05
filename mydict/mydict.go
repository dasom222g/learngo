package mydict

import (
	"errors"
)

type Dictionary map[string]string

var errorNotFound = errors.New("Not found")

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		// key 가 존재하는 경우
		return value, nil
	}
	// key가 존재하지 않는 경우
	return value, errorNotFound
}
