package containers

import (
	"gonum.org/v1/gonum/graph/simple"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	set := NewDisjointSet(nil)
	set.Add("hello")
	set.Add("world")
	set.Add("golang")
	set.Union("hello", "world")
	set.Union("hello", "golang")
	println(set.String())
	if set.sets["world"].parent != "golang" {
		t.Error("not expected parent of `world`")
	}
	set.find("hello")
	if set.sets["hello"].parent != "golang" {
		t.Error("not expected parent of `hello`")
	}
}

// Test for simple undirected graph:
//      1->2
//      1->3
//      2->3
func TestHasCycle(t *testing.T) {
	g := simple.NewDirectedGraph()
	n1 := g.NewNode()
	g.AddNode(n1)
	n2 := g.NewNode()
	g.AddNode(n2)
	n3 := g.NewNode()
	g.AddNode(n3)
	e12 := g.NewEdge(n1, n2)
	e13 := g.NewEdge(n1, n3)
	e23 := g.NewEdge(n2, n3)
	g.SetEdge(e12)
	g.SetEdge(e13)
	g.SetEdge(e23)
	if !HasCycle(g) {
		t.Error("graph should have cycle")
	}
	g.RemoveEdge(n2.ID(), n3.ID())
	if HasCycle(g) {
		t.Error("graph should not have cycle")
	}
}
