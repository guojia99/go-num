package iterators

import "github.com/guojia99/go-num/trees"

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

func NewRbtreeIterator(t trees.Tree) Iterator {
	return &RbtreeIterator{tree: t, node: nil, pos: Start}
}
