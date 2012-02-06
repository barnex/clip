package clip

// This file implements the "next" command

import (
	"os"
)

// Register the command
func init() {
	player.command["next"] = Next
}

func Next(args []string) (resp string, err os.Error) {
	if len(args) > 0 {
		err = os.NewError("next does not take arguments yet")
		// TODO: next album, next artist, ...
		return
	}

	if len(player.playlist) == 0 {
		err = os.NewError("playlist empty")
		return
	}

	if player.current+1 < len(player.playlist) {
		player.backend.Stop() // TODO: not thread safe!!
		// will advance to next automatically // TODO: player.next...
	}

	return
}
