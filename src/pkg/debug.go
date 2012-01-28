package clip

import (
	"os"
	"fmt"
)

func Debug(msg ...interface{}) {
	fmt.Println(msg...)
}


func Err(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
	os.Exit(3)
}
