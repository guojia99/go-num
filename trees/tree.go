package trees

type Tree interface {
	Search(node *Node, key interface{}) (index int, found bool)
	Get(key interface{}) (value interface{}, found bool)
	GetNode(key interface{}) (node *Node)
	Set(key interface{}, value interface{})
	Sets([]*NodeData)
	Remove(key interface{})
	Clear()

	Keys() []interface{}
	Values() []interface{}
	Left() *Node
	Right() *Node

	Empty() bool
	Length() int
	Height() int

	Println() string
}

func NewBtree(deep int64) Tree { return newBtree(deep) }
