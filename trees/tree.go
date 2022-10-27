package trees

import "fmt"

type Tree interface {
	Search(node *Node, key interface{}) (index int, found bool)
	Set(key interface{}) (value interface{}, found bool)
	Put(key interface{}, value interface{})
	Remove(key interface{})
	Clear()

	Node(key interface{}) *Node
	Keys() []interface{}
	Values() []interface{}
	Left() *Node
	LeftKey() interface{}
	LeftValue() interface{}
	Right() *Node
	RightKey() interface{}
	RightValue() interface{}

	Empty() bool
	Size() int
	Height()

	fmt.Stringer
}
