package clip

// This file uses mplayer as a playback back-end.

import (
	"exec"
	"log"
)

type MPlayer struct {
}


func (m *MPlayer) Play(file string) chan int {
	cmd := exec.Command("mplayer", "-really-quiet", file)
	done := make(chan int)
	go func() {
		out, err := cmd.CombinedOutput()
		Check(err)
		log.Println("mplayer output", string(out))
		done <- 1
	}()
	return done
}

func (m *MPlayer) Stop() {
	
}
