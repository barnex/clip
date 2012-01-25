package main

import("clip"
"fmt")

func main(){
	lib := clip.NewLib()
	lib.Add("/home/arne/music/kraftwerk/autobahn.mp3")
	fmt.Print(lib)
}
