package main

// This file implements the "add" command

import (
	"fmt"
)

func (player *Player) Add(args []string) (resp, err string) {
	if len(args) == 0 {
		err = ("nothing specified, nothing added")
		return
	}

	found := 0
	notfound := 0
	errstr := "not found:"
	for _, arg := range args {
		items := player.Find(arg)

		// TODO: item may be directory, add recursively
		if len(items) == 0 {
			errstr += " " + arg
			notfound++
		}

		player.playlist.Append(items...)
		found += len(items)

	}

	if notfound > 0 {
		err = errstr
	}
	if found > 0 {
		resp = fmt.Sprint("added ", found, " files to playlist")
	}
	return
}
