package dto

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"fmt"
	"sort"
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

func (f Folder) Get(fileName string) *File {
	for _, file := range f.Files {
		if file.Name == fileName {
			return file
		}
	}
	return nil
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

func (folders Folders) Sort(sortOption, sortFlag string) (Folders, error) {
	if len(folders) == 0 {
		return nil, nil
	}

	sorted := append(Folders{}, folders...)

	switch sortOption {
	case constants.OptionSortByName:
		if sortFlag == constants.FlagSortAsc {
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].Name < sorted[j].Name
			})
			return sorted, nil
		} else if sortFlag == constants.FlagSortDesc {
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].Name > sorted[j].Name
			})
			return sorted, nil
		}
		return nil, errors.ErrSortFlag

	case constants.OptionSortByCreated:
		if sortFlag == constants.FlagSortAsc {
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].CreatedAt.Before(sorted[j].CreatedAt)
			})
			return sorted, nil
		} else if sortFlag == constants.FlagSortDesc {
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].CreatedAt.After(sorted[j].CreatedAt)
			})
			return sorted, nil
		}
		return nil, errors.ErrSortFlag

	default:
		return nil, errors.ErrSortOption
	}
}

func (folders Folders) Remove(folderName string) (Folders, error) {
	for idx, folder := range folders {
		if folder.Name == folderName {
			return append(folders[:idx], folders[idx+1:]...), nil
		}
	}
	return nil, errors.ErrFolderNotExists
}

type File struct {
	CreatedAt   time.Time
	Description string
	Name        string
}

func (f File) String() string {
	return fmt.Sprintf("[File: %v, Desc: %v, createdAt: %v]", f.Name, f.Description, f.CreatedAt)
}
