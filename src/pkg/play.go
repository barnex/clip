package clip

// This file implements the "play" command

import (
	"os"
)

// Register the command
func init() {
	command["play"] = Play
}

func Play(args []string) (resp string, err os.Error) {
	if len(args) == 0 {
		err = os.NewError("nothing specified, nothing played")
		return
	}
	for _, arg := range args {
		items := library.Find(arg)
		if len(items) == 0 {
			err = os.NewError(arg + " not found")
			return
		}
		for _, i := range items {
			go func() {
				backend.Play(i.file)
				playedChan <- 1
			}()
		}
	}
	return
}
