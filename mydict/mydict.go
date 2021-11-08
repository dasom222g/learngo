package mydict

import (
	"errors"
	// "fmt"
)

// hashmap 타입은 기본적으로 포인터를 내포하므로 receiver는 복제가 아닌 참조를 하여 실제 데이터를 조작함
type Dictionary map[string]string

var (
	errorNotFound   = errors.New("Not found")
	errorExists     = errors.New("That word alreay exists")
	errorCantUpdate = errors.New("Can't update non-existing word")
)

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		// key 가 존재하는 경우
		return value, nil
	}
	// key가 존재하지 않는 경우
	return value, errorNotFound
}

func (d Dictionary) Add(key, definition string) error {
	_, err := d.Search(key)
	switch err {
	case errorNotFound: // 존재하지 않는 경우 단어 추가
		d[key] = definition
	case nil: // 이미 존재하는 경우
		return errorExists
	}
	return nil
}

func (d Dictionary) Update(key, definition string) error {
	_, err := d.Search(key)
	switch err {
	case errorNotFound: // 값이 없는 경우
		return errorCantUpdate
	case nil: // 값이 있는 경우
		d[key] = definition
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case errorNotFound:
		return err
	case nil:
		delete(d, key)
	}
	return nil
}
