package main

import (
	"clip"
	"os"
)

func main() {
	lib := clip.NewLib()
	lib.AddPath("/home/arne/music/kraftwerk/autobahn.mp3")
	lib.AddPath("/home/arne/personal/kraftwerk/autobahn.mp3")
	lib.WriteTo(os.Stdout)
}
