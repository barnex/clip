package clip

// This file implements the player state.
// The player communicates over RPC (Remote Procedure Call)
// with the client. The client forks a player daemon (clip -d)
// if none is yet running and sends RPC calls to it.

import (
	"os"
)

var player Player = Init()


func Init() Player {
	var player Player
	player.Init()
	return player
}

func (p *Player) Init() {
	// TODO: no good for client mode!
	Debug("player initialized")
	p.library = NewLib()
	p.playlist = ItemArray([]*Item{})
	p.playing = false
	p.current = -1
	p.command = make(map[string]Command)
	p.port = ":25274"
	p.backend = new(MPlayer)
	p.callChan = make(chan Call)
	p.playedChan = make(chan int)
}

// Here the player state is stored.
type Player struct {
	library    *Lib // the player's library
	playlist   ItemArray
	current    int // current track
	playing    bool
	backend    Backend
	callChan   chan Call // calls ("play", ...) are sent here
	playedChan chan int
	command    map[string]Command // the player's commands
	port       string             // default RPC port
}


// A command (e.g. "add") takes string arguments provided by the user
// (e.g. "/home/me/music" )and returns a message to the user (e.g. "added 7 files").
type Command func([]string) (string, os.Error)


func MainDaemon(args []string) {
	player.MainDaemon(args)
}

// Main loop for daemon mode
func (p *Player) MainDaemon(args []string) {
	go p.serveRPC()
	// event loop
	for {
		select {
		case call := <-p.callChan:
			resp := call.Exec()
			call.respChan <- resp
		}
	}
}
