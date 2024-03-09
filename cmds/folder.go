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
	register(FolderCreate{})
}

type FolderCreate struct{}

func (cmd FolderCreate) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, folderName, desc := args[0], args[1], ""
	if len(args) == 3 {
		desc = args[2]
	}

	if !dao.CreateFolder(userName, folderName, desc) {
		fmt.Fprintf(os.Stderr, "Error: %v\n", errors.ErrUnknown)
		return errors.ErrUnknown
	}

	fmt.Fprintf(os.Stdout, "Create %s successfully.\n", folderName)
	return nil
}

func (cmd FolderCreate) Name() string {
	return constants.FolderCreateCmd
}

func (cmd FolderCreate) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FolderCreate) Usage() {
	fmt.Fprintf(os.Stdout, "Usage: %v [username] [foldername] [description]?\n", cmd.Name())
}

func (cmd FolderCreate) validate(args []string) error {
	if len(args) != 2 && len(args) != 3 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	userName, folderName := args[0], args[1]

	if !common.ValidUserName(userName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", userName)
		return errors.ErrUserName
	} else if !common.ValidFolderName(folderName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", folderName)
		return errors.ErrFolderName
	}

	if dao.GetUser(userName) == nil {
		fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		return errors.ErrUserNotExists
	} else if dao.GetUserFolder(userName, folderName) != nil {
		fmt.Fprintf(os.Stderr, "Error: The %v has already existed.\n", folderName)
		return errors.ErrFolderExists
	}

	return nil
}
