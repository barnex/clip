package clip

import (
	"rpc"
	"fmt"
)

func MainClient(args []string) {
	client, err1 := rpc.DialHTTP("tcp", "localhost"+PORT)
	Check(err1)
	var resp string
	err2 := client.Call("Daemon.Call", args, &resp)
	Check(err2)
	fmt.Println(resp)
}
