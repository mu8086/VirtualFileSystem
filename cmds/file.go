package cmds

import (
	"VirtualFileSystem/constants"
	"fmt"
)

func init() {
	register(FileCreate{})
}

type FileCreate struct{}

func (cmd FileCreate) Execute() {
	return
}

func (cmd FileCreate) Name() string {
	return constants.FileCreateCmd
}

func (cmd FileCreate) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FileCreate) validate() bool {
	return true
}
