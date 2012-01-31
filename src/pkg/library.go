package clip

// This file implements the Library data structure.

import (
	"io"
	"os"
)

// Stores a music Library
type Lib struct {
	fs     *Node // filesystem root
	//lookup []Item
}

//type Item struct {
//	name string
//	file *Node
//}

// Constructs a new Library
func NewLib() *Lib {
	return &Lib{&Node{"/", nil, nil}}
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
