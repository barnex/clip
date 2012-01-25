package clip

import(
	"strings"
)

// Stores a music Library
type Lib struct{
	Root *Node
}

// Constructs a new Library
func NewLib()*Lib{
	return &Lib{NewNode("", nil)}
}

func(lib*Lib)Add(str string){
	lib.Root.Add(str)
}


// Node in the Library's file tree
type Node struct{
	str string
	parent *Node
	children []*Node
}

func NewNode(str string, parent *Node)*Node{
	node := &Node{str, parent, []*Node{}}
	if parent != nil{parent.children = append(parent.children, node)}
	return node
}

func(n*Node)Add(str string){
	Debug("Node.Add", str)
	slash := strings.Index(str, "/")
	root := str[:slash]
	Debug("root=", root)
	base := str[slash+1:]
	Debug("base=", base)
	child := n.Child(root)
	if child == nil{
		child = NewNode(root, n)
	}
}

func(n*Node)String()string{
	str := n.str
	for p := n.parent; p!=nil; p=p.parent{
		str = p.String() + "/" + str
	}
	return str
}

func(n*Node)Child(str string)*Node{
	for _,c:=range n.children{
		if c.str == str{return c}
	}
	return nil
}
