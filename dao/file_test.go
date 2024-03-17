package dao

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCreateFile(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("CreateFileUserA")
	userNameB := strings.ToLower("CreateFileUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")

	file1 := &dto.File{Name: fileName1}
	folder1 := &dto.Folder{Name: folderName1, Files: dto.Files{file1}}
	userA := &dto.User{Name: userNameA, Folders: dto.Folders{folder1}}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		userName   string
		folderName string
		fileName   string
		desc       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "user not exists",
			args: args{
				userName:   userNameB,
				folderName: folderName1,
				fileName:   fileName1,
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				userName:   userNameA,
				folderName: folderName2,
				fileName:   fileName1,
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "file exists",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				fileName:   fileName1,
			},
			wantErr: errors.ErrFileExists,
		},
		{
			name: "normal",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				fileName:   fileName2,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFile(tt.args.userName, tt.args.folderName, tt.args.fileName, tt.args.desc); err != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFiles(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("CreateFileUserA")
	userNameB := strings.ToLower("CreateFileUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")
	fileName3 := strings.ToLower("file3")
	timestamp1 := time.Now()
	timestamp2 := timestamp1.Add(10 * time.Second)
	timestamp3 := timestamp1.Add(20 * time.Second)

	file1 := &dto.File{Name: fileName1, CreatedAt: timestamp3}
	file2 := &dto.File{Name: fileName2, CreatedAt: timestamp1}
	file3 := &dto.File{Name: fileName3, CreatedAt: timestamp2}

	sortByCreatedAscFiles := dto.Files{file2, file3, file1}
	sortByCreatedDescFiles := dto.Files{file1, file3, file2}
	sortByNameAscFiles := dto.Files{file1, file2, file3}
	sortByNameDescFiles := dto.Files{file3, file2, file1}

	files := dto.Files{file3, file1, file2}
	folder1 := &dto.Folder{Name: folderName1, Files: files}
	userA := &dto.User{Name: userNameA, Folders: dto.Folders{folder1}}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		userName   string
		folderName string
		sortOption string
		sortFlag   string
	}
	tests := []struct {
		name    string
		args    args
		want    dto.Files
		wantErr error
	}{
		{
			name: "user not exists",
			args: args{
				userName:   userNameB,
				folderName: folderName1,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    nil,
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				userName:   userNameA,
				folderName: folderName2,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    nil,
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "invalid sort option",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				sortOption: constants.OptionSortByCreated + "+",
				sortFlag:   constants.FlagSortAsc,
			},
			want:    nil,
			wantErr: errors.ErrSortOption,
		},
		{
			name: "invalid sort flag",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc + "+",
			},
			want:    nil,
			wantErr: errors.ErrSortFlag,
		},
		{
			name: "sort by created asc",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByCreatedAscFiles,
			wantErr: nil,
		},
		{
			name: "sort by created desc",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByCreatedDescFiles,
			wantErr: nil,
		},
		{
			name: "sort by name asc",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByNameAscFiles,
			wantErr: nil,
		},
		{
			name: "sort by name desc",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByNameDescFiles,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFiles(tt.args.userName, tt.args.folderName, tt.args.sortOption, tt.args.sortFlag)
			if err != tt.wantErr {
				t.Errorf("GetFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFile(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("CreateFileUserA")
	userNameB := strings.ToLower("CreateFileUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")

	file1 := &dto.File{Name: fileName1}

	files := dto.Files{file1}
	folder1 := &dto.Folder{Name: folderName1, Files: files}
	userA := &dto.User{Name: userNameA, Folders: dto.Folders{folder1}}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		userName   string
		folderName string
		fileName   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "user not exists",
			args: args{
				userName:   userNameB,
				folderName: folderName1,
				fileName:   fileName1,
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				userName:   userNameA,
				folderName: folderName2,
				fileName:   fileName1,
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "file not exists",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				fileName:   fileName2,
			},
			wantErr: errors.ErrFileNotExists,
		},
		{
			name: "normal",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
				fileName:   fileName1,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFile(tt.args.userName, tt.args.folderName, tt.args.fileName); err != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
