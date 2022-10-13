// A* Search Algorithm
//1.  Initialize the open list
//2.  Initialize the closed list put the starting node on the open
//    list (you can leave its f at zero)
//3.  while the open list is not empty
//    a) find the node with the least f on the open list, call it "q"
//    b) pop q off the open list
//    c) generate q's 8 successors and set their parents to q
//    d) for each successor
//        i) if successor is the goal, stop search
//        ii) else, compute both g and h for successor,
//          successor.g = q.g + distance between successor and q
//          successor.h = distance from goal to successor (This can be done using many
//          ways, we will discuss three heuristics-Manhattan, Diagonal and Euclidean Heuristics)
//          successor.f = successor.g + successor.h
//        iii) if a node with the same position as successor is in the OPEN list which has a  lower f than successor,
//             skip this successor
//        iV) if a node with the same position as successor  is in the CLOSED list which has
//            a lower f than successor, skip this successor otherwise, add  the node to the open list
//     end (for loop)
//    e) push q on the closed list
//    end (while loop)

package pathfinding

import (
	"container/heap"
	"math"

	"k8s.io/utils/strings/slices"

	"github.com/guojia99/go-num/pathfinding/area"
)

type anode struct {
	*area.Node
	F, G, H int
}

type (
	anodes    []*anode
	anodeDict map[*area.Node]*anode
	nodeLink  map[*area.Node]*area.Node
)

func (a anodes) Len() int            { return len(a) }
func (a anodes) Less(i, j int) bool  { return a[i].F < a[j].F }
func (a anodes) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a *anodes) Push(x interface{}) { *a = append(*a, x.(*anode)) }
func (a *anodes) Pop() interface{} {
	old := *a
	length := len(old)
	x := old[length-1]
	*a = old[0 : length-1]
	return x
}

func (a *anode) CalcNG() int { return a.G + 1 }
func (a *anode) CalcNH(node *area.Node) int {
	return int(math.Abs(float64(a.X-node.X)) + math.Abs(float64(a.Y-node.Y)))
}

func AStat(start, end *area.Node, roadTyp string, obstacleList ...string) (road []*area.Node, newArea *area.Area) {
	var (
		openDict  = make(anodeDict)
		closeDict = make(anodeDict)
		linkList  = make(nodeLink)
		ads       = make(anodes, 0)
	)
	startANode, endANode := &anode{Node: start}, &anode{Node: end}
	openDict[start] = startANode
	heap.Init(&ads)
	heap.Push(&ads, startANode)

	back := false

	for {
		cNode := heap.Pop(&ads).(*anode)
		delete(openDict, cNode.Node)
		closeDict[cNode.Node] = cNode

		for _, val := range cNode.Around() {
			if _, ok := closeDict[val]; ok {
				continue
			}
			if slices.Contains(obstacleList, val.Typ) {
				continue
			}
			if _, ok := openDict[val]; ok {
				parentNode := linkList[val]
				parentANode := closeDict[parentNode]
				if parentANode.CalcNG() < parentANode.CalcNH(cNode.Node) {
					cNode = parentANode
					heap.Push(&ads, cNode)
					g, h := cNode.CalcNG(), endANode.CalcNH(val)
					openDict[val] = &anode{Node: val, G: g, H: h, F: g + h}
					break
				}
			}

			g, h := cNode.CalcNG(), endANode.CalcNH(val)
			node := &anode{Node: val, G: g, H: h, F: g + h}
			heap.Push(&ads, node)
			linkList[val] = cNode.Node
			if val.ID == end.ID {
				back = true
				break
			}
			openDict[val] = node
		}
		if len(openDict) == 0 || back {
			break
		}
	}
	if back {
		road = append(road, end)
		newArea = area.CreateArea(start.Area.X, start.Area.Y)

		for i := 0; i < start.Area.X; i++ {
			for j := 0; j < start.Area.Y; j++ {
				n, _ := start.Area.GetNode(i, j)
				newArea.SetNode(i, j, n.Typ)
			}
		}

		for {
			f, ok := linkList[end]
			f.Typ = roadTyp
			road = append(road, f)
			newArea.SetNode(f.X, f.Y, roadTyp)
			if f.ID == start.ID || !ok {
				break
			}
			end = f
		}
	}
	return
}
