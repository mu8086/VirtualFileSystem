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
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func handleCommand(input string) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return
	}

	command := fields[0]
	args := fields[1:]

	cmd := cmds.Get(command)
	if cmd == nil {
		fmt.Fprintf(os.Stdout, "invalid command: %v\n", command)
		return
	}

	fmt.Fprintf(os.Stdout, "[accept] command: %v(%v), args: %v\n", command, cmd.Name(), args)
	cmd.Execute()
}

func prompt() {
	fmt.Fprintf(os.Stdout, ">> ")
}
