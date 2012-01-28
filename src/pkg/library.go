package clip

import (
	"io"
	"os"
)

// Stores a music Library
type Lib struct {
	*Node
}

// Constructs a new Library
func NewLib() *Lib {
	return &Lib{&Node{"", nil, nil}}
}

func (lib *Lib) WriteTo(out io.Writer) (n int, err os.Error) {
	printf := func(node *Node) {
		n1, err1 := node.WriteTo(out)
		n += n1
		if err1 != nil {
			err = err1
		}
	}
	lib.Walk(printf)
	return
}
