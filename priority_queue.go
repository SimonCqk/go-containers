package containers

import "reflect"

// Element indicate the elements contained in priority queue,
// including a real value and its priority.
type Element struct {
	Value    interface{}
	Priority int
}

// LessCompFunc for comparing two Element instance.
func elementLessComp(a, b interface{}) bool {
	e1 := a.(*Element)
	e2 := b.(*Element)
	if e1.Priority == e2.Priority {
		if _, ok := lessComparator[reflect.ValueOf(e1.Value).Kind()]; ok {
			return lessComparator[reflect.ValueOf(e1.Value).Kind()](e1.Value, e2.Value)
		}
	}
	return e1.Priority < e2.Priority
}

// NewPriorityQueue make a new priority queue instance.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		minHeap: NewHeap(elementLessComp),
	}
}

// PriorityQueue implementation, meanwhile it's a max heap.
type PriorityQueue struct {
	minHeap *Heap
}

// Push add a new element with its priority into priority queue,
// since queue holds a minHeap, let the internal priority be the
// negative priority of the origin, so min-max relationship is
// reversed.
func (q *PriorityQueue) Push(elem interface{}, priority int) {
	q.minHeap.Push(&Element{Value: elem, Priority: -priority})
}

// Peek return the element with max priority.
func (q *PriorityQueue) Peek() *Element {
	if q.Size() == 0 {
		return nil
	}
	top := q.minHeap.Top().(*Element)
	// recover its real priority.
	return &Element{Value: top.Value, Priority: -top.Priority}
}

// Pop remove the element with max priority and return it.
func (q *PriorityQueue) Pop() *Element {
	if q.Size() == 0 {
		return nil
	}
	top := q.minHeap.Pop().(*Element)
	top.Priority = -top.Priority
	return top
}

func (q *PriorityQueue) Size() int {
	return q.minHeap.Size()
}
