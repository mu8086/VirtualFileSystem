package errors

import stdErr "errors"

var (
	ErrUnknown = stdErr.New("unknown")

	ErrArgSize = stdErr.New("invalid arg size")

	ErrCmdExists   = stdErr.New("command exists")
	ErrCmdNotFound = stdErr.New("command not found")
	ErrCmdParse    = stdErr.New("command parse failed")
	ErrCmdRegister = stdErr.New("command register failed")

	ErrUserExists = stdErr.New("user exists")
	ErrUserName   = stdErr.New("username invalid chars")
)
