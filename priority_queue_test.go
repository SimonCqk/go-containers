package containers

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type elements []*Element

func (e elements) Len() int           { return len(e) }
func (e elements) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e elements) Less(i, j int) bool { return e[i].Priority < e[j].Priority }

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	elements := elements(make([]*Element, 0, 10))
	for i := 0; i < 10; i++ {
		elem := &Element{Value: i, Priority: rnd.Intn(10000)}
		elements = append(elements, elem)
		pq.Push(elem.Value, elem.Priority)
	}
	sort.Sort(elements)
	for i := 9; i >= 0; i-- {
		elem := pq.Peek()
		if elem == nil {
			t.Error("empty priority queue when Peek.")
		}
		if elem.Priority != elements[i].Priority ||
			elem.Value != elements[i].Value {
			t.Errorf("%v is expected, got %v", elements[i], *elem)
		}
		elem = pq.Pop()
		if elem == nil {
			t.Error("empty priority queue when Peek.")
		}
		if elem.Priority != elements[i].Priority ||
			elem.Value != elements[i].Value {
			t.Errorf("%v is expected, got %v", elements[i], *elem)
		}
	}
}
