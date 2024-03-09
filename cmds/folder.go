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
	register(FolderRename{})
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

	if err := dao.CreateFolder(userName, folderName, desc); err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		case errors.ErrFolderExists:
			fmt.Fprintf(os.Stderr, "Error: The %v has already existed.\n", folderName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
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
	return nil
}

type FolderRename struct{}

func (cmd FolderRename) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, folderName, newFolderName := args[0], args[1], args[2]

	if err := dao.UpdateFolder(userName, folderName, newFolderName); err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		case errors.ErrFolderNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", folderName)
		case errors.ErrFolderExists:
			fmt.Fprintf(os.Stderr, "Error: The %v has already existed.\n", newFolderName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
	}

	fmt.Fprintf(os.Stdout, "Rename %v to %v successfully.\n", folderName, newFolderName)
	return nil
}

func (cmd FolderRename) Name() string {
	return constants.FolderRenameCmd
}

func (cmd FolderRename) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FolderRename) Usage() {
	fmt.Fprintf(os.Stdout, "Usage: %v [username] [foldername] [new-folder-name]\n", cmd.Name())
}

func (cmd FolderRename) validate(args []string) error {
	if len(args) != 3 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	userName, folderName, newFolderName := args[0], args[1], args[2]

	if !common.ValidUserName(userName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", userName)
		return errors.ErrUserName
	} else if !common.ValidFolderName(folderName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", folderName)
		return errors.ErrFolderName
	} else if !common.ValidFolderName(newFolderName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", newFolderName)
		return errors.ErrFolderName
	}
	return nil
}
