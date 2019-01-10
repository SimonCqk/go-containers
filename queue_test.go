package pond

import "testing"

func TestQueueAdd(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 1000; i++ {
		q.Add(i)
	}
	for i := 0; i < 1000; i++ {
		ele, err := q.Peek()
		if err != nil {
			t.Error("error: ", err)
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
		ele, err = q.Pop()
		if err != nil {
			t.Error("error: ", err)
		}
		if ele = ele.(int); ele != i {
			t.Error(i, " is expected")
		}
	}
}
