package graph

import (
	"reflect"
	"testing"
)

type Value struct {
	Test string
}

func TestAdd(t *testing.T) {
	got := New()
	got.Add(1, "test", &Value{Test: "foo"})
	got.Add(2, "thing", 27)

	want := &Graph{
		Vertices: map[uint64]Vertex{
			1: {Id: 1, Label: "test", Value: &Value{Test: "foo"}},
			2: {Id: 2, Label: "thing", Value: 27},
		},
		Edges:     map[uint64]Edge{},
		Adjacency: map[uint64]map[uint64]int{},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected: %v\nGot: %v\n", want, got)
	}
}

func TestEdge(t *testing.T) {
	got := New()
	got.Add(1, "test", "a")
	got.Add(2, "test", "b")
	got.Add(3, "test", "c")

	got.Edge(1, 3, "link")

	want := &Graph{
		Vertices: map[uint64]Vertex{
			1: {Id: 1, Label: "test", Value: "a"},
			2: {Id: 2, Label: "test", Value: "b"},
			3: {Id: 3, Label: "test", Value: "c"},
		},
		Edges: map[uint64]Edge{
			1: {Id: 1, From: 1, To: 3, Label: "link"},
			2: {Id: 2, From: 3, To: 1, Label: "link"},
		},
		Adjacency: map[uint64]map[uint64]int{
			1: {3: 1},
			3: {1: 1},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected: %v\nGot: %v\n", want, got)
	}
}
