package cmds

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"strings"
	"testing"
)

func TestFileCreate_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("fileCreateUserA")
	userNameB := strings.ToLower("fileCreateUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	fileName1 := strings.ToLower("file1")

	createUserCmd := Get(constants.UserCreateCmd)
	createUserCmd.Execute([]string{userNameA})

	createFolderCmd := Get(constants.FolderCreateCmd)
	createFolderCmd.Execute([]string{userNameA, folderName1})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FileCreate
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
				[]string{emptyStr, folderName1, fileName1},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{userNameA, emptyStr, fileName1},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "filename invalid",
			args: args{
				[]string{userNameA, folderName1, emptyStr},
			},
			wantErr: errors.ErrFileName,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, folderName1, fileName1, "description"},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameB, folderName1, fileName1},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{userNameA, folderName2, fileName1},
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "file exists",
			args: args{
				[]string{userNameA, folderName1, fileName1},
			},
			wantErr: errors.ErrFileExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FileCreate.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileRemove_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("fileRemoveUserA")
	userNameB := strings.ToLower("fileRemoveUserB")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")

	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{userNameA})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{userNameA, folderName1})

	fileCreateCmd := Get(constants.FileCreateCmd)
	fileCreateCmd.Execute([]string{userNameA, folderName1, fileName1})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FileRemove
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
				[]string{emptyStr, folderName1, fileName1},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{userNameA, emptyStr, fileName1},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "filename invalid",
			args: args{
				[]string{userNameA, folderName1, emptyStr},
			},
			wantErr: errors.ErrFileName,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, folderName1, fileName1},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameB, folderName1, folderName1},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{userNameA, folderName2, fileName1},
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "file not exists",
			args: args{
				[]string{userNameA, folderName1, fileName2},
			},
			wantErr: errors.ErrFileNotExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FileRemove.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFilesList_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("folderListUserA")
	userNameB := strings.ToLower("folderListUserB")
	userNameC := strings.ToLower("folderListUserC")
	folderName1 := strings.ToLower("folder1")
	folderName2 := strings.ToLower("folder2")
	folderName3 := strings.ToLower("folder3")
	fileName1 := strings.ToLower("file1")
	fileName2 := strings.ToLower("file2")

	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{userNameA})
	userCreateCmd.Execute([]string{userNameB})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{userNameA, folderName1})
	folderCreateCmd.Execute([]string{userNameA, folderName2})

	fileCreateCmd := Get(constants.FileCreateCmd)
	fileCreateCmd.Execute([]string{userNameA, folderName1, fileName1})
	fileCreateCmd.Execute([]string{userNameA, folderName1, fileName2})

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     FilesList
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
			name: "normal with default sortOption, sortFlag",
			args: args{
				[]string{userNameA, folderName1},
			},
			wantErr: nil,
		},
		{
			name: "normal with empty folder",
			args: args{
				[]string{userNameA, folderName2},
			},
			wantErr: nil,
		},
		{
			name: "normal",
			args: args{
				[]string{userNameA, folderName1, constants.OptionSortByName, constants.FlagSortDesc},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{userNameC, folderName1},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "user have no folder",
			args: args{
				[]string{userNameB, folderName1},
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{userNameA, folderName3},
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "sort option invalid",
			args: args{
				[]string{userNameA, folderName1, emptyStr, constants.FlagSortAsc},
			},
			wantErr: errors.ErrSortOption,
		},
		{
			name: "sort flag invalid",
			args: args{
				[]string{userNameA, folderName1, constants.OptionSortByName, emptyStr},
			},
			wantErr: errors.ErrSortFlag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("FilesList.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
