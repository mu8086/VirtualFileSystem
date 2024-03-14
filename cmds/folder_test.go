package cmds

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"strings"
	"testing"
)

func TestFolderCreate_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("folderCreateUserA")
	userNameB := strings.ToLower("folderCreateUserB")
	folderName := strings.ToLower("folder1")

	createUserCmd := Get(constants.UserCreateCmd)
	createUserCmd.Execute([]string{userNameA})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FolderCreate
		args    args
		wantErr error
	}{
		{
			name: "args size invalid",
			args: args{
				[]string{},
			},
			wantErr: errors.ErrArgSize,
		},
		{
			name: "username invalid",
			args: args{
				[]string{emptyStr, folderName},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{userNameA, emptyStr},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, folderName, "description"},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameB, folderName},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder exists",
			args: args{
				[]string{userNameA, folderName},
			},
			wantErr: errors.ErrFolderExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FolderCreate.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFoldersList_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("folderListUserA")
	userNameB := strings.ToLower("folderListUserB")
	userNameC := strings.ToLower("folderListUserC")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")

	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{userNameA})
	userCreateCmd.Execute([]string{userNameB})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{userNameA, folderName2, "descritpion2"})
	folderCreateCmd.Execute([]string{userNameA, folderName1, "descritpion1"})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FoldersList
		args    args
		wantErr error
	}{
		{
			name: "args size invalid",
			args: args{
				[]string{},
			},
			wantErr: errors.ErrArgSize,
		},
		{
			name: "username invalid",
			args: args{
				[]string{emptyStr},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "normal with default sortOption, sortFlag",
			args: args{
				[]string{userNameA},
			},
			wantErr: nil,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, constants.OptionSortByName, constants.FlagSortAsc},
			},
			wantErr: nil,
		},
		{
			name: "user have no folder",
			args: args{
				[]string{userNameB},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameC},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "sort option invalid",
			args: args{
				[]string{userNameA, emptyStr, constants.FlagSortAsc},
			},
			wantErr: errors.ErrSortOption,
		},
		{
			name: "sort flag invalid",
			args: args{
				[]string{userNameA, constants.OptionSortByName, emptyStr},
			},
			wantErr: errors.ErrSortFlag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FoldersList.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFolderRemove_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("folderRemoveUserA")
	userNameB := strings.ToLower("folderRemoveUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")

	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{userNameA})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{userNameA, folderName1, "descritpion1"})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FolderRemove
		args    args
		wantErr error
	}{
		{
			name: "args size invalid",
			args: args{
				[]string{},
			},
			wantErr: errors.ErrArgSize,
		},
		{
			name: "username invalid",
			args: args{
				[]string{emptyStr, folderName1},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{userNameA, emptyStr},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, folderName1},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameB, folderName1},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{userNameA, folderName2},
			},
			wantErr: errors.ErrFolderNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FolderRemove.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFolderRename_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("folderRenameUserA")
	userNameB := strings.ToLower("folderRenameUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")

	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{userNameA})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{userNameA, folderName1})
	folderCreateCmd.Execute([]string{userNameA, folderName2})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FolderRename
		args    args
		wantErr error
	}{
		{
			name: "args size invalid",
			args: args{
				[]string{},
			},
			wantErr: errors.ErrArgSize,
		},
		{
			name: "username invalid",
			args: args{
				[]string{emptyStr, folderName1, folderName3},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{userNameA, emptyStr, folderName3},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "new foldername invalid",
			args: args{
				[]string{userNameA, folderName1, emptyStr},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, folderName1, folderName3},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameB, folderName1, folderName3},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{userNameA, folderName1, folderName3},
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "new folder exists",
			args: args{
				[]string{userNameA, folderName2, folderName3},
			},
			wantErr: errors.ErrFolderExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FolderRename.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
