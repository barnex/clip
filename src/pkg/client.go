package clip

import (
	"rpc"
	"fmt"
	"time"
	"exec"
)

func MainClient(args []string) {
	client := dialDaemon()
	var resp string
	err2 := client.Call("Daemon.Call", args, &resp)
	Check(err2)
	fmt.Println(resp)
}

func dialDaemon() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", "localhost"+PORT)
	if err != nil {
		forkDaemon()
	}
	trials := 0
	for err != nil && trials < 10 {
		client, err = rpc.DialHTTP("tcp", "localhost"+PORT)
		time.Sleep(100e6)
		trials++
	}
	Check(err)
	return client
}


func forkDaemon() {
	cmd := exec.Command("clip", "-d")
	err := cmd.Start()
	Check(err)
}
