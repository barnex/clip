package clip

// This file uses mplayer as a playback back-end.

import (
	"exec"
)

type MPlayer struct {
	*exec.Cmd
}

func (m *MPlayer) Playing() bool {
	return m.Cmd != nil
}

func (m *MPlayer) Play(file string) chan int {
	m.Stop()
	m.Cmd = exec.Command("mplayer", "-realy-quiet", file)
	done := make(chan int)
	go func() {
		err := m.Cmd.Run()
		Check(err)
		done <- 1
	}()
	return done
}

func (m *MPlayer) Stop() {
	if m.Cmd != nil {
		err := m.Cmd.Process.Kill()
		Check(err)
		m.Cmd = nil
	}
}
