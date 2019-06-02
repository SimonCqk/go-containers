package containers

import "errors"

type DisjointSet struct {
	sets map[interface{}]*subset
}

type subset struct {
	rank   int
	parent interface{}
}

// NewDisjointSet create a new disjoint set with init values.
func NewDisjointSet(vals []interface{}) *DisjointSet {
	s := DisjointSet{}
	for _, val := range vals {
		s.sets[val] = &subset{parent: nil}
	}
	return &s
}

// Add add a new value into disjoint set.
func (s *DisjointSet) Add(val interface{}) {
	if _, present := s.sets[val]; present {
		return
	}
	s.sets[val] = &subset{parent: nil}
}

// Remove remove a exist value from disjoint value.
func (s *DisjointSet) Remove(val interface{}) error {
	val, present := s.sets[val]
	if !present {
		return errors.New("target value doesn't present in disjoint set")
	}
	delete(s.sets, val)
	return nil
}

// find return parent node of specified value.
func (s *DisjointSet) find(val interface{}) interface{} {
	if s.sets[val].parent != val {
		s.sets[val].parent = s.find(val)
	}
	return s.sets[val].parent
}

// Union make two values unite together.
func (s *DisjointSet) Union(x, y interface{}) {
	xRoot := s.find(x)
	yRoot := s.find(y)
	if s.sets[xRoot].rank < s.sets[yRoot].rank {
		s.sets[xRoot].parent = yRoot
	} else if s.sets[xRoot].rank > s.sets[yRoot].rank {
		s.sets[yRoot].parent = xRoot
	} else {
		s.sets[xRoot].parent = yRoot
		s.sets[xRoot].rank++
	}
}
