package area

import "fmt"

type NodeType string

const (
	Flat     = "□"
	Obstacle = "#"
	Road     = "+"
)

const nodeIdFormat = "%d-%d"

func getNodeID(x int, y int) string { return fmt.Sprintf(nodeIdFormat, x, y) }

type Node struct {
	Area *Area
	ID   string
	X, Y int
	Typ  string
}

func (a *Area) NewNode(x, y int, typ string) *Node {
	return &Node{
		ID:   getNodeID(x, y),
		Area: a,
		X:    x,
		Y:    y,
		Typ:  typ,
	}
}

func (n *Node) Around() []*Node {
	x, y := n.X, n.Y
	als := [4][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}}

	var out = make([]*Node, 0)
	for _, val := range als {
		node, ok := n.Area.Nodes[getNodeID(val[0], val[1])]
		if ok {
			out = append(out, node)
		}
	}
	return out
}
