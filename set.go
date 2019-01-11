package containers

import "github.com/pkg/errors"

type Set struct {
	core map[interface{}]bool
}

func NewSet() *Set {
	return &Set{core: make(map[interface{}]bool)}
}

func (s *Set) Add(elem interface{}) {
	if _, ok := s.core[elem]; !ok {
		s.core[elem] = true
	}
}

func (s *Set) All() []interface{} {
	all := make([]interface{}, 0, len(s.core))
	for each := range s.core {
		all = append(all, each)
	}
	return all
}

func (s *Set) Remove(elem interface{}) error {
	if _, ok := s.core[elem]; !ok {
		return errors.New("no such element in Set")
	}
	delete(s.core, elem)
	return nil
}

func (s *Set) Size() int {
	return len(s.core)
}
