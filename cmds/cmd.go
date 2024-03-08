package cmds

import (
	"VirtualFileSystem/errors"
	"fmt"
	"os"
)

type Cmd interface {
	Execute([]string) error
	Name() string
	String() string
	validate([]string) error
}

var cmds map[string]Cmd

func init() {
	cmds = make(map[string]Cmd)
}

func Get(cmdName string) Cmd {
	if cmd, exists := cmds[cmdName]; exists {
		return cmd
	}
	return nil
}

func AvailableCmds() string {
	s := ""
	for _, cmd := range cmds {
		s += " " + cmd.String()
	}
	if len(s) != 0 {
		s = s[1:]
	}
	return s
}

func register(cmd Cmd) bool {
	if cmd == nil {
		fmt.Fprintf(os.Stderr, "register: %v\n", errors.ErrCmdRegister)
		return false
	}

	if _, exists := cmds[cmd.Name()]; exists {
		fmt.Fprintf(os.Stderr, "register: %v\n", errors.ErrCmdExists)
		return false
	}

	cmds[cmd.Name()] = cmd
	return true
}
