package link_list

import (
	"testing"
	"fmt"
)

func TestLinkList_Add(t *testing.T) {
	list:= NewLinkList()
	for i:=0;i<1000;i++{
		list.Add(i)

	}
	fmt.Println(list.Get(101))
	fmt.Println(list.Get(0))
	fmt.Println(list.Get(999))
	//
	list.Print()
	//list.DePrint()
}
func TestLinkList_Insert(t *testing.T) {
	list:= NewLinkList()
	list.Insert(0,1)
	list.Insert(1,2)
	list.Insert(2,123)

	list.Print()
}
func TestLinkList_AddLast(t *testing.T) {
	list:= NewLinkList()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(123)

	list.Print()
}
func TestLinkList_Del(t *testing.T) {
	list:= NewLinkList()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	//list.Print()
	//
	list.Print()

	fmt.Println("---")

	list.Del(0)
	list.Print()

	//
	fmt.Println("---")
	//
	list.DelLast()
	list.Print()
	//fmt.Println(list.Len())


}