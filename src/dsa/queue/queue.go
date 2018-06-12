package queue

import (
	. "dsa/list"
	"dsa/link_list"
)

type Queue struct {
	list *link_list.LinkList
}
func NewQueue() (*Queue) {
	q := new(Queue)
	q.list = link_list.NewLinkList()

	return q
}

func (this *Queue)Fist() Type {
	return this.list.Get(0)
}
func (this *Queue)Last()  Type{
	return this.list.GetLast()

}
func (this *Queue)EnQueue(v Type) int {
	return this.list.Add(v)

}
func  (this *Queue)DeQueue()Type  {
	return this.list.Del(0)
}
func (this *Queue) IsEmpty() bool {
	return this.list.IsEmpty()
}
func (this *Queue) Len() int {
	return this.list.Len()
}



