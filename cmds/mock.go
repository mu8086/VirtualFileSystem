package cmds

import (
	"VirtualFileSystem/errors"
	"fmt"
	"os"
)

func init() {
	//register(Mock{})
}

type Mock struct{}

func (cmd Mock) Execute(args []string) error {
	if err := cmd.validate(args); err != nil {
		return err
	}

	// do something
	arg0, arg1 := args[0], args[1]
	fmt.Fprintf(os.Stdout, "arg0: %v, arg1: %v\n", arg0, arg1)

	fmt.Fprintf(os.Stdout, "Execute %v successfully.\n", cmd)
	return nil
}

func (cmd Mock) String() string {
	return "MockCmd"
}

func (cmd Mock) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v [username] [foldername]\n", cmd)
}

func (cmd Mock) validate(args []string) error {
	if len(args) != 2 {
		cmd.Usage()
		return errors.ErrArgSize
	}

	arg0, arg1 := args[0], args[1]
	fmt.Fprintf(os.Stdout, "arg0: %v, arg1: %v\n", arg0, arg1)

	// validate fields:
	//
	// if err != nil {
	//     return err
	// }

	return nil
}
