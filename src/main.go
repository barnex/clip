package main

// This file implements the CLIP main function

import (
	"clip"
	"flag"
)

// Command-line flags for special modes
// not normally used by the user.
var (
	flag_complete *bool = flag.Bool("c", false, "bash completion of arguments")
	flag_daemon   *bool = flag.Bool("d", false, "run in daemon mode")
)

func main() {
	flag.Parse()

	//	if *flag_complete {
	//		clip.AutoComplete(flag.Args())
	//		return
	//	}

	if *flag_daemon {
		clip.NewPlayer().Daemon()
	}

	clip.MainClient(flag.Args())
}
