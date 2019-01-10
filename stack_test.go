package containers

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i <= 1000; i++ {
		s.Push(i)
	}
	for i := 1000; i >= 0; i-- {
		ele, err := s.Peek()
		if err != nil {
			t.Error("error: ", err)
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
		ele, err = s.Pop()
		if err != nil {
			t.Error("error: ", err)
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
	}
}
