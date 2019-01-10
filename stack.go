package containers

import (
	"errors"
)

const initStackSize = 8

type Stack struct {
	buffer            []interface{}
	top, bottom, size int
}

func NewStack() *Stack {
	return &Stack{buffer: make([]interface{}, initStackSize)}
}

func (s *Stack) Push(elem interface{}) {
	if s.size == len(s.buffer) {
		s.grow()
	}
	s.buffer[s.top] = elem
	s.top++
	s.size++
}

// Peek return the first element in stack but never remove it.
func (s *Stack) Peek() (elem interface{}, err error) {
	if s.size <= 0 {
		return nil, errors.New("call Peek() on a empty queue")
	}
	return s.buffer[s.top-1], nil
}

// PopHead remove the first element in stack.
func (s *Stack) Pop() (elem interface{}, err error) {
	if s.size <= 0 {
		return nil, errors.New("call Pop() on a empty queue")
	}
	elem = s.buffer[s.top-1]
	s.buffer[s.top-1] = nil
	s.top--
	s.size--
	return
}

// Size return current number of elements hold in stack.
func (s *Stack) Size() int {
	return s.size
}

// grow scale the stack buffer by doubling up to queue.size.
func (s *Stack) grow() {
	newBuff := make([]interface{}, s.size<<1)
	copy(newBuff, s.buffer)
	s.buffer = newBuff
}
