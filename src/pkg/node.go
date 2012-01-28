package clip

import (
	"strings"
	"fmt"
	"io"
	"os"
)

// Node in the Library's file tree.
type Node struct {
	file     string
	parent   *Node
	children []*Node
}


// Walks through the tree and applies function f to each Node.
func (this *Node) Walk(f func(*Node)) {
	f(this)
	for _, c := range this.children {
		c.Walk(f)
	}
}

// Construct new node with given parent
// and link the parent-child pointers.
func (parent *Node) NewChild(file string) (child *Node) {
	child = &Node{file, parent, nil}
	parent.children = append(parent.children, child)
	return
}


// Returns full path represented by this node.
func (n *Node) String() string {
	str := n.file
	for p := n.parent; p != nil; p = p.parent {
		str = p.file + str
	}
	return str
}

// Write full path to out.
func (this *Node) WriteTo(out io.Writer) (n int, err os.Error) {
	n, err = fmt.Fprintln(out, this)
	return
}

// Add a slash-separated path to the tree.
func (n *Node) Add(path string) {
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
	child := n.Child(root)
	if child == nil {
		child = n.NewChild(root)
	}

	// recursively add base
	if base != "" {
		child.Add(base)
	}
}

// Get a child by its file string.
func (n *Node) Child(file string) *Node {
	for _, c := range n.children {
		if c.file == file {
			return c
		}
	}
	return nil
}
