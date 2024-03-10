package dao

import (
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
	"time"
)

func CreateFile(userName, folderName, fileName, desc string) error {
	user := GetUser(userName)
	if user == nil {
		return errors.ErrUserNotExists
	}

	folder := user.Folders.Get(folderName)
	if folder == nil {
		return errors.ErrFolderNotExists
	}

	file := folder.Get(fileName)
	if file != nil {
		return errors.ErrFileExists
	}

	folder.Files = append(folder.Files, &dto.File{
		CreatedAt:   time.Now(),
		Description: desc,
		Name:        fileName,
	})
	return nil
}

func RemoveFile(userName, folderName, fileName string) error {
	user := GetUser(userName)
	if user == nil {
		return errors.ErrUserNotExists
	}

	folder := user.Folders.Get(folderName)
	if folder == nil {
		return errors.ErrFolderNotExists
	}

	files, err := folder.Files.Remove(fileName)
	if err != nil {
		return err
	}

	folder.Files = files
	return nil
}
