package dao

import (
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
	"time"
)

func CreateFolder(userName, folderName, desc string) error {
	user := GetUser(userName)
	if user == nil {
		return errors.ErrUserNotExists
	}

	folder := user.Folders.Get(folderName)
	if folder != nil {
		return errors.ErrFolderExists
	}

	user.Folders = append(user.Folders, &dto.Folder{
		CreatedAt:   time.Now(),
		Name:        folderName,
		Description: desc,
		Files:       nil,
	})
	return nil
}

func RemoveFolder(userName, folderName string) error {
	user := GetUser(userName)
	if user == nil {
		return errors.ErrUserNotExists
	}

	folders, err := user.Folders.Remove(folderName)
	if err != nil {
		return err
	}

	user.Folders = folders
	return nil
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
