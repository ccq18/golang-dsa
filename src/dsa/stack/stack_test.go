package stack

import (
	"testing"
	"fmt"
	"reflect"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push("aasas")
	//4321
	for !s.IsEmpty() {
		i:=s.Pop()
		fmt.Println(reflect.TypeOf(i))
		//fmt.Println(strconv.FormatInt(i))
	}
}