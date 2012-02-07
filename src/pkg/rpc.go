package clip

// This file implements the Remote Procedure Call between
// the clip daemon and client front-end

import (
	"os"
	"fmt"
	"rpc"
	"net"
	"http"
	"reflect"
	"unicode"
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

// Aliased type to define RPC methods on.
type PlayerRPC Player

// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
func (rpc *PlayerRPC) Call(args []string, resp *string) (err os.Error) {
	Debug("PlayerRPC.Call", args)

	player := (*Player)(rpc)

	cmd := args[0]  // first arg is command (e.g.: "play")
	args = args[1:] // rest are arguments (e.g.: "jazz")

	// convert first character to uppercase
	first := unicode.ToUpper(int(cmd[0]))
	Cmd := string(first) + cmd[1:] // (e.g.: Play)

	// resolve the command using reflection
	p := reflect.ValueOf(player)
	m := p.MethodByName(Cmd)
	Debug("MethodByName", Cmd, ":", m)
	if m.Kind() == reflect.Invalid {
		err = os.NewError("clip: '" + cmd + "' is not a clip command. See clip help.")
		return
	}

	// set up method arguments
	ins := m.Type().NumIn()
	var callArgs []reflect.Value
	switch ins {
	default:
		err = os.NewError(fmt.Sprint("Bug: wrong number of ins: ", ins))
		return
	case 0:
		if len(args) > 0 {
			err = os.NewError(fmt.Sprint(cmd, " does not take arugments"))
			return
		}
		callArgs = []reflect.Value{}
	case 1:
		if len(args) == 0 {
			err = os.NewError(fmt.Sprint(cmd, " needs an argument"))
			return
		}
		callArgs = []reflect.Value{reflect.ValueOf(args)}
	}

	// call the method
	r := m.Call(callArgs)
	*resp = r[0].Interface().(string)   // by convention, response is 1st return value
	errStr := r[1].Interface().(string) // by convention, error is 2nd return value
	if errStr != "" {
		err = os.NewError(errStr)
	}

	return
}
