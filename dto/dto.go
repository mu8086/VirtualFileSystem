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

type Folder struct {
	CreatedAt   time.Time
	Description string
	Files       Files
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

func (folders Folders) String() (s string) {
	for _, folder := range folders {
		s += folder.Name

		if len(folder.Description) != 0 {
			s += " " + folder.Description
		}

		s += fmt.Sprintf(" %v\n", folder.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return s
}

func (folders Folders) Remove(folderName string) (Folders, error) {
	for idx, folder := range folders {
		if folder.Name == folderName {
			folder.Files = nil
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

type Files []*File

func (files Files) Sort(sortOption, sortFlag string) (Files, error) {
	if len(files) == 0 {
		return nil, nil
	}

	sorted := append(Files{}, files...)

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

func (files Files) String() (s string) {
	for _, file := range files {
		s += file.Name

		if len(file.Description) != 0 {
			s += " " + file.Description
		}

		s += fmt.Sprintf(" %v\n", file.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return s
}

func (files Files) Remove(fileName string) (Files, error) {
	for idx, file := range files {
		if file.Name == fileName {
			return append(files[:idx], files[idx+1:]...), nil
		}
	}
	return nil, errors.ErrFileNotExists
}
