package stack

import (
	. "dsa/list"
	"dsa/link_list"
)

type Stack struct {
	list *link_list.LinkList
}

func NewStack() (*Stack) {

	s := new(Stack)
	s.list = link_list.NewLinkList()

	return s
}
func (s *Stack) Push(i Type) int {
	return s.list.Add(i)
}
func (s *Stack) Pop() Type {
	return s.list.Del(s.list.Len() - 1)
}
func (s *Stack) GetTop() Type {
	return  s.list.GetLast()
}
func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}
func (s *Stack) Len() int {
	return s.list.Len()
}