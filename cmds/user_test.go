package cmds

import (
	"VirtualFileSystem/errors"
	"testing"
)

func TestUserCreate_Execute(t *testing.T) {
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
				[]string{""},
			},
			wantErr: errors.ErrUserName,
		},
		{
			name: "normal",
			args: args{
				[]string{"userA"},
			},
			wantErr: nil,
		},
		{
			name: "user exists",
			args: args{
				[]string{"userA"},
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
