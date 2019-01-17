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
		ele := s.Peek()
		if ele == nil {
			t.Error("empty stack when Peek.")
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
		ele = s.Pop()
		if ele == nil {
			t.Error("empty stack when Pop. ")
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
	}
}
