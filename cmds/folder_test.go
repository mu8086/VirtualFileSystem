package cmds

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"testing"
)

func TestFolderCreate_Execute(t *testing.T) {
	createUserCmd := Get(constants.UserCreateCmd)
	createUserCmd.Execute([]string{"folderCreateUserA"})

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
				[]string{"", "folder1"},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{"folderCreateUserA", ""},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "normal",
			args: args{
				[]string{"folderCreateUserA", "folder1", "description"},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{"folderCreateUserB", "folder1"},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder exists",
			args: args{
				[]string{"folderCreateUserA", "folder1"},
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
	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{"folderListUserA"})
	userCreateCmd.Execute([]string{"folderListUserB"})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{"folderListUserA", "folder2", "descritpion2"})
	folderCreateCmd.Execute([]string{"folderListUserA", "folder1", "descritpion1"})

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
				[]string{""},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "normal with default sortOption, sortFlag",
			args: args{
				[]string{"folderListUserA"},
			},
			wantErr: nil,
		},
		{
			name: "normal",
			args: args{
				[]string{"folderListUserA", "--sort-name", "desc"},
			},
			wantErr: nil,
		},
		{
			name: "user have no folder",
			args: args{
				[]string{"folderListUserB"},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{"folderListUserC"},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "sort option invalid",
			args: args{
				[]string{"folderListUserA", "", "asc"},
			},
			wantErr: errors.ErrSortOption,
		},
		{
			name: "sort flag invalid",
			args: args{
				[]string{"folderListUserA", "--sort-name", ""},
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
	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{"folderRemoveUserA"})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{"folderRemoveUserA", "folder1", "descritpion1"})

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
				[]string{"", "folder1"},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{"folderRemoveUserA", ""},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "normal",
			args: args{
				[]string{"folderRemoveUserA", "folder1"},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{"folderRemoveUserB", "folder1"},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{"folderRemoveUserA", "folder2"},
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
	userCreateCmd := Get(constants.UserCreateCmd)
	userCreateCmd.Execute([]string{"folderRenameUserA"})

	folderCreateCmd := Get(constants.FolderCreateCmd)
	folderCreateCmd.Execute([]string{"folderRenameUserA", "folder1"})
	folderCreateCmd.Execute([]string{"folderRenameUserA", "folder2"})

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
				[]string{"", "folder1", "folder3"},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "foldername invalid",
			args: args{
				[]string{"folderRenameUserA", "", "folder3"},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "new foldername invalid",
			args: args{
				[]string{"folderRenameUserA", "folder1", ""},
			},
			wantErr: errors.ErrFolderName,
		},
		{
			name: "normal",
			args: args{
				[]string{"folderRenameUserA", "folder1", "folder3"},
			},
			wantErr: nil,
		},
		{
			name: "user not exists",
			args: args{
				[]string{"folderRenameUserB", "folder1", "folder3"},
			},
			wantErr: errors.ErrUserNotExists,
		},
		{
			name: "folder not exists",
			args: args{
				[]string{"folderRenameUserA", "folder1", "folder3"},
			},
			wantErr: errors.ErrFolderNotExists,
		},
		{
			name: "new folder exists",
			args: args{
				[]string{"folderRenameUserA", "folder2", "folder3"},
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
