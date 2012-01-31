package clip

// This file implements the Library data structure.

import (
	"io"
	"os"
	"strings"
)

// Stores a music Library
type Lib struct {
	fs     *Node // filesystem root
	items ItemArray
}
// Constructs a new Library
func NewLib() *Lib {
	return &Lib{&Node{"/", nil, nil}, []Item{}}
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
			lib.AddPath(lib.fs, arg)
			lib.Import(arg + "/" + f)
		}
		return
	}

	if info.IsRegular() {
		lib.AddPath(lib.fs, arg)
		return
	}
}

// Add a slash-separated path to the tree.
func (lib*Lib) AddPath(node *Node, path string) {
	// remove leading slash from path,
	// root node is already present
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	// split path into root and base 
	slash := strings.Index(path, "/")
	root, base := path, ""
	if slash != -1 {
		root, base = path[:slash+1], path[slash+1:]
	}

	// add root as a new child if not yet present
	child := node.Child(root)
	if child == nil {
		child = node.NewChild(root)
	}

	// recursively add base
	if base != "" {
		lib.AddPath(child, base)
	}
}


func(lib*Lib)Lookup(item string)*Node{
	return lib.items.Lookup(item)
}


// Print the entire library recursively
func (lib *Lib) WriteTo(out io.Writer) (n int, err os.Error) {
	printf := func(node *Node) {
		n1, err1 := node.WriteTo(out)
		n += n1
		if err1 != nil {
			err = err1
		}
	}
	lib.fs.Walk(printf)
	return
}
