package trees

import (
	"fmt"
	"reflect"
	"testing"
)

func getNode(key int) *NodeData {
	return &NodeData{key, key * key}
}

// TestBTree_Get_OrderedInput 有序输入
func TestBTree_Get_OrderedInput(t *testing.T) {
	ts := NewBtree(5)
	ts.Sets([]*NodeData{
		{1, 1 * 1},
		{2, 2 * 2},
		{3, 3 * 3},
		{4, 4 * 4},
		{5, 5 * 5},
		{1000 * 1000, 1},
	})

	tests := []struct {
		key       interface{}
		wantValue interface{}
		wantFound bool
	}{
		{-1, nil, false},
		{1, 1, true},
		{2, 2 * 2, true},
		{3, 3 * 3, true},
		{4, 4 * 4, true},
		{1000 * 1000, 1, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("get%d", tt.key), func(t *testing.T) {
			gotValue, gotFound := ts.Get(tt.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

// TestBtree_Get_UnorderedInput 无序输入
func TestBtree_Get_UnorderedInput(t *testing.T) {
	ts := NewBtree(5)
	ts.Sets([]*NodeData{
		{100, 100},
		{29, 29},
		{188, 188},
		{6, 6},
		{10, 10},
		{19, 10},
		{19, 19},
		{11, 11},
	})

	tests := []struct {
		key       interface{}
		wantValue interface{}
		wantFound bool
	}{
		{-1, nil, false},
		{100, 100, true},
		{29, 29, true},
		{188, 188, true},
		{6, 6, true},
		{10, 10, true},
		{19, 19, true},
		{11, 11, true},
		{12, nil, false},
		{-111, nil, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("get%d", tt.key), func(t *testing.T) {
			gotValue, gotFound := ts.Get(tt.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
func newBaseTestBtree() Tree {
	// BTree is like:
	// https://www.cs.usfca.edu/~galles/visualization/BTree.html
	/*
					   8
		          /	    	\
				4		  	  12
			  /   \     	/    \
			2		6		10    14
		   / \     / \     / \    / \
		  1   3   5  7    9  11  13  15
	*/
	ts := NewBtree(3)
	ts.Sets([]*NodeData{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
		{10, 10},
		{11, 11},
		{12, 12},
		{13, 13},
		{14, 14},
		{15, 15},
	})
	return ts
}

func newLoopTestBtree(deep int, loop int) Tree {
	ts := NewBtree(int64(deep))
	for i := 0; i < loop; i++ {
		ts.Set(i, i)
	}
	return ts
}

func TestBTree_GetNode_Size(t *testing.T) {
	ts := newBaseTestBtree()
	/*
					   8
		          /	    	\
				4		  	  12
			  /   \     	/    \
			2		6		10    14
		   / \     / \     / \    / \
		  1   3   5  7    9  11  13  15
	*/
	tests := []struct {
		key      int
		wantSize int
	}{
		{8, 15},
		{4, 7},
		{2, 3},
		{1, 1},
		{3, 1},
		{6, 3},
		{5, 1},
		{7, 1},
		{12, 7},
		{10, 3},
		{9, 1},
		{11, 1},
		{14, 3},
		{13, 1},
		{15, 1},
		{1000000, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("get%d", tt.key), func(t *testing.T) {
			gotValue := ts.GetNode(tt.key).Size()
			if !reflect.DeepEqual(gotValue, tt.wantSize) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantSize)
			}
		})
	}
}

func TestBTree_Remove(t *testing.T) {
	ts := newBaseTestBtree()
	tests := []struct {
		key int
	}{
		{key: 10},
		{key: 1},
		{key: 11},
		{key: 20},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("romove_key-%d", tt.key), func(t *testing.T) {
			ts.Remove(tt.key)
			gotValue, gotFound := ts.Get(tt.key)
			if gotFound {
				t.Errorf("want value is `nil`, gotValue %+v", gotValue)
			}
		})
	}
}
