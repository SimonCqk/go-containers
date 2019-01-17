package containers

import (
	"fmt"
)

type Set struct {
	core map[interface{}]bool
}

func NewSet() *Set {
	return &Set{core: make(map[interface{}]bool)}
}

// Add add a new element into set.
func (s *Set) Add(elem interface{}) {
	if _, ok := s.core[elem]; !ok {
		s.core[elem] = true
	}
}

// All return all elements in set.
func (s *Set) All() []interface{} {
	all := make([]interface{}, 0, len(s.core))
	for each := range s.core {
		all = append(all, each)
	}
	return all
}

// Remove remove the specified element in set.
func (s *Set) Remove(elem interface{}) error {
	if _, ok := s.core[elem]; !ok {
		return fmt.Errorf("no such element in Set")
	}
	delete(s.core, elem)
	return nil
}

// Size return current number of elements in set.
func (s *Set) Size() int {
	return len(s.core)
}
