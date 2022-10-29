package trees

import "fmt"

type Tree interface {
	Search(node *Node, key interface{}) (index int, found bool)
	Get(key interface{}) (value interface{}, found bool)
	Set(key interface{}, value interface{})
	Remove(key interface{})
	Clear()

	Keys() []interface{}
	Values() []interface{}
	Left() *Node
	Right() *Node

	Empty() bool
	Size() int
	Height() int

	fmt.Stringer
}

func NewBtree(deep int) Tree { return newBtree(deep) }
