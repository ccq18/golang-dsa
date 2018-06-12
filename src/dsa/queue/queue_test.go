package queue

import (
	"testing"
	"fmt"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Println(q.Len())

	//123
	for !q.IsEmpty() {
		fmt.Println(q.DeQueue())
	}
	fmt.Println(q.Len())
}