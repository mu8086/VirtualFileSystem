package cmds

import (
	"VirtualFileSystem/errors"
	"strings"
	"testing"
)

func TestUserCreate_Execute(t *testing.T) {
	emptyStr := ""
	userNameA := strings.ToLower("userCreateUserA")

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		cmd     UserCreate
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
			name: "normal",
			args: args{
				[]string{userNameA},
			},
			wantErr: nil,
		},
		{
			name: "user exists",
			args: args{
				[]string{userNameA},
			},
			wantErr: errors.ErrUserExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cmd.Execute(tt.args.args); err != tt.wantErr {
				t.Errorf("UserCreate.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
