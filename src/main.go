package main

import (
	"clip"
	"os"
)

func main() {
	lib := clip.NewLib()
	lib.Add("/home/arne/music/kraftwerk/autobahn.mp3")
	lib.WriteTo(os.Stdout)
}
