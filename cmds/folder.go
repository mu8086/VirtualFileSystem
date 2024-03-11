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
	register(FolderCreate{})
	register(FoldersList{})
	register(FolderRemove{})
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

func (cmd FolderCreate) String() string {
	return constants.FolderCreateCmd
}

func (cmd FolderCreate) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername] [description]?\n", cmd)
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

type FoldersList struct{}

func (cmd FoldersList) Execute(args []string) error {
	// If neither --sort-name nor --sort-created is provided, sort the list by [foldername] in ascending order
	if len(args) == 1 {
		args = append(args, constants.OptionSortByName, constants.FlagSortAsc)
	}

	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, sortOption, sortFlag := args[0], args[1], args[2]

	sortedFolders, err := dao.GetFolders(userName, sortOption, sortFlag)
	if err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
	} else if sortedFolders == nil || len(sortedFolders) == 0 {
		fmt.Fprintf(os.Stdout, "Warning: The %v doesn't have any folders.\n", userName)
		return nil
	}

	fmt.Fprintf(os.Stdout, "%v", strings.ReplaceAll(
		sortedFolders.String(),
		"\n",
		fmt.Sprintf(" %v\n", userName)))

	return nil
}

func (cmd FoldersList) String() string {
	return constants.FoldersListCmd
}

func (cmd FoldersList) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [%v|%v] [%v|%v]\n", cmd,
		constants.OptionSortByName, constants.OptionSortByCreated,
		constants.FlagSortAsc, constants.FlagSortDesc)
}

func (cmd FoldersList) validate(args []string) error {
	if len(args) != 3 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	userName, sortOption, sortFlag := args[0], args[1], args[2]

	if !common.ValidUserName(userName) {
		fmt.Fprintf(os.Stderr, "Error: The %v contain invalid chars.\n", userName)
		return errors.ErrUserName
	} else if !common.ValidSortOption(sortOption) {
		cmd.Usage()
		return errors.ErrSortOption
	} else if !common.ValidSortFlag(sortFlag) {
		cmd.Usage()
		return errors.ErrSortFlag
	}
	return nil
}

type FolderRemove struct{}

func (cmd FolderRemove) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	userName, folderName := args[0], args[1]

	if err := dao.RemoveFolder(userName, folderName); err != nil {
		switch err {
		case errors.ErrUserNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", userName)
		case errors.ErrFolderNotExists:
			fmt.Fprintf(os.Stderr, "Error: The %v doesn't exist.\n", folderName)
		default:
			fmt.Fprintf(os.Stderr, "Unknown Error: %v\n", err)
		}
		return err
	}

	fmt.Fprintf(os.Stdout, "Delete %v successfully.\n", folderName)
	return nil
}

func (cmd FolderRemove) String() string {
	return constants.FolderRemoveCmd
}

func (cmd FolderRemove) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername]\n", cmd)
}

func (cmd FolderRemove) validate(args []string) error {
	if len(args) != 2 {
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

func (cmd FolderRename) String() string {
	return constants.FolderRenameCmd
}

func (cmd FolderRename) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername] [new-folder-name]\n", cmd)
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
