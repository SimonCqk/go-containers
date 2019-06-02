package containers

import (
	"bytes"
	"errors"
	"fmt"
	"gonum.org/v1/gonum/graph"
)

type DisjointSet struct {
	sets map[interface{}]*rankedNode
}

type rankedNode struct {
	rank   int
	parent interface{}
}

// NewDisjointSet create a new disjoint set with init values.
func NewDisjointSet(vals []interface{}) *DisjointSet {
	s := DisjointSet{sets: make(map[interface{}]*rankedNode)}
	for _, val := range vals {
		s.Add(val)
	}
	return &s
}

// Add add a new value into disjoint set.
func (s *DisjointSet) Add(val interface{}) {
	if _, present := s.sets[val]; present {
		return
	}
	s.sets[val] = &rankedNode{parent: val}
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
		s.sets[val].parent = s.find(s.sets[val].parent)
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

func (s *DisjointSet) String() string {
	var buff bytes.Buffer
	for val, node := range s.sets {
		buff.WriteString(fmt.Sprintf("%v -> %v\n", val, node.parent))
	}
	return buff.String()
}

// HasCycle detect whether cycle exists in undirected graph.
func HasCycle(graph graph.Graph) bool {
	set := NewDisjointSet(nil)
	nodes := graph.Nodes()
	if nodes.Len() == 0 {
		return false
	}
	for nodes.Next() {
		node := nodes.Node()
		// find nodes directly connected with this source node.
		toNodes := graph.From(node.ID())
		if toNodes.Len() == 0 {
			continue
		}
		set.Add(node.ID())
		for toNodes.Next() {
			toNode := toNodes.Node()
			edge := graph.Edge(node.ID(), toNode.ID())
			set.Add(edge.To().ID())
			x := set.find(edge.From().ID())
			y := set.find(edge.To().ID())
			if x == y {
				return true
			}
			set.Union(x, y)
		}
	}
	return false
}
