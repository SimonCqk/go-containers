package containers

import (
	"container/heap"
	"github.com/pkg/errors"
	"reflect"
)

const initHeapSize = 8

var lessComparator map[reflect.Kind]LessCompFunc

// LessCompFunc provide a entry for compare two instance by customized rule.
type LessCompFunc func(a, b interface{}) bool

// Minimum heap implementation.
type Heap struct {
	core *heapCore
}

// NewHeap make a heap, if contained type are not built-in types, you
// should specify a LessCompFunc to compare two interface{}.
func NewHeap(lessCompFunc ...LessCompFunc) *Heap {
	return &Heap{core: &heapCore{
		buffer: make([]interface{}, 0, initHeapSize),
		// default nil, lazy bind when first element pushed into heap.
		lessComparator: append(lessCompFunc, nil)[0],
	}}
}

// Push append a new element to heap.
func (h *Heap) Push(elem interface{}) {
	// lessComparator is not specified, take it as a built-in type,
	// or panic.
	if h.core.lessComparator == nil {
		if _, ok := lessComparator[reflect.ValueOf(elem).Kind()]; ok {
			h.core.lessComparator = lessComparator[reflect.ValueOf(elem).Kind()]
		} else {
			panic("unsupported type, or specify a LessComparator when make a new Heap")
		}
	}
	heap.Push(h.core, elem)
}

// Pop remove and return the min value element in heap, if heap is empty,
// return a error.
func (h *Heap) Pop() (interface{}, error) {
	if h.core.Len() == 0 {
		return nil, errors.New("pop on a empty heap")
	}
	return heap.Pop(h.core), nil
}

// Top return the top(min) element in heap.
func (h *Heap) Top() interface{} {
	if h.Size() == 0 {
		return nil
	}
	return h.core.buffer[h.core.Len()-1]
}

func (h *Heap) Size() int {
	return h.core.Len()
}

// basic implementation of container.heap.Interface
type heapCore struct {
	buffer         []interface{}
	lessComparator LessCompFunc
}

func (h *heapCore) grow() {
	newBuff := make([]interface{}, len(h.buffer)<<1)
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
	return h.lessComparator(h.buffer[i], h.buffer[j])
}

func (h *heapCore) Push(elem interface{}) {
	h.buffer = append(h.buffer, elem)
}

func (h *heapCore) Pop() interface{} {
	n := len(h.buffer) - 1
	min := h.buffer[n]
	h.buffer = h.buffer[:n]
	return min
}

func init() {
	// dirty basic type less comparator
	lessComparator = map[reflect.Kind]LessCompFunc{
		reflect.Uint:    func(a, b interface{}) bool { return a.(uint) < b.(uint) },
		reflect.Uint8:   func(a, b interface{}) bool { return a.(uint8) < b.(uint8) },
		reflect.Uint16:  func(a, b interface{}) bool { return a.(uint16) < b.(uint16) },
		reflect.Uint32:  func(a, b interface{}) bool { return a.(uint32) < b.(uint32) },
		reflect.Uint64:  func(a, b interface{}) bool { return a.(uint64) < b.(uint64) },
		reflect.Int:     func(a, b interface{}) bool { return a.(int) < b.(int) },
		reflect.Int8:    func(a, b interface{}) bool { return a.(int8) < b.(int8) },
		reflect.Int16:   func(a, b interface{}) bool { return a.(int16) < b.(int16) },
		reflect.Int32:   func(a, b interface{}) bool { return a.(int32) < b.(int32) },
		reflect.Int64:   func(a, b interface{}) bool { return a.(int64) < b.(int64) },
		reflect.Float32: func(a, b interface{}) bool { return a.(float32) < b.(float32) },
		reflect.Float64: func(a, b interface{}) bool { return a.(float64) < b.(float64) },
		reflect.String:  func(a, b interface{}) bool { return a.(string) < b.(string) },
	}
}
