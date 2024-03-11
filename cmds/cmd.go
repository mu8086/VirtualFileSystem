package cmds

import (
	"VirtualFileSystem/constants"
	"VirtualFileSystem/errors"
	"fmt"
	"os"
	"sort"
)

type Cmd interface {
	Execute([]string) error
	String() string
	Usage()
	validate([]string) error
}

var cmds map[string]Cmd
var cmdsStr string

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
	if len(cmdsStr) == 0 {
		slice := []string{}
		for _, cmd := range cmds {
			slice = append(slice, cmd.String())
		}

		sort.Strings(slice)
		for _, cmd := range slice {
			cmdsStr += fmt.Sprintf(" [%v]", cmd)
		}

		if len(cmdsStr) > 0 {
			cmdsStr = cmdsStr[1:]
		}
	}
	return "Available Commands: " + cmdsStr
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
	fmt.Fprintf(os.Stdout, "%v\n", AvailableCmds())
	return nil
}

func (cmd CmdsList) String() string {
	return constants.CmdsList
}

func (cmd CmdsList) Usage() {

}

func (cmd CmdsList) validate(args []string) error {
	return nil
}
