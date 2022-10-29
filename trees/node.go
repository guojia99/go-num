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

func (n *Node) AddOtherNodeChildren(otherNode *Node) {
	n.SetChildrenList(otherNode.Children)
}

func (n *Node) prependOtherNodeChildren(otherNode *Node) {
	children := append([]*Node(nil), otherNode.Children...)
	n.SetChildrenList(children)
}

func (n *Node) SetChildrenList(nodes []*Node) {
	if n.Children == nil {
		n.Children = make([]*Node, 0)
	}
	n.Children = append(n.Children, nodes...)
	for _, node := range nodes {
		node.Parent = n
	}
}

func (n *Node) AddData(data ...*NodeData) {
	n.Data = append(n.Data, data...)
}

func (n *Node) AddChildren(nodes ...*Node) {
	n.Children = append(n.Children, nodes...)
}

func (n *Node) DeleteChildren(index int) {
	if index > len(n.Children) {
		return
	}
	copy(n.Children[index:], n.Children[index+1:])
	n.Children[len(n.Children)-1] = nil
	n.Children = n.Children[:len(n.Children)-1]
}

func (n *Node) deleteData(index int) {
	copy(n.Data[index:], n.Data[index+1:])
	n.Data[len(n.Data)-1] = nil
	n.Data = n.Data[:len(n.Data)-1]
}
