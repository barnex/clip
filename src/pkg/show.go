package clip

// This file implements the "show" command

import (
	"os"
)

// Register the command
func init() {
	player.command["show"] = Show
}

func Show(args []string) (resp string, err os.Error) {
	// keep map of Showers/Stringers?
	// also resolve in tags/only if not found?	
	// show playlist, artists, library, ...
}
