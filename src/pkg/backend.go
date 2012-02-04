package clip

// This file provides the player back-end interface.

import ()

type Backend interface {
	Play(file string) // play a file, block until done
	Stop()            // stop playing
}
