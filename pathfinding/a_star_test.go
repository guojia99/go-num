package pathfinding

import (
	"fmt"
	"testing"

	"github.com/guojia99/go-num/pathfinding/area"
)

func TestAStat(t *testing.T) {
	f := area.Flat
	r := area.Road
	b := "#"

	cells := [][]string{
		{f, b, f, f, f, f, f, f, f, f},
		{f, f, f, b, b, b, b, b, f, f},
		{f, b, f, f, f, f, f, b, f, f},
		{f, f, f, f, f, f, f, f, f, f},
		{b, b, b, b, b, b, b, f, f, f},
		{f, f, f, f, f, f, b, f, f, f},
		{f, f, f, f, f, f, b, f, f, f},
		{f, f, b, b, f, f, b, b, f, b},
		{f, b, f, f, b, f, f, f, f, b},
		{f, f, f, f, f, b, f, b, f, b},
		{f, f, f, f, f, f, b, f, b, b},
		{f, f, f, f, f, f, f, f, f, b},
	}

	a, _ := area.With2DSlideCreateArea(cells)
	start, _ := a.GetNode(0, 0)
	end, _ := a.GetNode(0, 11)
	_, newA := AStat(start, end, r, b)
	fmt.Println(newA)
}
