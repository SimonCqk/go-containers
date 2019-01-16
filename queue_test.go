package containers

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 1000; i++ {
		q.Push(i)
	}
	for i := 0; i < 1000; i++ {
		ele := q.Peek()
		if ele == nil {
			t.Error("got nil when Peek.")
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
		ele = q.Pop()
		if ele == nil {
			t.Error("got nil when Pop.")
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
	}
}
