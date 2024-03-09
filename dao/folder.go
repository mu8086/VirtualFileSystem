package dao

import (
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
	"time"
)

func CreateFolder(userName, folderName, desc string) bool {
	user := GetUser(userName)
	if user == nil {
		return false
	}

	user.Folders = append(user.Folders, &dto.Folder{
		CreatedAt:   time.Now(),
		Name:        folderName,
		Description: desc,
		Files:       nil,
	})

	return true
}

func UpdateFolder(userName, folderName, newFolderName string) error {
	user := GetUser(userName)
	if user == nil {
		return errors.ErrUserNotExists
	}

	folder := user.Folders.Get(folderName)
	if folder == nil {
		return errors.ErrFolderNotExists
	} else if user.Folders.Get(newFolderName) != nil {
		return errors.ErrFolderExists
	}

	folder.Name = newFolderName
	return nil
}
