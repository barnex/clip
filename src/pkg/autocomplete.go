package clip

// This file implements bash programmable completion.

import (
	"os"
	"fmt"
	"strings"
)

// Register the command
func init() {
	command["complete"] = Complete
}

func Complete(args []string) (resp string, err os.Error) {
	panic("please implement me")
	return
}


// Used by "clip -c", invoked by bash completion.
// args:
//	word line
// E.g.:
//	clip alpha beta<TAB>
// yields args:
//	beta clip alpha beta
func AutoComplete(args []string) {
	if len(args) == 0 {
		return // should not happen
	}
	if len(args) == 1 {
		// fix for word = "" (omitted by bash)
		args = []string{"", args[0], ""}
	}
	word := args[0]
	//cmd := args[1]
	//line := args[2:]
	if len(args) == 3 {
		completeCommands(word)
		return
	}
}


// Auto-complete function for player commands like 
//	add ls play ...
func completeCommands(prefix string) {
	for cmd, _ := range command {
		if strings.HasPrefix(cmd, prefix) {
			fmt.Print(cmd, " ")
		}
	}
}
