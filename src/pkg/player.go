package clip

// This file implements the player state.
// The player communicates over RPC (Remote Procedure Call)
// with the client. The client forks a player daemon (clip -d)
// if none is yet running and sends RPC calls to it.

import (
	"sync"
)

type Player struct {
	*Lib // the player's library
	playlist ItemArray
	current  int // current track
	playing  bool
	backend  Backend
	port     string // default RPC port
	sync.Mutex
}

func(p*Player)API()API{
	return API{p}
}

func(p*Player)RPC()RPC{
	return RPC{p}
}

func NewPlayer() *Player {
	p := new(Player)
	p.init()
	return p
}

func (p *Player) init() {
	Debug("player initialized")
	p.Lib = NewLib()
	p.playlist = ItemArray([]*Item{})
	p.playing = false
	p.current = -1
	p.port = ":25274"
	p.backend = new(MPlayer)
}

// Main loop for daemon mode
func (p *Player) Daemon() {
	// TODO: heartbeat here
	p.serveRPC()
}
