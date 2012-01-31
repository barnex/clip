package clip

// This file implements the "add" command

import (
	"os"
)

// Register the command
func init() {
	command["add"] = Add
}

func Add(args []string) (resp string, err os.Error) {
	if len(args) == 0 {
		err = os.NewError("nothing specified, nothing added")
		return
	}
	for _, arg := range args {
		AddPath(library.fs, arg)
		resp += "Added " + arg + "\n"
	}
	return
}
