/*
Algorithm	Average	Worst case
Space	O(n)	O(n)
Search	O(log n)	O(log n)
Insert	O(log n)	O(log n)
Delete	O(log n)	O(log n)

Copy to References: https://en.wikipedia.org/wiki/B-tree
- Every node has at most m children.
- Every internal node has at least ⌈m/2⌉ children.
- Every non-leaf node has at least two children.
- All leaves appear on the same level and carry no information.
- A non-leaf node with k children contains k−1 keys.
*/

package trees

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/guojia99/go-num/trees/iterators"
	"github.com/guojia99/go-num/trees/utils"
)

const minChildrenNum = 3

func newBtree(deep int) *BTree {
	if deep < minChildrenNum {
		deep = minChildrenNum
	}
	return &BTree{deep: deep}
}

type BTree struct {
	root *Node
	size int // total number of key
	deep int // max children number
}

func (b *BTree) Search(node *Node, key interface{}) (index int, found bool) {
	mid, low, high := 0, 0, len(node.Data)-1
	for low <= high {
		mid = (high + low) / 2
		switch utils.Compare(key, node.Data[mid].Key) {
		case utils.Less:
			high = mid - 1
		case utils.Equal:
			return mid, true
		case utils.MoreThan:
			low = mid + 1
		case utils.Different:
			return low, false
		}
	}
	return low, false
}

func (b *BTree) Get(key interface{}) (value interface{}, found bool) {
	node, index, ok := b.searchForm(b.root, key)
	if ok {
		return node.Data[index].Value, true
	}
	return nil, false
}

func (b *BTree) Set(key interface{}, value interface{}) {
	data := &NodeData{key, value}
	if b.root == nil {
		b.root = &Node{
			Children: make([]*Node, 0),
			Data:     []*NodeData{data},
			Back:     0,
		}
		b.size += 1
		return
	}
	if b.insert(b.root, data) {
		b.size++
	}
}

func (b *BTree) Remove(key interface{}) {
	node, index, ok := b.searchForm(b.root, key)
	if !ok {
		return
	}
	b.size--
	if b.isLeaf(node) {
		dKey := node.Data[index].Key
		node.deleteData(index)
		b.balance(node, dKey)
		if len(b.root.Data) == 0 {
			b.root = nil
		}
		return
	}
	leftNode := b.rightNode(node.Children[index])
	leftDataIndex := len(leftNode.Data) - 1
	node.Data[index] = leftNode.Data[leftDataIndex]
	dKey := leftNode.Data[leftDataIndex].Key // back
	leftNode.deleteData(leftDataIndex)
	b.balance(leftNode, dKey)
}
func (b *BTree) Clear() { b.root, b.size = nil, 0 }

func (b *BTree) Keys() []interface{} {
	keys := make([]interface{}, b.size)
	iterator := iterators.NewRbtreeIterator(b)
	for i := 0; iterator.Next(); i++ {
		keys[i] = iterator.Key()
	}
	return keys
}

func (b *BTree) Values() []interface{} {
	values := make([]interface{}, b.size)
	iterator := iterators.NewRbtreeIterator(b)
	for i := 0; iterator.Next(); i++ {
		values[i] = iterator.Value()
	}
	return values
}

func (b *BTree) Left() *Node  { return b.leftNode(b.root) }
func (b *BTree) Right() *Node { return b.rightNode(b.root) }
func (b *BTree) Empty() bool  { return b.size == 0 }
func (b *BTree) Size() int    { return b.size }
func (b *BTree) Height() int  { return b.root.Height() }

func (b *BTree) String() string {
	var buffer = new(bytes.Buffer)
	buffer.WriteString("B-Tree:")
	if !b.Empty() {
		b.stringLoop(buffer, b.root, 0)
	}
	return buffer.String()
}

func (b *BTree) stringLoop(buffer *bytes.Buffer, node *Node, loopNum int) {
	for i := 0; i < len(node.Data)+1; i++ {
		if i < len(node.Children) {
			b.stringLoop(buffer, node.Children[i], loopNum)
		}
		if i < len(node.Data) {
			buffer.WriteString(strings.Repeat("\t", loopNum))
			buffer.WriteString(fmt.Sprintf("%+v", node.Data[loopNum].Key) + "\n")
		}
	}
}

func (b *BTree) leftNode(base *Node) *Node {
	if b.Empty() {
		return nil
	}
	cur := base
	for {
		if b.isLeaf(cur) {
			return cur
		}
		cur = cur.Children[0]
	}
}

func (b *BTree) rightNode(base *Node) *Node {
	if b.Empty() {
		return nil
	}
	cur := base
	for {
		if b.isLeaf(cur) {
			return cur
		}
		cur = cur.Children[len(cur.Children)-1]
	}
}

func (b *BTree) isLeaf(node *Node) bool { return len(node.Children) == 0 }

func (b *BTree) insert(node *Node, data *NodeData) bool {
	if b.isLeaf(node) {
		return b.insertIntoLeaf(node, data)
	}
	return b.insertIntoInternal(node, data)
}

func (b *BTree) insertIntoLeaf(node *Node, data *NodeData) bool {
	inIndex, ok := b.Search(node, data.Key)
	if ok {
		node.Data[inIndex] = data
		return false
	}

	node.Data = append(node.Data, nil)
	copy(node.Data[inIndex+1:], node.Data[inIndex:])
	node.Data[inIndex] = data
	b.organizeNodes(node)
	return true
}

func (b *BTree) insertIntoInternal(node *Node, data *NodeData) bool {
	inIndex, ok := b.Search(node, data.Key)
	if ok {
		node.Data[inIndex] = data
		return false
	}
	return b.insert(node.Children[inIndex], data)
}

// organizeNodes 理枝
func (b *BTree) organizeNodes(node *Node) {
	if !(len(node.Data) > b.deep-1) {
		return
	}

	middle := (b.deep - 1) / 2
	if node == b.root {
		left := &Node{Data: append([]*NodeData{nil}, b.root.Data[:middle]...)}
		right := &Node{Data: append([]*NodeData{nil}, b.root.Data[middle+1:]...)}
		if !b.isLeaf(b.root) {
			left.SetChildrenList(append([]*Node(nil), b.root.Children[:middle+1]...))
			right.SetChildrenList(append([]*Node(nil), b.root.Children[middle+1:]...))
		}
		root := &Node{
			Data:     []*NodeData{b.root.Data[middle]},
			Children: []*Node{left, right},
		}
		left.Parent, right.Parent = root, root
		b.root = root
		return
	}

	parent := node.Parent
	left := &Node{Data: append([]*NodeData{nil}, node.Data[:middle]...)}
	right := &Node{Data: append([]*NodeData{nil}, node.Data[middle+1:]...)}
	if !b.isLeaf(node) {
		left.SetChildrenList(append([]*Node(nil), node.Children[:middle+1]...))
		right.SetChildrenList(append([]*Node(nil), node.Children[middle+1:]...))
	}

	inIndex, _ := b.Search(parent, node.Data[middle].Key)

	parent.AddData(nil)
	copy(parent.Data[inIndex+1:], parent.Data[inIndex:])
	parent.Data[inIndex] = node.Data[middle]
	parent.Children[inIndex] = left

	parent.AddChildren(nil)
	copy(parent.Children[inIndex+2:], parent.Children[inIndex+1:])
	parent.Children[inIndex+1] = right

	b.organizeNodes(parent)
}

// searchForm 依据不同的node和key寻找到位置和相应数据
func (b *BTree) searchForm(startNode *Node, key interface{}) (node *Node, index int, found bool) {
	if b.Empty() {
		return nil, -1, false
	}

	node = startNode
	for {
		index, found = b.Search(node, key)
		if found {
			return node, index, true
		}
		if b.isLeaf(node) {
			return nil, -1, false
		}
		node = node.Children[index]
	}
}

// balance 自平衡树解构
/*
- 	如果缺少元素节点的右兄弟存在且拥有多余的元素，那么向左旋转将父节点的分隔值复制到缺少元素节点的最后（分隔值被移下来；缺少元素的节点现在有最小数量的元素）
	将父节点的分隔值替换为右兄弟的第一个元素（右兄弟失去了一个节点但仍然拥有最小数量的元素）树又重新平衡
-   否则，如果缺少元素节点的左兄弟存在且拥有多余的元素，那么向右旋转将父节点的分隔值复制到缺少元素节点的第一个节点（分隔值被移下来；缺少元素的节点现在有最小数量的元素）
	将父节点的分隔值替换为左兄弟的最后一个元素（左兄弟失去了一个节点但仍然拥有最小数量的元素）树又重新平衡
-   否则，如果它的两个直接兄弟节点都只有最小数量的元素，那么将它与一个直接兄弟节点以及父节点中它们的分隔值合并将分隔值复制到左边的节点（左边的节点可以是缺少元素的节点或者拥有最小数量元素的兄弟节点）
	将右边节点中所有的元素移动到左边节点（左边节点现在拥有最大数量的元素，右边节点为空）, 将父节点中的分隔值和空的右子树移除（父节点失去了一个元素）
-   如果父节点是根节点并且没有元素了，那么释放它并且让合并之后的节点成为新的根节点（树的深度减小）否则，如果父节点的元素数量小于最小值，重新平衡父节点
*/
func (b *BTree) balance(node *Node, key interface{}) {
	if node == nil || len(node.Data) >= b.minChildren()-1 {
		return
	}

	leftBrother, leftBrotherIdx := b.leftBrother(node, key)
	if leftBrother != nil && len(leftBrother.Data) > b.minChildren()-1 {
		node.Data = append([]*NodeData{node.Parent.Data[leftBrotherIdx]}, node.Data...)
		node.Parent.Data[leftBrotherIdx] = leftBrother.Data[len(leftBrother.Data)-1]
		leftBrother.deleteData(len(leftBrother.Data) - 1)
		if !b.isLeaf(leftBrother) {
			leftBrotherRightChild := leftBrother.Children[len(leftBrother.Children)-1]
			leftBrotherRightChild.Parent = node
			node.Children = append([]*Node{leftBrotherRightChild}, node.Children...)
			leftBrother.DeleteChildren(len(leftBrother.Children) - 1)
		}
	}

	rightBrother, rightBrotherIdx := b.rightBrother(node, key)
	if rightBrother != nil && len(rightBrother.Data) > b.minChildren()-1 {
		node.Data = append(node.Data, node.Parent.Data[rightBrotherIdx-1])
		node.Parent.Data[rightBrotherIdx-1] = rightBrother.Data[0]
		rightBrother.deleteData(0)
		if !b.isLeaf(rightBrother) {
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

	if node.Parent == b.root && len(b.root.Data) == 0 {
		b.root = node
		node.Parent = nil
		return
	}
	b.balance(node.Parent, key)
}

func (b *BTree) leftBrother(node *Node, key interface{}) (brother *Node, brotherIndex int) {
	if node.Parent != nil {
		brotherIndex, _ = b.Search(node.Parent, key)
		brotherIndex--
		if brotherIndex >= 0 && brotherIndex < len(node.Parent.Children) {
			return node.Parent.Children[brotherIndex], brotherIndex
		}
	}
	return nil, -1
}
func (b *BTree) rightBrother(node *Node, key interface{}) (brother *Node, brotherIndex int) {
	if node.Parent != nil {
		brotherIndex, _ = b.Search(node.Parent, key)
		brotherIndex--
		if brotherIndex < len(node.Parent.Children) {
			return node.Parent.Children[brotherIndex], brotherIndex
		}
	}
	return nil, -1
}

func (b *BTree) maxChildren() int { return b.deep }
func (b *BTree) minChildren() int { return (b.deep + 1) / 2 }
