/*
Algorithm	Average	Worst case
Space	O(n)	O(n)
search	O(log n)	O(log n)
Insert	O(log n)	O(log n)
Delete	O(log n)	O(log n)

Copy to References: https://en.wikipedia.org/wiki/B-tree
- Every node has at most m children.
- Every internal node has at least âm/2â children.
- Every non-leaf node has at least two children.
- All leaves appear on the same level and carry no information.
- A non-leaf node with k children contains kâ1 keys.
*/

package trees

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/guojia99/go-num/trees/utils"
)

const minChildrenNum = 3

type BTree struct {
	Root *Node `json:"root,omitempty"`
	// size total number of key.
	Size int64 `json:"size,omitempty"`
	// deep max children number.
	Deep int64 `json:"deep,omitempty"`
}

func newBtree(deep int64) *BTree {
	if deep < minChildrenNum {
		deep = minChildrenNum
	}
	var b = new(BTree)
	b.Deep = deep
	return b
}

func (x *BTree) Search(node *Node, key interface{}) (index int, found bool) {
	var mid int
	low, high := 0, len(node.Data)-1
	for low <= high {
		mid = (high + low) / 2
		comp := utils.Compare(key, node.Data[mid].Key)
		switch comp {
		case utils.MoreThan:
			low = mid + 1
		case utils.Less:
			high = mid - 1
		case utils.Equal:
			return mid, true
		case utils.Different:
			return low, false
		}
	}
	return low, false
}

func (x *BTree) Get(key interface{}) (value interface{}, found bool) {
	node, index, ok := x.searchForm(x.Root, key)
	if ok {
		return node.Data[index].Value, true
	}
	return nil, false
}

func (x *BTree) GetNode(key interface{}) (node *Node) {
	node, _, _ = x.searchForm(x.Root, key)
	return
}

func (x *BTree) Set(key interface{}, value interface{}) {

	x.set(&NodeData{Key: key, Value: value})
}

func (x *BTree) Sets(datas []*NodeData) {
	for _, data := range datas {
		x.set(data)
	}
}

func (x *BTree) set(data *NodeData) {
	if x.Root == nil {
		x.Root = &Node{
			Children: make([]*Node, 0),
			Data:     []*NodeData{data},
			Back:     0,
		}
		x.Size += 1
		return
	}
	if x.insert(x.Root, data) {
		x.Size++
	}
}

func (x *BTree) Remove(key interface{}) {
	node, index, ok := x.searchForm(x.Root, key)
	if !ok {
		return
	}
	x.Size--
	if x.isLeaf(node) {
		dKey := node.Data[index].Key
		node.deleteData(index)
		x.balance(node, dKey)
		if len(x.Root.Data) == 0 {
			x.Root = nil
		}
		return
	}
	leftNode := x.rightNode(node.Children[index])
	leftDataIndex := len(leftNode.Data) - 1
	node.Data[index] = leftNode.Data[leftDataIndex]
	dKey := leftNode.Data[leftDataIndex].Key // back
	leftNode.deleteData(leftDataIndex)
	x.balance(leftNode, dKey)
}
func (x *BTree) Clear() { x.Root, x.Size = nil, 0 }

func (x *BTree) Keys() []interface{} {
	keys := make([]interface{}, x.Size)
	iterator := NewRbtreeIterator(x)
	for i := 0; iterator.Next(); i++ {
		keys[i] = iterator.Key()
	}
	return keys
}

func (x *BTree) Values() []interface{} {
	values := make([]interface{}, x.Size)
	iterator := NewRbtreeIterator(x)
	for i := 0; iterator.Next(); i++ {
		values[i] = iterator.Value()
	}
	return values
}

func (x *BTree) Left() *Node  { return x.leftNode(x.Root) }
func (x *BTree) Right() *Node { return x.rightNode(x.Root) }
func (x *BTree) Empty() bool  { return x.Size == 0 }
func (x *BTree) Length() int  { return int(x.Size) }
func (x *BTree) Height() int  { return x.Root.Height() }

func (x *BTree) Println() string {
	var buffer = new(bytes.Buffer)
	buffer.WriteString("B-Tree:")
	if !x.Empty() {
		x.stringLoop(buffer, x.Root, 0)
	}
	return buffer.String()
}

func (x *BTree) stringLoop(buffer *bytes.Buffer, node *Node, loopNum int) {
	for i := 0; i < len(node.Data)+1; i++ {
		if i < len(node.Children) {
			x.stringLoop(buffer, node.Children[i], loopNum)
		}
		if i < len(node.Data) {
			buffer.WriteString(strings.Repeat("\t", loopNum))
			buffer.WriteString(fmt.Sprintf("%+v", node.Data[loopNum].Key) + "\n")
		}
	}
}

func (x *BTree) leftNode(base *Node) *Node {
	if x.Empty() {
		return nil
	}
	cur := base
	for {
		if x.isLeaf(cur) {
			return cur
		}
		cur = cur.Children[0]
	}
}

func (x *BTree) rightNode(base *Node) *Node {
	if x.Empty() {
		return nil
	}
	cur := base
	for {
		if x.isLeaf(cur) {
			return cur
		}
		cur = cur.Children[len(cur.Children)-1]
	}
}

func (x *BTree) isLeaf(node *Node) bool { return len(node.Children) == 0 }

func (x *BTree) insert(node *Node, data *NodeData) bool {
	if x.isLeaf(node) {
		return x.insertIntoLeaf(node, data)
	}
	return x.insertIntoInternal(node, data)
}

func (x *BTree) insertIntoLeaf(node *Node, data *NodeData) bool {
	inIndex, ok := x.Search(node, data.Key)
	if ok {
		node.Data[inIndex] = data
		return false
	}

	node.Data = append(node.Data, nil)
	copy(node.Data[inIndex+1:], node.Data[inIndex:])
	node.Data[inIndex] = data
	x.organizeNodes(node)
	return true
}

func (x *BTree) insertIntoInternal(node *Node, data *NodeData) bool {
	inIndex, ok := x.Search(node, data.Key)
	if ok {
		node.Data[inIndex] = data
		return false
	}
	return x.insert(node.Children[inIndex], data)
}

// organizeNodes çĉ
func (x *BTree) organizeNodes(node *Node) {
	if !(len(node.Data) > int(x.Deep-1)) {
		return
	}

	middle := (x.Deep - 1) / 2

	// set root
	if node == x.Root {
		left := &Node{Data: append([]*NodeData(nil), x.Root.Data[:middle]...)}
		// old: left := &Node{Data: x.Root.Data[:middle]}
		right := &Node{Data: append([]*NodeData(nil), x.Root.Data[middle+1:]...)}
		if !x.isLeaf(x.Root) {
			left.SetChildrenList(x.Root.Children[:middle+1])
			right.SetChildrenList(x.Root.Children[middle+1:])
		}
		root := &Node{
			Data:     []*NodeData{x.Root.Data[middle]},
			Children: []*Node{left, right},
		}
		left.Parent, right.Parent = root, root
		x.Root = root
		return
	}
	// set not root
	parent := node.Parent
	left := &Node{Data: append([]*NodeData(nil), node.Data[:middle]...), Parent: parent}
	right := &Node{Data: append([]*NodeData(nil), node.Data[middle+1:]...), Parent: parent}
	if !x.isLeaf(node) {
		left.SetChildrenList(node.Children[:middle+1])
		right.SetChildrenList(node.Children[middle+1:])
	}

	inIndex, _ := x.Search(parent, node.Data[middle].Key)

	parent.AddData(nil)
	copy(parent.Data[inIndex+1:], parent.Data[inIndex:])
	parent.Data[inIndex] = node.Data[middle]
	parent.Children[inIndex] = left

	parent.AddChildren(nil)
	copy(parent.Children[inIndex+2:], parent.Children[inIndex+1:])
	parent.Children[inIndex+1] = right

	x.organizeNodes(parent)
}

// searchForm ä?ĉ?ä¸ċçnodeċkeyċŻğĉ?ċ°ä½ç½?ċç¸ċşĉ°ĉ?
func (x *BTree) searchForm(startNode *Node, key interface{}) (node *Node, index int, found bool) {
	if x.Empty() {
		return nil, -1, false
	}

	node = startNode
	for {
		index, found = x.Search(node, key)
		if found {
			return node, index, true
		}
		if x.isLeaf(node) {
			return nil, -1, false
		}
		node = node.Children[index]
	}
}

// balance èŞċı³èĦĦĉ è§£ĉ
/*
- 	ċĤĉçĵşċ°ċç´ èçıçċ³ċċĵċ­ċ¨ä¸ĉ?ĉċ¤ä½çċç´ ïĵé£äıċċ·Ĥĉè½Ĵċ°çĥèçıçċéċĵċ¤ċĥċ°çĵşċ°ċç´ èçıçĉċïĵċéċĵè˘Ğç§ğä¸ĉ?ïĵçĵşċ°ċç´ çèçıç°ċ¨ĉĉċ°ĉ°éçċç´ ïĵ
	ċ°çĥèçıçċéċĵĉżĉ˘ä¸şċ³ċċĵççĴĴä¸ä¸Şċç´ ïĵċ³ċċĵċ¤ħċğäşä¸ä¸Şèçıä½äğçĥĉ?ĉĉċ°ĉ°éçċç´ ïĵĉ ċéĉ°ċı³èĦĦ
-   ċĤċïĵċĤĉçĵşċ°ċç´ èçıçċ·Ĥċċĵċ­ċ¨ä¸ĉ?ĉċ¤ä½çċç´ ïĵé£äıċċ³ĉè½Ĵċ°çĥèçıçċéċĵċ¤ċĥċ°çĵşċ°ċç´ èçıççĴĴä¸ä¸Şèçıïĵċéċĵè˘Ğç§ğä¸ĉ?ïĵçĵşċ°ċç´ çèçıç°ċ¨ĉĉċ°ĉ°éçċç´ ïĵ
	ċ°çĥèçıçċéċĵĉżĉ˘ä¸şċ·Ĥċċĵçĉċä¸ä¸Şċç´ ïĵċ·Ĥċċĵċ¤ħċğäşä¸ä¸Şèçıä½äğçĥĉ?ĉĉċ°ĉ°éçċç´ ïĵĉ ċéĉ°ċı³èĦĦ
-   ċĤċïĵċĤĉċ?çä¸¤ä¸Şç´ĉ?ċċĵèçıé½ċŞĉĉċ°ĉ°éçċç´ ïĵé£äıċ°ċ?ä¸ä¸ä¸Şç´ĉ?ċċĵèçıäğ?ċçĥèçıä¸­ċ?äğĴçċéċĵċċıĥċ°ċéċĵċ¤ċĥċ°ċ·Ĥè?ıçèçıïĵċ·Ĥè?ıçèçıċŻäğ?ĉŻçĵşċ°ċç´ çèçıĉèĉ?ĉĉċ°ĉ°éċç´ çċċĵèçıïĵ
	ċ°ċ³è?ıèçıä¸­ĉĉçċç´ ç§ğċ¨ċ°ċ·Ĥè?ıèçıïĵċ·Ĥè?ıèçıç°ċ¨ĉ?ĉĉċ¤§ĉ°éçċç´ ïĵċ³è?ıèçıä¸şçİşïĵ, ċ°çĥèçıä¸­çċéċĵċçİşçċ³ċ­ĉ ç§ğé¤ïĵçĥèçıċ¤ħċğäşä¸ä¸Şċç´ ïĵ
-   ċĤĉçĥèçıĉŻĉ ıèçıċıĥä¸ĉ²Ħĉċç´ äşïĵé£äıéĉ?ċ?ċıĥä¸è?İċċıĥäıċçèçıĉä¸şĉ°çĉ ıèçıïĵĉ çĉ·ħċşĤċċ°ïĵċĤċïĵċĤĉçĥèçıçċç´ ĉ°éċ°äşĉċ°ċĵïĵéĉ°ċı³èĦĦçĥèçı
*/
func (x *BTree) balance(node *Node, key interface{}) {
	if node == nil || len(node.Data) >= x.minChildren()-1 {
		return
	}

	leftBrother, leftBrotherIdx := x.leftBrother(node, key)
	if leftBrother != nil && len(leftBrother.Data) > x.minChildren()-1 {
		node.Data = append([]*NodeData{node.Parent.Data[leftBrotherIdx]}, node.Data...)
		node.Parent.Data[leftBrotherIdx] = leftBrother.Data[len(leftBrother.Data)-1]
		leftBrother.deleteData(len(leftBrother.Data) - 1)
		if !x.isLeaf(leftBrother) {
			leftBrotherRightChild := leftBrother.Children[len(leftBrother.Children)-1]
			leftBrotherRightChild.Parent = node
			node.Children = append([]*Node{leftBrotherRightChild}, node.Children...)
			leftBrother.DeleteChildren(len(leftBrother.Children) - 1)
		}
	}

	rightBrother, rightBrotherIdx := x.rightBrother(node, key)
	if rightBrother != nil && len(rightBrother.Data) > x.minChildren()-1 {
		node.Data = append(node.Data, node.Parent.Data[rightBrotherIdx-1])
		node.Parent.Data[rightBrotherIdx-1] = rightBrother.Data[0]
		rightBrother.deleteData(0)
		if !x.isLeaf(rightBrother) {
			rightBrotherLeftChild := rightBrother.Children[0]
			rightBrotherLeftChild.Parent = node
			node.Children = append(node.Children, rightBrotherLeftChild)
			rightBrother.DeleteChildren(0)
		}
		return
	}

	if rightBrother != nil {
		node.Data = append(node.Data, node.Parent.Data[rightBrotherIdx-1])
		node.Data = append(node.Data, rightBrother.Data...)
		key = node.Parent.Data[rightBrotherIdx-1].Key
		node.Parent.deleteData(rightBrotherIdx - 1)
		node.Parent.Children[rightBrotherIdx].AddOtherNodeChildren(node)
		node.Parent.DeleteChildren(rightBrotherIdx)
	} else if leftBrother != nil {
		data := append([]*NodeData(nil), leftBrother.Data...)
		data = append(data, node.Parent.Data[leftBrotherIdx])
		node.Data = append(data, node.Data...)
		key = node.Parent.Data[leftBrotherIdx].Key

		node.Parent.deleteData(leftBrotherIdx)
		node.Parent.Children[leftBrotherIdx].prependOtherNodeChildren(node)
		node.Parent.DeleteChildren(leftBrotherIdx)
	}

	if node.Parent == x.Root && len(x.Root.Data) == 0 {
		x.Root = node
		node.Parent = nil
		return
	}
	x.balance(node.Parent, key)
}

func (x *BTree) leftBrother(node *Node, key interface{}) (brother *Node, brotherIndex int) {
	if node.Parent != nil {
		brotherIndex, _ = x.Search(node.Parent, key)
		brotherIndex--
		if brotherIndex >= 0 && brotherIndex < len(node.Parent.Children) {
			return node.Parent.Children[brotherIndex], brotherIndex
		}
	}
	return nil, -1
}
func (x *BTree) rightBrother(node *Node, key interface{}) (brother *Node, brotherIndex int) {
	if node.Parent != nil {
		brotherIndex, _ = x.Search(node.Parent, key)
		brotherIndex--
		if brotherIndex < len(node.Parent.Children) {
			return node.Parent.Children[brotherIndex], brotherIndex
		}
	}
	return nil, -1
}

func (x *BTree) maxChildren() int { return int(x.Deep) }
func (x *BTree) minChildren() int { return int((x.Deep + 1) / 2) }
