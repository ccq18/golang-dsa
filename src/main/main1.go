package main

import (
	//"dsa/list"
	//"fmt"
	"fmt"
	"time"
	"strconv"
)

func main() {
	//l := list.NewList()
	//fmt.Println(time.Now())
	//for i := 0; i < 100000000; i++ {
	//	l.Add(i)
	//	//l.Add(2)
	//	//l.Add(3)
	//}
	////fmtt.Pr
	//fmt.Println(l.Len())

	//python 36s
	//php 11s
	//golang 2s

	fmt.Println(time.Now())
	//var a []int
	a := make(map[string]int)

	f := func(i int)int {
		return i
	}
	// append works on nil slices.
	for i := 0; i < 10000000; i++ {
		a["i"+strconv.Itoa(i)] = f(i)
		//a = append(a, f(i))
	}
	fmt.Println(len(a))

	fmt.Println(time.Now())

	//fmt.Println(l.Get(0))
	//l.Clear()
	//fmt.Println(l.Get(0))

	//l.Print()
	//println(l.Get(0))
	//println(l.Get(1))
	//println(l.Get(2))
}
