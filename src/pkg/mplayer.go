package clip

// This file uses mplayer as a playback back-end.

import (
	"exec"
	"log"
)

type MPlayer struct {
	cmd *exec.Cmd
}


func (m *MPlayer) Play(file string) chan int {
	m.cmd = exec.Command("mplayer", "-really-quiet", file)
	done := make(chan int)
	go func() {
		out, err := m.cmd.CombinedOutput()
		Check(err)
		log.Println("mplayer output", string(out))
		done <- 1
	}()
	return done
}

func (m *MPlayer) Stop() {

}
