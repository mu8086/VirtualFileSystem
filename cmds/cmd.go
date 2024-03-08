package cmds

import (
	"fmt"
	"os"
)

type Cmd interface {
	Execute()
	Name() string
	String() string
	validate() bool
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
	return s[1:]
}

func register(cmd Cmd) bool {
	if cmd == nil {
		fmt.Fprintf(os.Stderr, "register cmd failed: error argument.\n")
		return false
	}

	if _, exists := cmds[cmd.Name()]; !exists {
		cmds[cmd.Name()] = cmd
		return true
	}

	fmt.Fprintf(os.Stderr, "register cmd failed: already exists.\n")
	return false
}
