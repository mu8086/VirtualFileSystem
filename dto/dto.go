package dto

import (
	"fmt"
	"time"
)

type User struct {
	Name    string
	Folders Folders
}

func (u User) String() string {
	return fmt.Sprintf("[User: %v, Folders: %v]", u.Name, u.Folders)
}

type Folder struct {
	CreatedAt   time.Time
	Description string
	Files       []*File
	Name        string
}

func (f Folder) String() string {
	return fmt.Sprintf("[Folder: %v, Desc: %v, CreatedAt: %v, Files: %v]", f.Name, f.Description, f.CreatedAt, f.Files)
}

type Folders []*Folder

func (folders Folders) Get(folderName string) *Folder {
	for _, folder := range folders {
		if folder.Name == folderName {
			return folder
		}
	}
	return nil
}

type File struct {
	createdAt   time.Time
	Description string
	Name        string
}

func (f File) String() string {
	return fmt.Sprintf("[File: %v, createdAt: %v]", f.Name, f.createdAt)
}
