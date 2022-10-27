package trees

type NodeData struct {
	Key   interface{} `json:"k"`
	Value interface{} `json:"v"`
}

type Node struct {
	Parent   *Node       `json:"p"`
	Children []*Node     `json:"c"`
	Data     []*NodeData `json:"d"`

	Back int8
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	size := 1
	for _, child := range n.Children {
		size += child.Size()
	}
	return size
}

func (n *Node) Height() int {
	node := n
	height := 0
	for ; n != nil; node = node.Children[0] {
		height++
		if len(n.Children) == 0 {
			break
		}
	}
	return height
}
