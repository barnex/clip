package clip

import (
	"os"
	"rpc"
	"net"
	"http"
)

type Daemon struct {
	lib *Lib
}

func NewDaemon() *Daemon {
	return &Daemon{NewLib()}
}


func (d *Daemon) AutoComplete(args []string, resp *string) (err os.Error) {
	*resp = "arne"
	return nil
}

func (d *Daemon) Call(args []string, resp *string) (err os.Error) {
	*resp = "called me?"
	return nil
}


const (
	PORT = ":25274"
)

func MainDaemon(args []string) {
	daemon := NewDaemon()
	rpc.Register(daemon)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", PORT)
	if e != nil {
		Err("listen error:", e)
	}
	Debug("Listening on port " + PORT)
	http.Serve(l, nil)
}
