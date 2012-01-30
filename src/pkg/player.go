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

// A command (e.g. "add") takes string arguments provided by the user
// (e.g. "/home/me/music" )and returns a message to the user (e.g. "added 7 files").
type Command func([]string) (string, os.Error)

// Here the player state is stored.
var (
	library *Lib               = NewLib() // the player's library
	command map[string]Command = make(map[string]Command) // the player's commands
	port    string             = ":25274" // default RPC port
)

// Dummy type to define RPC methods on.
type PlayerRPC struct{} 

// RPC-exported function used for auto-completion (clip -c).
// The command-line arguments are passed (see complete.bash)
// and a list of completions is returned in *resp.
func (d *PlayerRPC) AutoComplete(args []string, resp *string) (err os.Error) {
	*resp = "arne"
	return nil
}

// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
func (d *PlayerRPC) Call(args []string, resp *string) (err os.Error) {
	if len(args) == 0 {
		args = []string{""}
	}
	cmd := args[0]
	args = args[1:]
	Debug("player.call", cmd, args)
	f, ok := command[cmd]
	if !ok {
		err = os.NewError("no such command: " + cmd)
	} else {
		*resp, err = f(args)
	}
	return
}


// Main loop for daemon mode
func MainDaemon(args []string) {
	serveRPC()
}

// Start serving RPC calls from client instances.
func serveRPC() {
	rpc.Register(&PlayerRPC{})
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + port)
	http.Serve(conn, nil)
	//TODO: log errors.
}
