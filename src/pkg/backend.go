package clip

// This file provides the player back-end interface.

import ()

type Backend interface {
	Play(file string) chan int
	Stop()
}
