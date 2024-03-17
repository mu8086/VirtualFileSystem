package main

import (
	"VirtualFileSystem/cmds"
	"VirtualFileSystem/errors"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	loadViper("local.toml")

	fmt.Printf("%v\n", cmds.AvailableCmds())

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

func loadViper(configFile string) error {
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not load viper configuration: %v", err)
		return err
	}
	return nil
}

func parseCommand(input string) (string, []string) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "", nil
	}
	return fields[0], fields[1:]
}

func handleCommand(input string) error {
	cmdName, args := parseCommand(input)
	if cmdName == "" { // empty line
		return nil
	}

	cmd := cmds.Get(cmdName)
	if cmd == nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", errors.ErrCmdNotExists)
		return errors.ErrCmdNotExists
	}

	return cmd.Execute(args)
}

func prompt() {
	fmt.Fprintf(os.Stdout, ">> ")
}
