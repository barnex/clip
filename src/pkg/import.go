package clip

// This file implements the "import" command

import (
	"os"
)

// Register the command
func init() {
	player.command["import"] = Import
}

func Import(args []string) (resp string, err os.Error) {
	if len(args) == 0 {
		err = os.NewError("nothing specified, nothing imported")
		return
	}
	for _, arg := range args {
		player.library.Import(arg)
	}
	return
}
