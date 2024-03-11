package cmds

import (
	"VirtualFileSystem/common"
	"VirtualFileSystem/constants"
	"VirtualFileSystem/dao"
	"VirtualFileSystem/errors"
	"fmt"
	"os"
	"strings"
)

func init() {
	register(FileCreate{})
	register(FilesList{})
	register(FileRemove{})
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

type FileRemove struct{}

func (cmd FileRemove) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, folderName, fileName := args[0], args[1], args[2]
	if err := dao.RemoveFile(userName, folderName, fileName); err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		case errors.ErrFolderNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", folderName)
		case errors.ErrFileNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", fileName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
	}

	fmt.Fprintf(os.Stdout, "Delete %v in %v / %v successfully.\n", fileName, userName, folderName)
	return nil
}

func (cmd FileRemove) Name() string {
	return constants.FileRemoveCmd
}

func (cmd FileRemove) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FileRemove) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername] [filename]\n", cmd.Name())
}

func (cmd FileRemove) validate(args []string) error {
	if len(args) != 3 {
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

type FilesList struct{}

func (cmd FilesList) Execute(args []string) error {
	// If neither --sort-name nor --sort-created is provided, sort the list by [foldername] in ascending order
	if len(args) == 2 {
		args = append(args, constants.OptionSortByName, constants.FlagSortAsc)
	}

	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, folderName, sortOption, sortFlag := args[0], args[1], args[2], args[3]

	sortedFiles, err := dao.GetFiles(userName, folderName, sortOption, sortFlag)
	if err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		case errors.ErrFolderNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", folderName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
	} else if sortedFiles == nil || len(sortedFiles) == 0 {
		fmt.Fprintf(os.Stdout, "Warning: The %v doesn't have any files.\n", folderName)
		return nil
	}

	fmt.Fprintf(os.Stdout, "%v", strings.ReplaceAll(
		sortedFiles.String(),
		"\n",
		fmt.Sprintf(" %v %v\n", folderName, userName)))
	return nil
}

func (cmd FilesList) Name() string {
	return constants.FilesListCmd
}

func (cmd FilesList) String() string {
	return fmt.Sprintf("[%s]", cmd.Name())
}

func (cmd FilesList) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername] [%v|%v] [%v|%v]\n", cmd.Name(),
		constants.OptionSortByName, constants.OptionSortByCreated,
		constants.FlagSortAsc, constants.FlagSortDesc)
}

func (cmd FilesList) validate(args []string) error {
	if len(args) != 4 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	userName, folderName, sortOption, sortFlag := args[0], args[1], args[2], args[3]

	if !common.ValidUserName(userName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", userName)
		return errors.ErrUserName
	} else if !common.ValidFolderName(folderName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", folderName)
		return errors.ErrFolderName
	} else if !common.ValidSortOption(sortOption) {
		cmd.Usage()
		return errors.ErrSortOption
	} else if !common.ValidSortFlag(sortFlag) {
		cmd.Usage()
		return errors.ErrSortFlag
	}
	return nil
}
