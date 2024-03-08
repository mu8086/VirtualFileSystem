package cmds

import (
	"VirtualFileSystem/constants"
	"fmt"
)

func init() {
	register(UserCreate{})
}

type UserCreate struct{}

func (cmd UserCreate) Execute() {
	return
}

func (cmd UserCreate) Name() string {
	return constants.UserCreateCmd
}

func (cmd UserCreate) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd UserCreate) validate() bool {
	return true
}
