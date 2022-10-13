package area

import (
	"fmt"
	"testing"
)

func TestCreateArea(t *testing.T) {
	a := CreateArea(10, 11)
	fmt.Println(a.String())
}

func TestWith2DSlideCreateArea(t *testing.T) {
	cells := [][]string{
		{"#", "#", "#", "#", "-"},
		{"#", "#", "#", "-", "-"},
		{"#", "#", "-", "-"},
		{"#", "-", "#", "#", "-"},
		{"-", "#", "#"},
		{"#", "-", "#", "#", "-"},
	}
	a, err := With2DSlideCreateArea(cells)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a.String())
}
