package link_list

import (
	. "dsa/list"
	"fmt"
)

type LinkList struct {
	head *LinkNode
	last *LinkNode
	len  int
}
type LinkNode struct {
	value Type
	next  *LinkNode
	pre   *LinkNode
}

func NewLinkList() (*LinkList) {
	l := new(LinkList)
	l.len = 0
	return l
}
func (l *LinkList) Add(v Type) int {
	node := new(LinkNode)
	node.value = v
	if l.last == nil {
		l.head = node
		l.last = node
	} else {
		l.last.next = node
		node.pre = l.last
		l.last = node
	}
	return l.len
}

func (l *LinkList) Get(k int) (Type) {
	return l.getNode(k).value
}
func (l *LinkList) GetLast() (Type) {
	if l.len <= 0 {
		return nil
	}
	return l.last.value
}
func (l *LinkList) getNode(k int) (node *LinkNode) {

	if k > l.len/2 {
		node := l.last
		if node == nil {
			return nil
		}
		for i := l.len - 1; node != nil; node, i = node.pre, i-1 {
			if i == k {
				return node
			}
		}
	} else {
		node = l.head
		if node == nil {
			return nil
		}
		for i := 0; node != nil; node, i = node.next, i+1 {
			if i == k {
				return node
			}
		}
	}

	return nil
}
func (l *LinkList) AddLast(v Type) int {
	return l.Insert(l.len, v)
}
func (l *LinkList) AddFirst() int {
	return l.Insert(l.len, 0)
}
func (l *LinkList) DelLast() Type {
	if l.len <= 0 {
		return nil
	}
	return l.Del(l.len - 1)
}
func (l *LinkList) DelFirst() Type {
	if l.len <= 0 {
		return nil
	}
	return l.Del(0)
}
func (l *LinkList) Del(k int) Type {
	node := l.getNode(k)
	if node == nil {
		panic("节点不存在")
		//d(node)
	}
	//只有一个节点
	if node.pre == nil && node.next == nil {
		l.head = nil
		//是头节点
	} else if node.pre == nil {
		//
		node.next.pre = nil
		l.head = node.next

	} else if node.next == nil {
		node.pre.next = nil
		l.last = node.pre
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	l.len--
	return node.value
	//node
}
func (l *LinkList) Insert(k int, v Type) int {
	node := new(LinkNode)
	node.value = v

	if k == 0 {
		//头尾节点有一个为空则都为空
		if l.head == nil {
			l.last = node
		} else {
			node.next = l.head
			l.head.pre = node
		}
		l.head = node
	} else if k == l.len {
		//加入尾节点
		if l.last == nil {
			l.head = node
		} else {
			node.pre = l.last
			l.last.next = node
		}
		l.last = node
	} else {
		now := l.getNode(k)
		if now == nil {
			panic("当前节点不存在")
		}
		if now != nil {
			node.next = now
			node.pre = now.pre
			now.pre = node

		}
	}
	l.len++
	return l.len

}

//func (l *LinkList) Set(k int, v Type) {
//
//	l.arr[k] = v
//}

//func (l *LinkList) GetLast() (Type) {
//	if l.len <= 0 {
//		panic("当前长度为0")
//	}
//	return l.arr[l.len-1]
//}
//func (l *LinkList) Clear() {
//	l.len = 0
//}
//
func (l *LinkList) Map(f func(Type)) {
	node := l.head
	if node == nil {
		return
	}
	for ; node != nil; node = node.next {
		f(node.value)
	}
}
func (l *LinkList) DeMap(f func(Type)) {
	node := l.last
	if node == nil {
		return
	}
	for ; node != nil; node = node.pre {
		f(node.value)
	}

}
func (l *LinkList) Len() int {
	return l.len
}

func (l *LinkList) Print() {
	l.Map(func(p Type) {
		fmt.Println(p)
	})

}
func (l *LinkList) DePrint() {
	l.DeMap(func(p Type) {
		fmt.Println(p)
	})

}
func (l *LinkList) IsEmpty() bool {
	return l.len <= 0
}
