package clip

import (
	"os"
)

func init(){
	command["add"] = Add
}

func Add(args []string)(resp string, err os.Error){
	for _, arg := range args{
		library.AddPath(arg)
	}
	return
}
