package containers

import (
	"container/heap"
	"github.com/pkg/errors"
)

const initHeapSize = 8

// Less is a interface for comparing two values stored in
// Heap. If this < l, then return true, else return false.
type Less interface {
	Less(l Less) bool
}

// Minimum heap implementation.
type Heap struct {
	core *heapCore
}

func NewHeap() *Heap {
	return &Heap{core: &heapCore{
		buffer: make([]Less, 0, initHeapSize),
	}}
}

func (h *Heap) Push(elem Less) {
	heap.Push(h.core, elem)
}

func (h *Heap) Pop() (interface{}, error) {
	if h.core.Len() == 0 {
		return nil, errors.New("pop on a empty heap")
	}
	return heap.Pop(h.core), nil
}

func (h *Heap) Size() int {
	return h.core.Len()
}

// basic implementation of container.heap.Interface
type heapCore struct {
	heap.Interface
	buffer []Less
}

func (h *heapCore) grow() {
	newBuff := make([]Less, len(h.buffer)<<1)
	copy(newBuff, newBuff)
	h.buffer = newBuff
}

func (h *heapCore) Len() int {
	return len(h.buffer)
}

func (h *heapCore) Swap(i, j int) {
	h.buffer[i], h.buffer[j] = h.buffer[j], h.buffer[i]
}

func (h *heapCore) Less(i, j int) bool {
	return h.buffer[i].Less(h.buffer[j])
}

func (h *heapCore) Push(elem Less) {
	h.buffer = append(h.buffer, elem)
}

func (h *heapCore) Pop() interface{} {
	min := h.buffer[0]
	h.buffer = h.buffer[1:]
	return min
}
