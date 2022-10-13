package area

import (
	"errors"
	"fmt"
)

type Area struct {
	X, Y  int
	Nodes map[string]*Node
}

func CreateArea(x, y int) *Area {
	a := &Area{
		X: x, Y: y,
		Nodes: make(map[string]*Node),
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			a.Nodes[getNodeID(i, j)] = a.NewNode(i, j, Flat)
		}
	}
	return a
}

func With2DSlideCreateArea(in [][]string) (*Area, error) {
	if len(in) == 0 {
		return nil, errors.New("can't not use empty area row")
	}

	if len(in[0]) == 0 {
		return nil, errors.New("can't not use empty area col")
	}
	var a = CreateArea(len(in[0]), len(in))

	for i := 0; i < a.Y; i++ {
		for j := 0; j < a.X; j++ {
			if len(in) <= i {
				continue
			}
			if len(in[i]) <= j {
				continue
			}
			a.SetNode(j, i, in[i][j])
		}
	}
	return a, nil
}

func (a *Area) GetNode(x, y int) (*Node, error) {
	id := getNodeID(x, y)
	if _, ok := a.Nodes[id]; !ok {
		return nil, fmt.Errorf("not find grid by %s", id)
	}
	return a.Nodes[id], nil
}

func (a *Area) SetNode(x, y int, typ string) bool {
	id := getNodeID(x, y)
	if _, ok := a.Nodes[id]; !ok {
		return false
	}
	a.Nodes[id].Typ = typ
	return true
}

func (a *Area) String() string {
	var out string
	for i := 0; i < a.Y; i++ {
		for j := 0; j < a.X; j++ {
			node := a.Nodes[getNodeID(j, i)]
			out += string(node.Typ) + " "
		}
		out += "\n"
	}
	return out
}
