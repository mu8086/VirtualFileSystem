package main

import (
	"VirtualFileSystem/errors"
	"testing"
)

func Test_loadViper(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "local.toml",
			args:    args{configFile: "local.toml"},
			wantErr: false,
		},
		{
			name:    "not exists file",
			args:    args{configFile: "local2.toml"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadViper(tt.args.configFile); (err != nil) != tt.wantErr {
				t.Errorf("loadViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handleCommand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "empty line - 1",
			args:    args{input: "                "},
			wantErr: nil,
		},
		{
			name:    "empty line - 2",
			args:    args{input: ""},
			wantErr: nil,
		},
		{
			name:    "invalid command",
			args:    args{input: "test command"},
			wantErr: errors.ErrCmdNotExists,
		},
		{
			name:    "register user",
			args:    args{input: "register a"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handleCommand(tt.args.input); err != tt.wantErr {
				t.Errorf("handleCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
