package dto

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestFolder_Get(t *testing.T) {
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")
	folderName1 := strings.ToLower("folder1")
	emptyFolderName := strings.ToLower("emptyFolder")

	file1 := &File{Name: fileName1}
	folder1 := Folder{Name: folderName1, Files: Files{file1}}
	emptyFolder := Folder{Name: emptyFolderName, Files: Files{}}

	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		f    Folder
		args args
		want *File
	}{
		{
			name: "normal",
			f:    folder1,
			args: args{fileName: fileName1},
			want: file1,
		},
		{
			name: "no such file in the folder",
			f:    folder1,
			args: args{fileName: fileName2},
			want: nil,
		},
		{
			name: "search file in a empty folder",
			f:    emptyFolder,
			args: args{fileName: fileName2},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Get(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Folder.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolders_Get(t *testing.T) {
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")

	folder1 := Folder{Name: folderName1}
	folder2 := Folder{Name: folderName2}

	folders1 := Folders{&folder1, &folder2}
	emptyFolders := Folders{}

	type args struct {
		folderName string
	}
	tests := []struct {
		name    string
		folders Folders
		args    args
		want    *Folder
	}{
		{
			name:    "normal - 1",
			folders: folders1,
			args:    args{folderName: folderName1},
			want:    &folder1,
		},
		{
			name:    "normal - 2",
			folders: folders1,
			args:    args{folderName: folderName2},
			want:    &folder2,
		},
		{
			name:    "no such folder in folders",
			folders: folders1,
			args:    args{folderName: folderName3},
			want:    nil,
		},
		{
			name:    "search folder in a empty folders",
			folders: emptyFolders,
			args:    args{folderName: folderName1},
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.folders.Get(tt.args.folderName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Folders.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolders_Sort(t *testing.T) {
	timestamp1 := time.Now()
	timestamp2 := timestamp1.Add(10 * time.Second)
	timestamp3 := timestamp1.Add(20 * time.Second)

	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")

	folder1 := &Folder{Name: folderName1, CreatedAt: timestamp3}
	folder2 := &Folder{Name: folderName2, CreatedAt: timestamp1}
	folder3 := &Folder{Name: folderName3, CreatedAt: timestamp2}

	folders1 := Folders{folder3, folder1, folder2}
	emptyFolders := Folders{}

	sortByCreatedAscFolders := Folders{folder2, folder3, folder1}
	sortByCreatedDescFolders := Folders{folder1, folder3, folder2}
	sortByNameAscFolders := Folders{folder1, folder2, folder3}
	sortByNameDescFolders := Folders{folder3, folder2, folder1}

	type args struct {
		sortOption string
		sortFlag   string
	}
	tests := []struct {
		name    string
		folders Folders
		args    args
		want    Folders
		wantErr error
	}{
		{
			name:    "empty folders",
			folders: emptyFolders,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:    "sort by created asc",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByCreatedAscFolders,
			wantErr: nil,
		},
		{
			name:    "sort by created desc",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByCreatedDescFolders,
			wantErr: nil,
		},
		{
			name:    "sort by name asc",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByNameAscFolders,
			wantErr: nil,
		},
		{
			name:    "sort by name desc",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByNameDescFolders,
			wantErr: nil,
		},
		{
			name:    "invalid sort option",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByName + "+",
				sortFlag:   constants.FlagSortDesc,
			},
			want:    nil,
			wantErr: errors.ErrSortOption,
		},
		{
			name:    "sort by created with invalid sort flag",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortDesc + "+",
			},
			want:    nil,
			wantErr: errors.ErrSortFlag,
		},
		{
			name:    "sort by name with invalid sort flag",
			folders: folders1,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc + "+",
			},
			want:    nil,
			wantErr: errors.ErrSortFlag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.folders.Sort(tt.args.sortOption, tt.args.sortFlag)
			if err != tt.wantErr {
				t.Errorf("Folders.Sort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Folders.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolders_String(t *testing.T) {
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")
	folderName4 := strings.ToLower("folder4")

	timestamp1 := time.Now()
	timestamp2 := timestamp1.Add(10 * time.Second)
	timestamp3 := timestamp1.Add(20 * time.Second)
	timestamp4 := timestamp1.Add(30 * time.Second)

	desc1 := "desc1"
	desc2 := "desc2"
	desc3 := ""
	desc4 := "desc4"

	folder1 := &Folder{Name: folderName1, CreatedAt: timestamp1, Description: desc1}
	folder2 := &Folder{Name: folderName2, CreatedAt: timestamp2, Description: desc2}
	folder3 := &Folder{Name: folderName3, CreatedAt: timestamp3, Description: desc3}
	folder4 := &Folder{Name: folderName4, CreatedAt: timestamp4, Description: desc4}

	s1 := fmt.Sprintf("%v %v %v", folder1.Name, folder1.Description, folder1.CreatedAt.Format("2006-01-02 15:04:05"))
	s2 := fmt.Sprintf("%v %v %v", folder2.Name, folder2.Description, folder2.CreatedAt.Format("2006-01-02 15:04:05"))
	s3 := fmt.Sprintf("%v %v", folder3.Name, folder3.CreatedAt.Format("2006-01-02 15:04:05"))
	s4 := fmt.Sprintf("%v %v %v", folder4.Name, folder4.Description, folder4.CreatedAt.Format("2006-01-02 15:04:05"))

	emptyFolders := Folders{}
	folders1 := Folders{folder1, folder4, folder2, folder3}
	folders2 := Folders{folder3, folder1, folder4, folder2}

	wantS1 := fmt.Sprintf("%v\n%v\n%v\n%v\n", s1, s4, s2, s3)
	wantS2 := fmt.Sprintf("%v\n%v\n%v\n%v\n", s3, s1, s4, s2)

	tests := []struct {
		name    string
		folders Folders
		wantS   string
	}{
		{
			name:    "empty folders",
			folders: emptyFolders,
			wantS:   "",
		},
		{
			name:    "folders1",
			folders: folders1,
			wantS:   wantS1,
		},
		{
			name:    "folders2",
			folders: folders2,
			wantS:   wantS2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.folders.String(); gotS != tt.wantS {
				t.Errorf("Folders.String() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestFolders_Remove(t *testing.T) {
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")
	folderName4 := strings.ToLower("folder4")

	folder1 := &Folder{Name: folderName1}
	folder2 := &Folder{Name: folderName2}
	folder3 := &Folder{Name: folderName3}

	emptyFolders := Folders{}
	folders1 := Folders{folder2, folder3, folder1}
	folders2 := Folders{folder2, folder3, folder1}
	folders3 := Folders{folder2, folder3, folder1}
	folders4 := Folders{folder2, folder3, folder1}

	foldersRemove1 := Folders{folder2, folder3}
	foldersRemove2 := Folders{folder3, folder1}
	foldersRemove3 := Folders{folder2, folder1}

	type args struct {
		folderName string
	}
	tests := []struct {
		name    string
		folders Folders
		args    args
		want    Folders
		wantErr error
	}{
		{
			name:    "remove 1",
			folders: folders1,
			args:    args{folderName: folderName1},
			want:    foldersRemove1,
			wantErr: nil,
		},
		{
			name:    "remove 2",
			folders: folders2,
			args:    args{folderName: folderName2},
			want:    foldersRemove2,
			wantErr: nil,
		},
		{
			name:    "remove 3",
			folders: folders3,
			args:    args{folderName: folderName3},
			want:    foldersRemove3,
			wantErr: nil,
		},
		{
			name:    "remove a not exists folder",
			folders: folders4,
			args:    args{folderName: folderName4},
			want:    nil,
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name:    "remove from a empty folder",
			folders: emptyFolders,
			args:    args{folderName: folderName1},
			want:    nil,
			wantErr: errors.ErrFolderNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.folders.Remove(tt.args.folderName)
			if err != tt.wantErr {
				t.Errorf("[%v]Folders.Remove() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[%v]Folders.Remove() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFiles_Sort(t *testing.T) {
	timestamp1 := time.Now()
	timestamp2 := timestamp1.Add(10 * time.Second)
	timestamp3 := timestamp1.Add(20 * time.Second)

	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")
	fileName3 := strings.ToLower("file3")

	file1 := &File{Name: fileName1, CreatedAt: timestamp3}
	file2 := &File{Name: fileName2, CreatedAt: timestamp1}
	file3 := &File{Name: fileName3, CreatedAt: timestamp2}

	files1 := Files{file2, file1, file3}
	emptyFiles := Files{}

	sortByCreatedAscFiles := Files{file2, file3, file1}
	sortByCreatedDescFiles := Files{file1, file3, file2}
	sortByNameAscFiles := Files{file1, file2, file3}
	sortByNameDescFiles := Files{file3, file2, file1}

	type args struct {
		sortOption string
		sortFlag   string
	}
	tests := []struct {
		name    string
		files   Files
		args    args
		want    Files
		wantErr error
	}{
		{
			name:  "sort by created asc",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByCreatedAscFiles,
			wantErr: nil,
		},
		{
			name:  "sort by created desc",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByCreatedDescFiles,
			wantErr: nil,
		},
		{
			name:  "sort by name asc",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByNameAscFiles,
			wantErr: nil,
		},
		{
			name:  "sort by name desc",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByNameDescFiles,
			wantErr: nil,
		},
		{
			name:  "invalid sort option",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByName + "+",
				sortFlag:   constants.FlagSortDesc,
			},
			want:    nil,
			wantErr: errors.ErrSortOption,
		},
		{
			name:  "sort by created with invalid sort flag",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortDesc + "+",
			},
			want:    nil,
			wantErr: errors.ErrSortFlag,
		},
		{
			name:  "sort by name with invalid sort flag",
			files: files1,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc + "+",
			},
			want:    nil,
			wantErr: errors.ErrSortFlag,
		},
		{
			name:  "empty files",
			files: emptyFiles,
			args: args{
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.files.Sort(tt.args.sortOption, tt.args.sortFlag)
			if err != tt.wantErr {
				t.Errorf("Files.Sort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Files.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiles_String(t *testing.T) {
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")
	fileName3 := strings.ToLower("file3")
	fileName4 := strings.ToLower("file4")

	timestamp1 := time.Now()
	timestamp2 := timestamp1.Add(10 * time.Second)
	timestamp3 := timestamp1.Add(20 * time.Second)
	timestamp4 := timestamp1.Add(30 * time.Second)

	desc1 := "desc1"
	desc2 := "desc2"
	desc3 := ""
	desc4 := "desc4"

	file1 := &File{Name: fileName1, CreatedAt: timestamp1, Description: desc1}
	file2 := &File{Name: fileName2, CreatedAt: timestamp2, Description: desc2}
	file3 := &File{Name: fileName3, CreatedAt: timestamp3, Description: desc3}
	file4 := &File{Name: fileName4, CreatedAt: timestamp4, Description: desc4}

	s1 := fmt.Sprintf("%v %v %v", file1.Name, file1.Description, file1.CreatedAt.Format("2006-01-02 15:04:05"))
	s2 := fmt.Sprintf("%v %v %v", file2.Name, file2.Description, file2.CreatedAt.Format("2006-01-02 15:04:05"))
	s3 := fmt.Sprintf("%v %v", file3.Name, file3.CreatedAt.Format("2006-01-02 15:04:05"))
	s4 := fmt.Sprintf("%v %v %v", file4.Name, file4.Description, file4.CreatedAt.Format("2006-01-02 15:04:05"))

	emptyFiles := Files{}
	files1 := Files{file1, file4, file2, file3}
	files2 := Files{file3, file1, file4, file2}

	wantS1 := fmt.Sprintf("%v\n%v\n%v\n%v\n", s1, s4, s2, s3)
	wantS2 := fmt.Sprintf("%v\n%v\n%v\n%v\n", s3, s1, s4, s2)

	tests := []struct {
		name  string
		files Files
		wantS string
	}{
		{
			name:  "empty files",
			files: emptyFiles,
			wantS: "",
		},
		{
			name:  "files1",
			files: files1,
			wantS: wantS1,
		},
		{
			name:  "files2",
			files: files2,
			wantS: wantS2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.files.String(); gotS != tt.wantS {
				t.Errorf("Files.String() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestFiles_Remove(t *testing.T) {
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")
	fileName3 := strings.ToLower("file3")
	fileName4 := strings.ToLower("file4")

	file1 := &File{Name: fileName1}
	file2 := &File{Name: fileName2}
	file3 := &File{Name: fileName3}

	emptyFiles := Files{}
	files1 := Files{file2, file3, file1}
	files2 := Files{file2, file3, file1}
	files3 := Files{file2, file3, file1}
	files4 := Files{file2, file3, file1}

	filesRemove1 := Files{file2, file3}
	filesRemove2 := Files{file3, file1}
	filesRemove3 := Files{file2, file1}

	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		files   Files
		args    args
		want    Files
		wantErr error
	}{
		{
			name:    "remove 1",
			files:   files1,
			args:    args{fileName: fileName1},
			want:    filesRemove1,
			wantErr: nil,
		},
		{
			name:    "remove 2",
			files:   files2,
			args:    args{fileName: fileName2},
			want:    filesRemove2,
			wantErr: nil,
		},
		{
			name:    "remove 3",
			files:   files3,
			args:    args{fileName: fileName3},
			want:    filesRemove3,
			wantErr: nil,
		},
		{
			name:    "remove a not exists file",
			files:   files4,
			args:    args{fileName: fileName4},
			want:    nil,
			wantErr: errors.ErrFileNotExists,
		},
		{
			name:    "remove from a empty files",
			files:   emptyFiles,
			args:    args{fileName: fileName1},
			want:    nil,
			wantErr: errors.ErrFileNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.files.Remove(tt.args.fileName)
			if err != tt.wantErr {
				t.Errorf("Files.Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Files.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
