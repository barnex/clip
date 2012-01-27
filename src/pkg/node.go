package clip

import(
	"strings"
)

// Node in the Library's file tree
type Node struct{
	data string
	parent *Node
	children []*Node
}

// Construct new node with given parent (possibly nil)
// and link the parent-child pointers
func NewNode(data string, parent *Node)*Node{
	node := &Node{str, parent, []*Node{}}
	if parent != nil{parent.children = append(parent.children, node)}
	return node
}


// 
func(n*Node)String()string{
	str := n.str
	for p := n.parent; p!=nil; p=p.parent{
		str = p.String() + "/" + str
	}
	return str
}

//func(n*Node)Child(str string)*Node{
//	for _,c:=range n.children{
//		if c.str == str{return c}
//	}
//	return nil
//}
