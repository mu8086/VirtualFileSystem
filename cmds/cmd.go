package cmds

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"fmt"
	"os"
)

type Cmd interface {
	Execute([]string) error
	Name() string
	String() string
	Usage()
	validate([]string) error
}

var cmds map[string]Cmd

func init() {
	cmds = make(map[string]Cmd)

	register(CmdsList{})
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
		s += fmt.Sprintf(" [%v]", cmd)
	}
	if len(s) != 0 {
		s = s[1:]
	}
	return "Available Commands: " + s
}

func register(cmd Cmd) bool {
	if cmd == nil {
		fmt.Fprintf(os.Stderr, "register: %v\n", errors.ErrCmdRegister)
		return false
	}

	if _, exists := cmds[cmd.String()]; exists {
		fmt.Fprintf(os.Stderr, "register: %v\n", errors.ErrCmdExists)
		return false
	}

	cmds[cmd.String()] = cmd
	return true
}

type CmdsList struct{}

func (cmd CmdsList) Execute(args []string) error {
	return nil
}

func (cmd CmdsList) Name() string {
	return cmd.String()
}

func (cmd CmdsList) String() string {
	return constants.CmdsList
}

func (cmd CmdsList) Usage() {

}

func (cmd CmdsList) validate(args []string) error {
	return nil
}
