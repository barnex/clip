package clip

// This file implements the "play" command

import (
	"os"
)

// Register the command
func init() {
	player.command["play"] = Play
}

func Play(args []string) (resp string, err os.Error) {
	if len(args)>0{
		err = os.NewError("play does not take arguments yet")
		return
	}

	if player.playing{
		resp = "already playing"
	}
	
	if len(player.playlist) == 0{
		err = os.NewError("playlist empty")
		return
	}

	if player.current == -1{
		player.current = 0
	}
			go func() {
				player.backend.Play(player.playlist[player.current].file)
				player.playedChan <- 1
			}()
	return
}
