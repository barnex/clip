package clip

// This file implements the player state.
// The player communicates over RPC (Remote Procedure Call)
// with the client. The client forks a player daemon (clip -d)
// if none is yet running and sends RPC calls to it.

import (
	"os"
	"rpc"
	"net"
	"http"
)

var player Player = Init()


func Init() Player{
	var player Player
	player.Init()
	return player
}

func (p *Player) Init() {
	Debug("player initialized")
	p.library = NewLib()
	p.playlist = ItemArray([]*Item{})
	p.command = make(map[string]Command)
	p.port = ":25274"
	p.backend = new(MPlayer)
	p.callChan = make(chan Call)
	p.playedChan = make(chan int)
}

// Here the player state is stored.
type Player struct {
	library    *Lib               // the player's library
	playlist  ItemArray
	command    map[string]Command // the player's commands
	port       string             // default RPC port
	backend    Backend
	callChan   chan Call // calls ("play", ...) are sent here
	playedChan chan int
}


// A command (e.g. "add") takes string arguments provided by the user
// (e.g. "/home/me/music" )and returns a message to the user (e.g. "added 7 files").
type Command func([]string) (string, os.Error)


func MainDaemon(args []string){
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


// Start serving RPC calls from client instances.
func (p *Player) serveRPC() {
	rpc.Register((*PlayerRPC)(&player))
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", p.port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + p.port)
	http.Serve(conn, nil)
	//TODO: log errors.
}


// Dummy type to define RPC methods on.
type PlayerRPC Player

// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
func (rpc *PlayerRPC) Call(args []string, resp *string) (err os.Error) {
	p := (*Player)(rpc)
	call := NewCall(args)       // wrap args in Call struct
	p.callChan <- call          // send to event loop for execution
	callResp := <-call.respChan // wait for response
	*resp = callResp.Resp       // set return value
	return callResp.Err
}
