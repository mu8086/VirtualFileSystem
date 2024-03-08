package dto

import (
	"fmt"
	"time"
)

type User struct {
	Name    string
	Folders []*Folder
}

func (u User) String() string {
	return fmt.Sprintf("[User: %v:[%v]]", u.Name, u.Folders)
}

type Folder struct {
	createdAt time.Time
	Name      string
	Files     []*File
}

func (f Folder) String() string {
	return fmt.Sprintf("[Folder: %v, createdAt: %v, Files: %v]", f.Name, f.createdAt, f.Files)
}

type File struct {
	createdAt time.Time
	Name      string
}

func (f File) String() string {
	return fmt.Sprintf("[File: %v, createdAt: %v]", f.Name, f.createdAt)
}
