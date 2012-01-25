package clip

// Stores a music Library
type Lib struct{
	Root *Node
}

// Constructs a new Library
func NewLib()*Lib{
	return &Lib{&Node{"/", nil, []*Node{}}}
}

// Node in the Library's file tree
type Node struct{
	str string
	parent *Node
	children []*Node
}


