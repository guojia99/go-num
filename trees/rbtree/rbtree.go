package rbtree

import "github.com/guojia99/go-num/trees"

const minChildrenNum = 3

type RbTree struct {
	root        *trees.Node
	size        int // total number of key
	maxChildren int // max children number
}
