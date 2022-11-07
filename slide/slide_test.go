package slide

import (
	"fmt"
	"testing"
)

func Test_Intersection(t *testing.T) {
	a := []interface{}{1, 2, "3", 3, "4", 5}
	b := []interface{}{1, 999, 888, 777}
	c := []interface{}{11111, 1111111}

	fmt.Println(Intersection(a, b))
	fmt.Println(Intersection(b, c))
}

func BenchmarkIntersectionInt(b *testing.B) {
	b.StopTimer()
	newSlide := func(start int) []int {
		var out []int
		for i := start; i < start+100; i++ {
			out = append(out, i)
		}
		return out
	}

	aa := newSlide(1000)
	bb := append(newSlide(22222))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IntersectionInt(aa, bb)
	}
}
