package containers

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	h := NewHeap()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0, 10)
	for i := 0; i < 1000; i++ {
		num := rnd.Int()
		nums = append(nums, num)
		h.Push(num)
	}
	if h.Size() != 1000 {
		t.Error("size of heap expected as 1000, got ", h.Size())
	}
	sort.Ints(nums)
	for i := 0; i < 1000; i++ {
		num, err := h.Pop()
		if err != nil {
			t.Error("got error when Pop, error: ", err)
		}
		if num != nums[i] {
			t.Error(nums[i], " is expected, got ", num)
		}
	}
}
