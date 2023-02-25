package _33

import (
	"fmt"
	"testing"
)

func TestRecentCounter_Ping(t *testing.T) {
	c := Constructor()
	fmt.Println(c.Ping(1))
	fmt.Println(c.Ping(100))
	fmt.Println(c.Ping(3001))
	fmt.Println(c.Ping(3002))
}
