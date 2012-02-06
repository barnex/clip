package clip

// This file implements the "add" command

import (
	"os"
	"fmt"
)

// Register the command
func init() {
	player.command["add"] = Add
}

func Add(args []string) (resp string, err os.Error) {
	if len(args) == 0 {
		err = os.NewError("nothing specified, nothing added")
		return
	}

	found := 0
	notfound := 0
	errstr := "not found:"
	for _, arg := range args {
		items := player.library.Find(arg)

		// TODO: item may be directory, add recursively
		if len(items) == 0 {
			errstr += " " + arg
			notfound++
		}

		player.playlist.Append(items...)
		found += len(items)

	}

	if notfound > 0 {
		err = os.NewError(errstr)
	}
	if found > 0 {
		resp = fmt.Sprint("added ", found, " files to playlist")
	}
	return
}
