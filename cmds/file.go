package cmds

import (
	"VirtualFileSystem/constants"
	"fmt"
	"os"
)

func init() {
	register(FileCreate{})
}

type FileCreate struct{}

func (cmd FileCreate) Execute(args []string) error {
	return nil
}

func (cmd FileCreate) Name() string {
	return constants.FileCreateCmd
}

func (cmd FileCreate) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FileCreate) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername] [filename] [description]?\n", cmd.Name())
}

func (cmd FileCreate) validate(args []string) error {
	return nil
}
