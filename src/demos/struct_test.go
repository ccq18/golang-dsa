package demos

import (
	"fmt"
	"testing"
)

func TestC1(t *testing.T) {
	c := new(C1)
	fmt.Println(c.a)
	c2 := new(C2)
	c2.int=111
	fmt.Println(c2.int)
	c2.a=111

	fmt.Println(c2.C1.a)


}

