package main

import (
	"clip"
	"fmt"
	"os"
)

func main() {
	lib := clip.NewLib()
	lib.Add("/home/arne/music/kraftwerk/autobahn.mp3")
	fmt.Println(lib)
	n, err := lib.WriteTo(os.Stdout)
	clip.Debug("n", n, "err", err)
}
