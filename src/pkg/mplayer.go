package clip

// This file uses mplayer as a playback back-end.

import (
	"exec"
	"log"
)

type MPlayer struct {
	cmd *exec.Cmd
}


func (m *MPlayer) Play(file string) {
	m.cmd = exec.Command("mplayer", "-really-quiet", file)
	out, err := m.cmd.CombinedOutput()
	Check(err) // TODO: err==killed is OK
	log.Println("mplayer output", string(out))
}

func (m *MPlayer) Stop() {
	if m.cmd != nil {
		//err := 
		m.cmd.Process.Kill()
		//Check(err) // !! TODO: may already be finished in the meanwhile
	}
}
