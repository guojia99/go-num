package trees

type NodeData struct {
	Key   interface{} `json:"k,omitempty"`
	Value interface{} `json:"v,omitempty"`
}

type Node struct {
	Parent   *Node       `json:"p,omitempty"`
	Children []*Node     `json:"c,omitempty"`
	Data     []*NodeData `json:"d,omitempty"`
	Back     int64       `json:"b,omitempty"`
}

func (x *Node) Size() int {
	if x == nil {
		return 0
	}
	size := 1
	for _, child := range x.Children {
		size += child.Size()
	}
	return size
}

func (x *Node) Height() int {
	node := x
	height := 0
	for ; x != nil; node = node.Children[0] {
		height++
		if len(x.Children) == 0 {
			break
		}
	}
	return height
}

func (x *Node) AddOtherNodeChildren(otherNode *Node) {
	x.SetChildrenList(otherNode.Children)
}

func (x *Node) prependOtherNodeChildren(otherNode *Node) {
	children := append([]*Node(nil), otherNode.Children...)
	x.SetChildrenList(children)
}

func (x *Node) SetChildrenList(nodes []*Node) {
	x.Children = append([]*Node(nil), nodes...)
	for _, node := range nodes {
		node.Parent = x
	}
}

func (x *Node) AddData(data ...*NodeData) {
	x.Data = append(x.Data, data...)
}

func (x *Node) AddChildren(nodes ...*Node) {
	x.Children = append(x.Children, nodes...)
}

func (x *Node) DeleteChildren(index int) {
	if index > len(x.Children) {
		return
	}
	copy(x.Children[index:], x.Children[index+1:])
	x.Children[len(x.Children)-1] = nil
	x.Children = x.Children[:len(x.Children)-1]
}

func (x *Node) deleteData(index int) {
	copy(x.Data[index:], x.Data[index+1:])
	x.Data[len(x.Data)-1] = nil
	x.Data = x.Data[:len(x.Data)-1]
}
