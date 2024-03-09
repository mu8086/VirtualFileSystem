package dao

import (
	"VirtualFileSystem/dto"
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
