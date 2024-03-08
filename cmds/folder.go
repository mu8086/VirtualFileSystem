package cmds

import (
	"VirtualFileSystem/constants"
	"fmt"
)

func init() {
	register(FolderCreate{})
}

type FolderCreate struct{}

func (cmd FolderCreate) Execute(args []string) error {
	return nil
}

func (cmd FolderCreate) Name() string {
	return constants.FolderCreateCmd
}

func (cmd FolderCreate) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FolderCreate) validate(args []string) error {
	return nil
}
