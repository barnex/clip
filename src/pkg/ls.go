package clip

import (
	"os"
	"bytes"
)

func init(){
	command["ls"] = Ls
}

func Ls(args []string)(resp string, err os.Error){
	buf := bytes.NewBuffer([]byte{})
	library.WriteTo(buf)
	resp = buf.String()
	return
}
