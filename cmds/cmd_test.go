package cmds

import (
	"VirtualFileSystem/constants"
	"fmt"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		cmdName string
	}
	tests := []struct {
		name string
		args args
		want Cmd
	}{
		{
			name: "empty",
			args: args{
				cmdName: "",
			},
			want: nil,
		},
		{
			name: constants.FolderCreateCmd,
			args: args{
				cmdName: constants.FolderCreateCmd,
			},
			want: cmds[constants.FolderCreateCmd],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.cmdName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvailableCmds(t *testing.T) {
	tmpCmds, tmpCmdsStr := cmds, cmdsStr

	cmds = make(map[string]Cmd)
	cmds[constants.FolderCreateCmd] = FolderCreate{}
	cmdsStr = ""

	tests := []struct {
		name string
		want string
	}{
		{
			name: constants.FolderCreateCmd,
			want: fmt.Sprintf("Available Commands: [%v]", constants.FolderCreateCmd),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AvailableCmds(); got != tt.want {
				t.Errorf("AvailableCmds() = %v, want %v", got, tt.want)
			}
		})
	}

	cmds, cmdsStr = tmpCmds, tmpCmdsStr
}

func Test_register(t *testing.T) {
	var zeroCmd Cmd

	type args struct {
		cmd Cmd
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "zeroCmd",
			args: args{
				cmd: zeroCmd,
			},
			want: false,
		},
		{
			name: "exists command",
			args: args{
				cmd: FileCreate{},
			},
			want: false,
		},
		{
			name: "mock command",
			args: args{
				cmd: Mock{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := register(tt.args.cmd); got != tt.want {
				t.Errorf("register() = %v, want %v", got, tt.want)
			}
		})
	}
}
