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

func(lib*Lib)String()string{
	n := lib.root
	str = n.file
	for _,c:=range n.children{}
}
