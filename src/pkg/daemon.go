package clip

import (
	"rpc"
	"net"
	"http"
)

const(
	PORT = ":25274"
)

func MainDaemon(args []string){
	RPC := new(int)
	rpc.Register(RPC)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", PORT)
	if e != nil {
		Err("listen error:", e)
	}
	Debug("Listening on port " + PORT)
	http.Serve(l, nil)
}
