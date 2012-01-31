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
	lookup ItemArray
}
// Constructs a new Library
func NewLib() *Lib {
	return &Lib{&Node{"/", nil, nil}, []Item{}}
}

// Add a slash-separated path to the tree.
func AddPath(node *Node, path string) {
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
		AddPath(child, base)
	}
}


func (lib *Lib) AddFile(file string) {
	AddPath(lib.fs, file)
	//lib.lookup.Add(file)
}

func(lib*Lib)Lookup(item string)*Node{
	return lib.lookup.Lookup(item)
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
