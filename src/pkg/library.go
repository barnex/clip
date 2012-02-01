package clip

// This file implements the Library data structure.

import (
	"os"
	"fmt"
	"strings"
)

// Stores a music Library
type Lib struct {
}

// Constructs a new Library
func NewLib() *Lib {
	return &Lib{}
}

// Recursively import directory or file into library.
func(lib*Lib) Import(arg string) {
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
			lib.Import(arg + "/" + f)
		}
		return
	}

	if info.IsRegular() {
		lib.ImportFile(arg)
		return
	}
}

func(lib*Lib)ImportFile(file string){
	fmt.Println("import", file)
}
