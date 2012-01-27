package clip

import ()

// Stores a music Library
type Lib struct {
	root *Node
}

// Constructs a new Library
func NewLib() *Lib {
	return &Lib{NewNode("", nil)}
}

func (lib *Lib) Add(file string) {
	lib.root.Add(file)
}
