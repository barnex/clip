package clip

import (
	"os"
	"rpc"
	"fmt"
	"net"
	"http"
)

var (
	library *Lib
)

type PlayerRPC struct{} // dummy type

func (d *PlayerRPC) AutoComplete(args []string, resp *string) (err os.Error) {
	*resp = "arne"
	return nil
}

func (d *PlayerRPC) Call(args []string, resp *string) (err os.Error) {
	*resp = fmt.Sprint("called:", args)
	return nil
}


const (
	PORT = ":25274"
)

func MainDaemon(args []string) {
	rpc.Register(&PlayerRPC{})
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", PORT)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + PORT)
	http.Serve(conn, nil)
	//TODO: log errors.
}
