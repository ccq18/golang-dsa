package list

import "fmt"

type List struct {
	arr []Type
	len int
	mem int
}
type Type = interface{}

func NewList() (*List) {
	l := new(List)
	fmt.Println()
	l.mem = 2
	l.len = 0
	l.arr = make([]Type, l.mem)
	return l
}
func (l *List) Add(v Type) int {
	l.prepare(l.len + 1)
	l.arr[l.len] = v
	l.len++
	return l.len
}
func (l *List) prepare(newlen int) {
	if newlen > l.mem {
		l.mem *= 2
		arr := make([]Type, l.mem)
		copy(arr, l.arr)
		l.arr = arr
	}
}
func (l *List) Set(k int, v Type) {

	l.arr[k] = v
}
func (l *List) Get(k int) (Type) {
	if l.len <= k {
		panic("k:" + string(k) + "不存在")
	}
	return l.arr[k]
}
func (l *List) GetLast() (Type) {
	if l.len <= 0 {
		panic("当前长度为0")
	}
	return l.arr[l.len-1]
}
func (l *List) Clear() {
	l.len = 0
}

func (l *List) Map(f func(Type)) {
	for i := 0; i < l.len; i++ {
		f(l.arr[i])
	}
}

func (l *List) Del(k int) Type {

	if l.len <= k || k < 0 {
		panic("k:" + string(k) + "不存在")
	}
	v := l.arr[k]
	l.len -= 1
	for i := k; i < l.len; i++ {
		l.arr[i] = l.arr[i+1]
	}
	return v
}
func (l *List) Insert(k int, v Type) {
	l.prepare(l.len + 1)
	//将k 位置的所有值向后移动一位
	if (k < l.len) {
		for i := l.len - 1; i >= k; i-- {
			l.arr[i+1] = l.arr[i]
		}
	}

	l.len += 1
	l.arr[k] = v
}

func (l *List) Len() int {
	return l.len
}
func (l *List) Mem() int {
	return l.mem
}
func (l *List) Print() {
	l.Map(func(i Type) {
		fmt.Println(i)
	})
}
func (l *List) IsEmpty() bool {
	return l.len <= 0
}
