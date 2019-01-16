package containers

import (
	"math"
)

// init size of queue, must be a number of power of 2.
const initQueueSize = 8

// Queue is a data structure queue implementation based on ring buffer,
// it can be replaced by any Type inside buffer, so it's a generic impl.
// However, it's not Thread-Safe.
type Queue struct {
	buffer           []interface{}
	head, tail, size int
}

func NewQueue() *Queue {
	return &Queue{buffer: make([]interface{}, initQueueSize)}
}

// Push push a new element into queue
func (q *Queue) Push(elem interface{}) {
	if q.size == len(q.buffer) {
		q.grow()
	}

	q.buffer[q.tail] = elem
	// bitwise modulus
	q.tail = (q.tail + 1) & (len(q.buffer) - 1)
	q.size++
}

// Peek return the first element in queue but never remove it.
func (q *Queue) Peek() interface{} {
	if q.size <= 0 {
		return nil
	}
	return q.buffer[q.head]
}

// PopHead remove the first element in queue.
func (q *Queue) Pop() (elem interface{}) {
	if q.size <= 0 {
		return nil
	}
	n := len(q.buffer)
	elem = q.buffer[q.head]
	q.size--
	// shrink for saving memory
	if 4*q.size < n && n > initQueueSize {
		// size of buffer must be power of 2 for bitwise modulus.
		newSize := shiftToPowOf2(q.size)
		if newSize < initQueueSize {
			newSize = initQueueSize
		}
		newBuff := make([]interface{}, newSize)
		// deprecate head element directly.
		if q.tail > q.head {
			copy(newBuff, q.buffer[q.head+1:q.tail])
		} else {
			n := copy(newBuff, q.buffer[q.head+1:])
			copy(newBuff[n:], q.buffer[:q.tail])
		}
		// reset points
		q.head = 0
		q.tail = q.size
		q.buffer = newBuff
		return
	}

	q.buffer[q.head] = nil
	// bitwise modulus
	q.head = (q.head + 1) & (len(q.buffer) - 1)
	return
}

// Size return current number of elements hold in queue.
func (q *Queue) Size() int {
	return q.size
}

// grow scale the queue buffer by doubling up to queue.size.
func (q *Queue) grow() {
	newBuff := make([]interface{}, q.size<<1)

	if q.tail > q.head {
		copy(newBuff, q.buffer)
	} else {
		n := copy(newBuff, q.buffer[q.head:])
		copy(newBuff[n:], q.buffer[:q.tail])
	}

	q.head = 0
	q.tail = q.size
	q.buffer = newBuff
}

func isPowOf2(num int) bool {
	return num > 0 && (num&(num-1)) == 0
}

func shiftToPowOf2(num int) int {
	if isPowOf2(num) {
		return num
	}
	ceil := math.Ceil(math.Log2(float64(num)))
	return int(math.Pow(2.0, ceil))
}
