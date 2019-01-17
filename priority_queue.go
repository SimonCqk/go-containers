package containers

// Element indicate the elements contained in priority queue,
// including a real value and its priority.
type Element struct {
	Value    interface{}
	Priority int
}

func elementLessComp(a, b interface{}) bool {
	return a.(*Element).Priority < b.(*Element).Priority
}

// NewPriorityQueue make a new priority queue instance.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		minHeap: NewHeap(elementLessComp),
	}
}

type PriorityQueue struct {
	minHeap *Heap
}

func (q *PriorityQueue) Push(elem interface{}, priority int) {
	q.minHeap.Push(&Element{Value: elem, Priority: -priority})
}

func (q *PriorityQueue) Peek() *Element {
	if q.Size() == 0 {
		return nil
	}
	top := q.minHeap.Top().(*Element)
	return &Element{Value: top.Value, Priority: -top.Priority}
}

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
