package clip

import (
	"strings"
)

// Node in the Library's file tree
type Node struct {
	file     string
	parent   *Node
	children []*Node
}


// Walks through the tree and applies function f to each Node
func (n *Node) Walk(f func(*Node)) {
	f(n)
	for _, c := range n.children {
		c.Walk(f)
	}
}

// Construct new node with given parent (possibly nil)
// and link the parent-child pointers
func (parent *Node) NewChild(file string) (child *Node) {
	Debug("NewNode", file, parent)
	child = &Node{file, parent, nil}
	if parent != nil {
		parent.children = append(parent.children, child)
		Debug("node.parent.children=", parent.children)
	}
	Debug("NewNode=", child)
	return
}


// 
//func (n *Node) String() string {
//	str := n.file
//	for p := n.parent; p != nil; p = p.parent {
//		str = p.file + "/" + str
//	}
//	return str
//}


func (n *Node) Add(path string) {
	Debug("Node.Add", path)
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	slash := strings.Index(path, "/")
	root := path
	base := ""
	if slash != -1 {
		root = path[:slash]
		base = path[slash+1:]
	}
	Debug("root=", root)
	Debug("base=", base)
	child := n.Child(root)
	if child == nil {
		child = n.NewChild(root)
	}
	if base != "" {
		child.Add(base)
	}
}


func (n *Node) Child(file string) *Node {
	for _, c := range n.children {
		if c.file == file {
			return c
		}
	}
	return nil
}
