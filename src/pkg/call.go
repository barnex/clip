package clip


import (
	"os"
)


// Represents a call sent through RPC.
type Call struct {
	Args     []string  // CLI args (e.g. "play jazz")
	respChan chan Resp // send response here
}

// Response to a call.
type Resp struct {
	Resp string   // Response to user (e.g. "playing jazz")
	Err  os.Error // Error to user (e.g. "jazz not found")
}

func NewCall(args []string) Call {
	return Call{args, make(chan Resp)}
}


// Parse and execute call, return response.
func (call *Call) Exec() Resp {
	args := call.Args

	if len(args) == 0 {
		args = []string{""}
	}
	cmd := args[0]
	args = args[1:]
	Debug("player.call", cmd, args)
	f, ok := command[cmd]
	var resp Resp
	if !ok {
		resp.Err = os.NewError("no such command: " + cmd)
	} else {
		resp.Resp, resp.Err = f(args)
	}
	return resp
}

