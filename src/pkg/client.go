package clip

import (
	"rpc"
	"fmt"
	"time"
	"exec"
)

// Main loop for "client" mode (the normal mode).
func MainClient(args []string) {
	client := dialDaemon()
	var resp string
	err2 := client.Call("PlayerRPC.Call", args, &resp)
	Check(err2)
	fmt.Print(resp)
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
	cmd := exec.Command("clip", "-d")
	err := cmd.Start()
	Check(err)
}
