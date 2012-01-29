package clip

import (
	"os"
	"rpc"
	"net"
	"http"
)

type Command func([]string) (string, os.Error)

var (
	library *Lib               = NewLib()
	command map[string]Command = make(map[string]Command)
	port    string             = ":25274"
)

type PlayerRPC struct{} // dummy type

func (d *PlayerRPC) AutoComplete(args []string, resp *string) (err os.Error) {
	*resp = "arne"
	return nil
}

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


func MainDaemon(args []string) {
	serveRPC()
}

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
