package main

import (
	"VirtualFileSystem/cmds"
	"VirtualFileSystem/dao"
	"VirtualFileSystem/errors"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("available cmds: %v\n", cmds.AvailableCmds())

	lineScanner := bufio.NewScanner(os.Stdin)

	prompt()
	for lineScanner.Scan() {
		handleCommand(lineScanner.Text())
		prompt()
	}
	if err := lineScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input: %v\n", err)
	}
}

// TODO: implement error
func parseCommand(input string) (string, []string, error) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "", nil, nil
	}
	return fields[0], fields[1:], nil
}

func handleCommand(input string) {
	cmdName, args, err := parseCommand(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", errors.ErrCmdParse, err)
		return
	} else if cmdName == "" { // empty line
		return
	}

	// TODO: remove
	if cmdName == "all" {
		fmt.Printf("all user: %v\n", dao.GetAllUsers())
		return
	}

	cmd := cmds.Get(cmdName)
	if cmd == nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", errors.ErrCmdNotExists, cmdName)
		return
	}

	err = cmd.Execute(args)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "command: '%v' encountered err: %v\n", cmdName, err)
		return
	}
}

func prompt() {
	fmt.Fprintf(os.Stdout, ">> ")
}
