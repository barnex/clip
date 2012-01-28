package clip

import ()

// Stores a music Library
type Lib struct {
	*Node
}

// Constructs a new Library
func NewLib() *Lib {
	return &Lib{&Node{"", nil, nil}}
}


