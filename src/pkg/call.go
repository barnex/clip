package clip

import (
	"os"
)

type Call struct {
	args []string
	resp string
	err  os.Error
	resp chan *Call
}


func HandleCall(call *Call) {
	args := call.args

	if len(args) == 0 {
		args = []string{""}
	}
	cmd := args[0]
	args = args[1:]
	Debug("player.call", cmd, args)
	f, ok := command[cmd]
	if !ok {
		call.err = "no such command: " + cmd
	} else {
		call.resp, call.err = f(args)
	}

	call.respChan <- call
}
