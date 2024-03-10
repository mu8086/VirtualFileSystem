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
	register(FileCreate{})
}

type FileCreate struct{}

func (cmd FileCreate) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, folderName, fileName, desc := args[0], args[1], args[2], ""
	if len(args) == 4 {
		desc = args[3]
	}

	if err := dao.CreateFile(userName, folderName, fileName, desc); err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		case errors.ErrFolderNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", folderName)
		case errors.ErrFileExists:
			fmt.Fprintf(os.Stderr, "Error: The %v has already existed.\n", fileName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
	}

	fmt.Fprintf(os.Stdout, "Create %v in %v / %v successfully.\n", fileName, userName, folderName)
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
	if len(args) != 3 && len(args) != 4 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	userName, folderName, fileName := args[0], args[1], args[2]

	if !common.ValidUserName(userName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", userName)
		return errors.ErrUserName
	} else if !common.ValidFolderName(folderName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", folderName)
		return errors.ErrFolderName
	} else if !common.ValidFileName(fileName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", fileName)
		return errors.ErrFileName
	}
	return nil
}
