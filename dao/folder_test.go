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

func TestCreateFolder(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("getUserA")
	userNameB := strings.ToLower("getUserB")
	folder1 := strings.ToLower("folder1")
	folder2 := strings.ToLower("folder2")

	userA := &dto.User{Name: userNameA, Folders: dto.Folders{&dto.Folder{Name: folder1}}}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		userName   string
		folderName string
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
				folderName: folder1,
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder exists",
			args: args{
				userName:   userNameA,
				folderName: folder1,
			},
			wantErr: errors.ErrFolderExists,
		},
		{
			name: "normal",
			args: args{
				userName:   userNameA,
				folderName: folder2,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFolder(tt.args.userName, tt.args.folderName, tt.args.desc); err != tt.wantErr {
				t.Errorf("CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFolders(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	timestamp1 := time.Now()
	timestamp2 := timestamp1.Add(10 * time.Second)
	timestamp3 := timestamp1.Add(20 * time.Second)

	userNameA := strings.ToLower("getUserA")
	userNameB := strings.ToLower("getUserB")
	userNameC := strings.ToLower("getUserC")

	folder1 := &dto.Folder{Name: "folder1", CreatedAt: timestamp2}
	folder2 := &dto.Folder{Name: "folder2", CreatedAt: timestamp3}
	folder3 := &dto.Folder{Name: "folder3", CreatedAt: timestamp1}

	sortByCreatedAscFolders := dto.Folders{folder3, folder1, folder2}
	sortByCreatedDescFolders := dto.Folders{folder2, folder1, folder3}
	sortByNameAscFolders := dto.Folders{folder1, folder2, folder3}
	sortByNameDescFolders := dto.Folders{folder3, folder2, folder1}

	userA := &dto.User{
		Name:    userNameA,
		Folders: dto.Folders{folder2, folder3, folder1},
	}
	userB := &dto.User{
		Name: userNameB,
	}

	users = make(map[string]*dto.User)
	users[userNameA] = userA
	users[userNameB] = userB

	type args struct {
		userName   string
		sortOption string
		sortFlag   string
	}
	tests := []struct {
		name    string
		args    args
		want    dto.Folders
		wantErr error
	}{
		{
			name: "user not exists",
			args: args{
				userName:   userNameC,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    nil,
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "invalid sortOption",
			args: args{
				userName:   userNameA,
				sortOption: constants.OptionSortByCreated + "+",
				sortFlag:   constants.FlagSortAsc,
			},
			want:    nil,
			wantErr: errors.ErrSortOption,
		},
		{
			name: "invalid sortFlag",
			args: args{
				userName:   userNameA,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc + "+",
			},
			want:    nil,
			wantErr: errors.ErrSortFlag,
		},
		{
			name: "sort by created (asc)",
			args: args{
				userName:   userNameA,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByCreatedAscFolders,
			wantErr: nil,
		},
		{
			name: "sort by created (desc)",
			args: args{
				userName:   userNameA,
				sortOption: constants.OptionSortByCreated,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByCreatedDescFolders,
			wantErr: nil,
		},
		{
			name: "sort by name (asc)",
			args: args{
				userName:   userNameA,
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortAsc,
			},
			want:    sortByNameAscFolders,
			wantErr: nil,
		},
		{
			name: "sort by name (desc)",
			args: args{
				userName:   userNameA,
				sortOption: constants.OptionSortByName,
				sortFlag:   constants.FlagSortDesc,
			},
			want:    sortByNameDescFolders,
			wantErr: nil,
		},
		{
			name: "user have no folder",
			args: args{
				userName: userNameB,
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFolders(tt.args.userName, tt.args.sortOption, tt.args.sortFlag)
			if err != tt.wantErr {
				t.Errorf("GetFolders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFolders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFolder(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("getUserA")
	userNameB := strings.ToLower("getUserB")

	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")

	folder1 := &dto.Folder{Name: folderName1}

	userA := &dto.User{
		Name:    userNameA,
		Folders: dto.Folders{folder1},
	}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		userName   string
		folderName string
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
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				userName:   userNameA,
				folderName: folderName2,
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "normal",
			args: args{
				userName:   userNameA,
				folderName: folderName1,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFolder(tt.args.userName, tt.args.folderName); err != tt.wantErr {
				t.Errorf("RemoveFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateFolder(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("getUserA")
	userNameB := strings.ToLower("getUserB")

	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")

	folder1 := &dto.Folder{Name: folderName1}
	folder2 := &dto.Folder{Name: folderName2}

	userA := &dto.User{
		Name:    userNameA,
		Folders: dto.Folders{folder1, folder2},
	}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		userName      string
		folderName    string
		newFolderName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "user not exists",
			args: args{
				userName:      userNameB,
				folderName:    folderName1,
				newFolderName: folderName2,
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				userName:      userNameA,
				folderName:    folderName3,
				newFolderName: folderName2,
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "new folder exists",
			args: args{
				userName:      userNameA,
				folderName:    folderName1,
				newFolderName: folderName2,
			},
			wantErr: errors.ErrFolderExists,
		},
		{
			name: "normal",
			args: args{
				userName:      userNameA,
				folderName:    folderName1,
				newFolderName: folderName3,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateFolder(tt.args.userName, tt.args.folderName, tt.args.newFolderName); err != tt.wantErr {
				t.Errorf("UpdateFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
