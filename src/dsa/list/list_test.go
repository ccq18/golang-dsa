package list

import (
	"testing"
	"fmt"
)

func TestList_Add(t *testing.T) {
	l:= NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	println(l.Get(0))
	println(l.Get(1))
	println(l.Get(2))

}
func TestList_Clear(t *testing.T) {
	l:= NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Clear()

}
func TestList_Set(t *testing.T) {
	l:= NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Set(1,11)
}
func TestList_Map(t *testing.T) {
	l:= NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Map(func(i Type) {
		fmt.Println(i)
	})
}

func TestList_Del(t *testing.T) {
	l:= NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Del(1)
	//1 3
	l.Map(func(i Type) {
		fmt.Println(i)
	})
}

func TestList_Insert(t *testing.T) {
	l:= NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Insert(1,4)
	// 1 4 2 3
	l.Map(func(i Type) {
		fmt.Println(i)
	})
}