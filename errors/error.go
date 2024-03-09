package errors

import stdErr "errors"

var (
	ErrUnknown = stdErr.New("unknown")

	ErrArgSize = stdErr.New("arg invalid size")

	ErrCmdExists    = stdErr.New("command exists")
	ErrCmdNotExists = stdErr.New("Unrecognized command")
	ErrCmdParse     = stdErr.New("command parse failed")
	ErrCmdRegister  = stdErr.New("command register failed")

	ErrFolderExists    = stdErr.New("folder exists")
	ErrFolderName      = stdErr.New("folder name contains invalid chars")
	ErrFolderNotExists = stdErr.New("folder not exists")

	ErrSortFlag   = stdErr.New("sort flag not exists")
	ErrSortOption = stdErr.New("sort option not exists")

	ErrUserExists    = stdErr.New("user exists")
	ErrUserName      = stdErr.New("user name contains invalid chars")
	ErrUserNotExists = stdErr.New("user not exists")
)
