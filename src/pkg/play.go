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
		done := backend.Play(arg)
		<-done
	}
	return
}
