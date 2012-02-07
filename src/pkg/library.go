package clip

// This file implements the Library data structure.

import (
	"os"
	"io"
	"fmt"
	"strings"
)

// Stores a music Library
type Lib struct {
	items ItemArray
}

// Constructs a new Library
func NewLib() *Lib {
	lib := new(Lib)
	lib.init()
	return lib
}

// Initializes the library
func (lib*Lib)init(){
	lib.items =  ItemArray([]*Item{})
}

// Recursively import directory or file into library.
func (lib *Lib) Import(arg string) {
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
		lib.items = append(lib.items, NewFile(arg))
		return
	}
}

// Print the entire library recursively
func (lib *Lib) WriteTo(out io.Writer) (n int, err os.Error) {
	for _, item := range lib.items {
		N, ERR := fmt.Fprintln(out, item)
		if ERR != nil {
			err = ERR
		}
		n += N
	}
	return
}

// Find items based on tag
func (lib *Lib) Find(tag string) (items []*Item) {
	tag = Fuzzy(tag)
	items = []*Item{}
	for _, item := range lib.items {
		if item.tag == tag {
			items = append(items, item)
			Debug("find", tag, ":", item.file)
		}
	}
	return
}
