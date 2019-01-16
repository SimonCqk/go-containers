package containers

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestHeapInt(t *testing.T) {
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

func TestHeapFloat(t *testing.T) {
	h := NewHeap()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]float64, 0, 1000)
	for i := 0; i < 1000; i++ {
		num := rnd.Float64()
		nums = append(nums, num)
		h.Push(num)
	}
	if h.Size() != 1000 {
		t.Error("size of heap expected as 1000, got ", h.Size())
	}
	sort.Float64s(nums)
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

func TestHeapStr(t *testing.T) {
	h := NewHeap()
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const rndStrLen = 32
	strRnd := func(n int) string {
		b := make([]byte, n)
		for i := range b {
			b[i] = chars[rand.Intn(len(chars))]
		}
		return string(b)
	}
	strSlice := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		str := strRnd(rndStrLen)
		strSlice = append(strSlice, str)
		h.Push(str)
	}
	if h.Size() != 1000 {
		t.Error("size of heap expected as 1000, got ", h.Size())
	}
	sort.Strings(strSlice)
	for i := 0; i < 1000; i++ {
		str, err := h.Pop()
		if err != nil {
			t.Error("got error when Pop, error: ", err)
		}
		if str != strSlice[i] {
			t.Error(strSlice[i], " is expected, got ", str)
		}
	}
}

type newInt struct {
	a, b int
}

type newIntSlice []newInt

func (ni newIntSlice) Len() int           { return len(ni) }
func (ni newIntSlice) Swap(i, j int)      { ni[i], ni[j] = ni[j], ni[i] }
func (ni newIntSlice) Less(i, j int) bool { return (ni[i].a + ni[i].b) < (ni[j].a + ni[j].b) }

func TestHeapCustomized(t *testing.T) {

	newIntLess := func(a, b interface{}) bool {
		ai := a.(newInt)
		bi := b.(newInt)
		return (ai.a + ai.b) < (bi.a + bi.b)
	}

	h := NewHeap(newIntLess)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]newInt, 0, 1000)
	for i := 0; i < 1000; i++ {
		ni := newInt{a: rnd.Intn(10000), b: rnd.Intn(10000)}
		nums = append(nums, ni)
		h.Push(ni)
	}
	if h.Size() != 1000 {
		t.Error("size of heap expected as 1000, got ", h.Size())
	}
	sort.Sort(newIntSlice(nums))
	for i := 0; i < 1000; i++ {
		num, err := h.Pop()
		if err != nil {
			t.Error("got error when Pop, error: ", err)
		}
		ni := num.(newInt)
		if (ni.a + ni.b) != (nums[i].a + nums[i].b) {
			t.Error(nums[i], " is expected, got ", num)
		}
	}
}
