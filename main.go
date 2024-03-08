package main

import (
	"VirtualFileSystem/cmds"
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
		fmt.Fprintf(os.Stderr, "parse command failed, err: %+v", err)
		return
	}

	cmd := cmds.Get(cmdName)
	if cmd == nil {
		fmt.Fprintf(os.Stdout, "invalid command: %v\n", cmdName)
		return
	}

	fmt.Fprintf(os.Stdout, "[accept] command: %v(%v), args: %v\n", cmdName, cmd.Name(), args)
	cmd.Execute()
}

func prompt() {
	fmt.Fprintf(os.Stdout, ">> ")
}
