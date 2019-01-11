package containers

import (
	"sort"
	"testing"
)

func TestSetNormal(t *testing.T) {
	s := NewSet()
	for i := 0; i < 1000; i++ {
		// add i to Set twice
		for j := 0; j < 2; j++ {
			s.Add(i)
		}
	}
	if s.Size() != 1000 {
		t.Error("size of Set is expected ", 1000)
	}
	all := s.All()
	intAll := make([]int, 0, len(all))
	for idx := range all {
		ele := all[idx].(int)
		intAll = append(intAll, ele)
	}
	sort.Ints(intAll)
	for idx, i := range intAll {
		if intAll[idx] != i {
			t.Error(idx, " is expected, got ", i)
		}
	}
}
