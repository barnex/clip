package main

// This file implements the CLIP main function

import (
	"flag"
	"fmt"
	"runtime"
)

// Command-line flags for special modes
// not normally used by the user.
var (
	//flag_complete *bool = flag.Bool("c", false, "bash completion of arguments")
	flag_daemon  *bool = flag.Bool("d", false, "run in daemon mode")
	flag_version *bool = flag.Bool("v", false, "show version and exit")
)

func main() {
	flag.Parse()

	if *flag_version {
		fmt.Println(`CLIP 0.0.80 "Hupsje"`)
		fmt.Println("Go ", runtime.Version())
		return
	}

	//	if *flag_complete {
	//		clip.AutoComplete(flag.Args())
	//		return
	//	}

	if *flag_daemon {
		NewPlayer().Daemon()
	}

	MainClient(flag.Args())
}
