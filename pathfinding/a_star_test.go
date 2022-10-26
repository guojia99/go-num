package pathfinding

import (
	"fmt"
	"testing"

	"github.com/guojia99/go-num/pathfinding/area"
)

func TestAStat(t *testing.T) {
	f := area.Flat
	e := " "
	b := "#"
	road := "o"

	cells := [][]string{
		{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
		{f, e, e, e, e, e, e, b, e, e, e, e, e, b, e, f},
		{f, e, e, e, e, e, e, b, e, e, e, e, e, b, e, f},
		{f, e, e, e, e, e, e, b, e, e, e, e, e, b, e, f},
		{f, e, e, e, e, e, e, b, e, e, e, e, e, e, e, f},
		{f, e, e, e, e, e, e, b, e, e, e, e, e, b, e, f},
		{f, e, e, e, e, e, e, e, e, e, e, e, e, b, e, f},
		{f, b, b, b, b, b, b, b, e, e, e, e, e, b, e, f},
		{f, e, e, e, e, e, e, b, e, e, e, e, e, b, e, f},
		{f, e, e, e, e, e, e, b, b, b, b, b, b, b, e, f},
		{f, e, e, e, e, e, e, e, e, e, e, e, e, e, e, f},
		{f, e, e, e, e, e, e, e, e, e, e, e, e, e, e, f},
		{f, e, e, e, e, e, e, e, e, e, e, e, e, e, e, f},
		{f, e, e, e, e, e, e, e, e, e, e, e, e, e, e, f},
		{f, e, e, e, e, e, e, e, e, e, e, e, e, e, e, f},
		{f, f, f, f, f, f, f, f, f, f, f, f, f, f, f, f},
	}

	a, _ := area.With2DSlideCreateArea(cells)
	start, _ := a.GetNode(1, 1)
	end, _ := a.GetNode(3, 8)
	_, newA := AStat(start, end, road, b, f)
	fmt.Println(newA)
}
