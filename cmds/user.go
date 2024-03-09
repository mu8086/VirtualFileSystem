package cmds

import (
	"VirtualFileSystem/common"
	"VirtualFileSystem/constants"
	"VirtualFileSystem/dao"
	"VirtualFileSystem/errors"
	"fmt"
	"os"
)

func init() {
	register(UserCreate{})
}

type UserCreate struct{}

func (cmd UserCreate) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	name := args[0]
	if !dao.CreateUser(name) {
		fmt.Fprintf(os.Stderr, "Error: %v\n", errors.ErrUnknown)
		return errors.ErrUnknown
	}

	fmt.Fprintf(os.Stdout, "Add %s successfully.\n", name)
	return nil
}

func (cmd UserCreate) Name() string {
	return constants.UserCreateCmd
}

func (cmd UserCreate) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd UserCreate) Usage() {
	fmt.Fprintf(os.Stdout, "Usage: %v [username]\n", cmd.Name())
}

func (cmd UserCreate) validate(args []string) error {
	if len(args) != 1 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	name := args[0]
	if !common.ValidUserName(name) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", name)
		return errors.ErrUserName
	} else if dao.GetUser(name) != nil {
		fmt.Fprintf(os.Stderr, "Error: The %v has already existed.\n", name)
		return errors.ErrUserExists
	}

	return nil
}
