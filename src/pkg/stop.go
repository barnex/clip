package clip

// This file implements the "pause" command

import (
	"os"
)

// Register the command
func init() {
	player.command["stop"] = Stop
}

func Stop(args []string) (resp string, err os.Error) {
	if len(args) > 0 {
		err = os.NewError("stop does not take arguments")
		return
	}

	if !player.playing {
		resp = "Already stopped."
		return
	}

	player.playing = false
	player.backend.Stop()

	return
}
