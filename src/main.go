package main

import (
	"clip"
	"flag"
	"os"
)

var (
	flag_complete *bool = flag.Bool("c", false, "bash completion of arguments")
)

func main() {
	flag.Parse()

	if *flag_complete {
		clip.AutoComplete(flag.Args())
		return
	}

	lib := clip.NewLib()
	lib.AddPath("/home/arne/music/kraftwerk/autobahn.mp3")
	lib.AddPath("/home/arne/personal/kraftwerk/autobahn.mp3")
	lib.WriteTo(os.Stdout)
}
