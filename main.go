package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lineScanner := bufio.NewScanner(os.Stdin)

	prompt()
	for lineScanner.Scan() {
		parseCommand(lineScanner.Text())
		prompt()
	}
	if err := lineScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parseCommand(input string) {
}

func prompt() {
	fmt.Fprintf(os.Stdout, ">> ")
}
