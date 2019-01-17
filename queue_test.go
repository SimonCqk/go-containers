package containers

import (
	"container/list"
	"testing"
)

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

type listQueue struct {
	l *list.List
}

func (q *listQueue) PushBack(v interface{}) {
	if v == nil {
		return
	}
	q.l.PushBack(v)
}

func (q *listQueue) Front() *list.Element {
	return q.l.Front()
}

func (q *listQueue) Remove(e *list.Element) {
	if e == nil {
		return
	}
	q.l.Remove(e)
}

func (q *listQueue) Pop() *list.Element {
	elem := q.Front()
	q.Remove(elem)
	return elem
}

func BenchmarkQueuePush(t *testing.B) {
	q := NewQueue()
	for i := 0; i < t.N; i++ {
		q.Push(i)
	}
}

func BenchmarkListQueuePush(t *testing.B) {
	lq := listQueue{l: list.New()}
	for i := 0; i < t.N; i++ {
		lq.PushBack(i)
	}
}

func BenchmarkQueuePop(t *testing.B) {
	q := NewQueue()
	for i := 0; i < t.N; i++ {
		q.Push(i)
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		q.Pop()
	}
}

func BenchmarkListQueuePop(t *testing.B) {
	lq := listQueue{l: list.New()}
	for i := 0; i < t.N; i++ {
		lq.PushBack(i)
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		lq.Pop()
	}
}
