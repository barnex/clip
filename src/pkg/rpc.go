package clip

// This file implements the Remote Procedure Call between
// the clip daemon and front-ends

import (
	"os"
	"rpc"
	"net"
	"http"
)

// Start serving RPC calls from client instances.
func (player *Player) serveRPC() {
	rpc.Register((*PlayerRPC)(player))
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + port)
	http.Serve(conn, nil)
	//TODO: log errors.
}

// Dummy type to define RPC methods on.
type PlayerRPC Player


// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
func (rpc *PlayerRPC) Call(args []string, resp *string) (err os.Error) {
	Debug("PlayerRPC.Call", args)
	//	p := (*Player)(rpc)
	//	call := NewCall(args)       // wrap args in Call struct
	//	p.callChan <- call          // send to event loop for execution
	//	callResp := <-call.respChan // wait for response
	//	*resp = callResp.Resp       // set return value
	//	return callResp.Err
	return
}
