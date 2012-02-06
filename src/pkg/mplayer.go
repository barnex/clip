package clip

// This file uses mplayer as a playback back-end.

import (
	"exec"
	"log"
	"os"
	//"fmt"
)

type MPlayer struct {
	cmd *exec.Cmd
}


func (m *MPlayer) Play(file string) {
	m.cmd = exec.Command("mplayer", "-really-quiet", file)
	out, err := m.cmd.CombinedOutput()

	// If the command was killed with signal 9,
	// the player was just stopped, so we don't crash
	if waitmsg, ok := err.(*os.Waitmsg)	; ok{
		if waitmsg.WaitStatus != 9{
			Check(err)
		}
	}else{
		Check(err) // TODO: don't crash 
	}

	log.Println("mplayer output", string(out))
}

func (m *MPlayer) Stop() {
	if m.cmd != nil {
		//err := 
		m.cmd.Process.Kill()
		//Check(err) // !! TODO: may already be finished in the meanwhile
	}
}
