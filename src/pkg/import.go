package clip

// This file implements the "import" command

import (
	"os"
	"strings"
)

// Register the command
func init() {
	command["import"] = Import
}

func Import(args []string) (resp string, err os.Error) {
	if len(args) == 0 {
		err = os.NewError("nothing specified, nothing imported")
		return
	}
	for _, arg := range args {
		importFile(arg)
	}
	return
}


func importFile(arg string) {
	// rm trailing slash
	if strings.HasSuffix(arg, "/") {
		arg = arg[:len(arg)-1]
	}

	info, err := os.Stat(arg)
	Check(err) // TODO: dontcrash

	if info.IsDirectory() {
		dir, err := os.OpenFile(arg, os.O_RDONLY, 0777)
		Check(err)
		files, err2 := dir.Readdirnames(-1)
		Check(err2)
		for _, f := range files {
			importFile(arg + "/" + f)
		}
		return
	}

	if info.IsRegular() {
		library.fs.AddPath(arg)
		return
	}
}
