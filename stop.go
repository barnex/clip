package main

// This file implements the "stop" command

func (player *Player) Stop() (resp, err string) {

	if !player.playing {
		resp = "Already stopped."
		return
	}

	player.playing = false
	player.backend.Stop()

	return
}
