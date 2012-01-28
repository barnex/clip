package main

import (
	"clip"
	"flag"
	"os"
)

var (
	flag_complete *bool = flag.Bool("c", false, "bash completion of arguments")
	flag_daemon *bool = flag.Bool("d", false, "run in daemon mode")
)

func main() {
	flag.Parse()

	if *flag_complete {
		clip.AutoComplete(flag.Args())
		return
	}

	if *flag_daemon{
		clip.MainDaemon(flag.Args())
		return
	}

	lib := clip.NewLib()
	lib.AddPath("/home/arne/music/kraftwerk/autobahn.mp3")
	lib.AddPath("/home/arne/personal/kraftwerk/autobahn.mp3")
	lib.WriteTo(os.Stdout)
}
