package trees

type Position int8

const (
	Start Position = iota
	Between
	End
)

type Iterator interface {
	Next() bool
	Prev() bool
	First() bool
	Last() bool
	Key() interface{}
	Value() interface{}
}

func NewRbtreeIterator(t Tree) Iterator {
	return &RbtreeIterator{tree: t, node: nil, pos: Start}
}

type RbtreeIterator struct {
	tree Tree
	node *Node
	data *NodeData
	pos  Position
}

func (i *RbtreeIterator) Next() bool         { return i.goNext() }
func (i *RbtreeIterator) Prev() bool         { return i.goPrev() }
func (i *RbtreeIterator) Key() interface{}   { return i.data.Key }
func (i *RbtreeIterator) Value() interface{} { return i.data.Value }
func (i *RbtreeIterator) First() bool        { return i.goBetween().Next() }
func (i *RbtreeIterator) Last() bool         { return i.goEnd().Prev() }

func (i *RbtreeIterator) goNext() bool {
	if i.pos == End {
		return i.goEnd().pos == Between
	}

	if i.pos == Start {
		left := i.tree.Left()
		if left == nil {
			return i.goEnd().pos == Between
		}
		i.node, i.data = left, left.Data[0]
		return i.goBetween().pos == Between
	}

	fIndex, _ := i.tree.Search(i.node, i.data.Key)
	if fIndex+1 < len(i.node.Children) {
		i.node = i.node.Children[fIndex+1]
		for len(i.node.Children) > 0 {
			i.node = i.node.Children[0]
		}
		i.data = i.node.Data[0]
		return i.goBetween().pos == Between
	}

	if fIndex+1 < len(i.node.Data) {
		i.data = i.node.Data[fIndex+1]
		return i.goBetween().pos == Between
	}

	for i.node.Parent != nil {
		i.node = i.node.Parent
		pIndex, _ := i.tree.Search(i.node, i.data.Key)
		if pIndex < len(i.node.Data) {
			i.data = i.node.Data[pIndex]
			return i.goBetween().pos == Between
		}
	}
	return i.pos == Between
}

func (i *RbtreeIterator) goPrev() bool {
	if i.pos == Start {
		return i.goStart().pos == Between
	}

	if i.pos == End {
		right := i.tree.Right()
		if right == nil {
			return i.goStart().pos == Between
		}
		i.node, i.data = right, right.Data[len(right.Data)-1]
		return i.goBetween().pos == Between
	}

	fIndex, _ := i.tree.Search(i.node, i.data.Key)
	if fIndex < len(i.node.Children) {
		i.node = i.node.Children[fIndex]
		for len(i.node.Children) > 0 {
			i.node = i.node.Children[len(i.node.Children)-1]
		}
		i.data = i.node.Data[len(i.node.Data)-1]
		return i.goBetween().pos == Between
	}

	if fIndex-1 >= 0 {
		i.data = i.node.Data[fIndex-1]
		return i.goBetween().pos == Between
	}

	for i.node.Parent != nil {
		i.node = i.node.Parent
		pIndex, _ := i.tree.Search(i.node, i.data.Key)
		if pIndex-1 >= 0 {
			i.data = i.node.Data[pIndex-1]
			return i.goBetween().pos == Between
		}
	}
	return i.pos == Between
}

func (i *RbtreeIterator) goStart() *RbtreeIterator   { i.node, i.data, i.pos = nil, nil, Start; return i }
func (i *RbtreeIterator) goBetween() *RbtreeIterator { i.pos = Between; return i }
func (i *RbtreeIterator) goEnd() *RbtreeIterator     { i.node, i.data, i.pos = nil, nil, End; return i }
