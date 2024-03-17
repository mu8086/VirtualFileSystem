package dao

import (
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
)

var users map[string]*dto.User

func init() {
	users = make(map[string]*dto.User)
}

func CreateUser(name string) error {
	if GetUser(name) != nil {
		return errors.ErrUserExists
	}

	users[name] = &dto.User{
		Name: name,
	}
	return nil
}

func GetUser(name string) *dto.User {
	if user, exists := users[name]; exists {
		return user
	}
	return nil
}
