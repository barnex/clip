package clip

// This file implements bash programmable completion.

import (
	"fmt"
	"strings"
)

// Used by "clip -c", invoked by bash completion.
// args:
//	word line
// E.g.:
//	clip alpha beta<TAB>
// yields args:
//	beta clip alpha beta
func AutoComplete(args []string) {
	if len(args) == 1 {
		// fix for word = "" (omitted by bash)
		args = []string{"", args[0], ""}
	}
	word := args[0]
	//cmd := args[1]
	//line := args[2:]
	if len(args) == 3 {
		completeCommands(word)
	}
}

var commands []string = []string{"add", "play", "pause", "stop"}

func completeCommands(prefix string) {
	for _, cmd := range commands {
		if strings.HasPrefix(cmd, prefix) {
			fmt.Print(cmd, " ")
		}
	}
}
