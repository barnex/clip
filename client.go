package main

// This file implements the client main function.
// Invoked whenever the user types executes "clip".
// The client merely forwards the CLI arguments
// to the clip daemon and returns the response to
// the user.

import (
	"os"
	"net/rpc"
	"fmt"
	"time"
	"os/exec"
	"strings"
)

// RPC port
const port = ":2527"

// Main loop for "client" mode (the normal mode).
// Simply passes the arguments to the daemon and
// displays the result.
func MainClient(args []string) {
	client := dialDaemon()
	var resp string
	err := client.Call("RPC.Call", args, &resp)
	if err != nil {
		fmt.Fprint(os.Stderr, cleanup(err.Error()))
	}
	fmt.Print(cleanup(resp))
}

// cleanup newlines so string can be printed to stdout without redundant/missing newlines
func cleanup(str string) string {
	str = strings.Trim(str, "\n")
	if str != "" {
		return str + "\n"
	}
	return str
}

// Connect to the clip daemon for RPC communication.
// Starts the daemon if he's not yet running.
func dialDaemon() *rpc.Client {
	// try to call the daemon
	client, err := rpc.DialHTTP("tcp", "localhost"+port)

	// if daemon does not seem to be running, start him.
	const SLEEP = 10e6 // nanoseconds
	if err != nil {
		forkDaemon()
		time.Sleep(SLEEP)
	}

	// try again to call the daemon,
	// give him some time to come up.
	trials := 0
	for err != nil && trials < 10 {
		client, err = rpc.DialHTTP("tcp", "localhost"+port)
		time.Sleep(SLEEP)
		trials++
	}

	Check(err)
	return client
}

// Start the clip daemon.
func forkDaemon() {
	executable, err := os.Readlink("/proc/self/exe")
	Check(err)
	cmd := exec.Command(executable, "-d")
	err = cmd.Start()
	Check(err)
}

