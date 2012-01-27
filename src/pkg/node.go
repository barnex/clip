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

// Construct new node with given parent (possibly nil)
// and link the parent-child pointers
func NewNode(file string, parent *Node) *Node {
	node := &Node{file, parent, []*Node{}}
	if parent != nil {
		parent.children = append(parent.children, node)
	}
	Debug("NewNode=", node)
	return node
}


// 
func (n *Node) String() string {
	str := n.file
	for p := n.parent; p != nil; p = p.parent {
		str = p.file + "/" + str
	}
	return str
}


func (n *Node) Add(path string) {
	Debug("Node.Add", path)
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	slash := strings.Index(path, "/")
	root := path
	base := ""
	if slash != -1{
		root = path[:slash]
		base = path[slash+1:]
	}
	Debug("root=", root)
	Debug("base=", base)
	child := n.Child(root)
	if child == nil {
		child = NewNode(root, n)
	}
	if base != ""{child.Add(base)}
}


func (n *Node) Child(file string) *Node {
	for _, c := range n.children {
		if c.file == file {
			return c
		}
	}
	return nil
}
