/*
Algorithm	Average	Worst case
Space	O(n)	O(n)
search	O(log n)	O(log n)
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

// organizeNodes 理枝
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

// searchForm 依据不同的node和key寻找到位置和相应数据
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
		fmt.Println(brotherIndex)
		brotherIndex--
		if brotherIndex < len(node.Parent.Children) {
			return node.Parent.Children[brotherIndex], brotherIndex
		}
	}
	return nil, -1
}

func (x *BTree) maxChildren() int { return int(x.Deep) }
func (x *BTree) minChildren() int { return int((x.Deep + 1) / 2) }
