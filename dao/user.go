package dao

import (
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
	"fmt"
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

// TODO: remove
func GetAllUsers() string {
	s := ""
	for _, user := range users {
		s += fmt.Sprintf(" %v", user)
	}
	if len(s) != 0 {
		s = s[1:]
	}
	return s
}

func GetUser(name string) *dto.User {
	if user, exists := users[name]; exists {
		return user
	}
	return nil
}

func GetUserFolder(userName, folderName string) *dto.Folder {
	if user := GetUser(userName); user != nil {
		return user.Folders.Get(folderName)
	}
	return nil
}
