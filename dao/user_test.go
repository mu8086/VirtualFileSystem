package dao

import (
	"VirtualFileSystem/dto"
	"VirtualFileSystem/errors"
	"reflect"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("getUserA")
	userNameB := strings.ToLower("getUserB")
	userA := &dto.User{Name: userNameA}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "normal",
			args: args{
				name: userNameB,
			},
			wantErr: nil,
		},
		{
			name: "exists user",
			args: args{
				name: userNameA,
			},
			wantErr: errors.ErrUserExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.name); err != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tmp := users
	defer func() {
		users = tmp
	}()

	userNameA := strings.ToLower("getUserA")
	userNameB := strings.ToLower("getUserB")
	userA := &dto.User{Name: userNameA}

	users = make(map[string]*dto.User)
	users[userNameA] = userA

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *dto.User
	}{
		{
			name: "normal - exists user",
			args: args{
				name: userNameA,
			},
			want: userA,
		},
		{
			name: "not exists user",
			args: args{
				name: userNameB,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUser(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
