package graph

import (
	"reflect"
	"testing"
)

type Thing struct {
	Test string
}

func TestAdd(t *testing.T) {
	got := New()
	got.Add(1, "test", &Thing{Test: "foo"})
	got.Add(2, "thing", 27)

	want := &Graph{
		Vertices: map[uint64]Vertex{
			1: {Id: 1, Label: "test", Value: &Thing{Test: "foo"}},
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

func TestValue(t *testing.T) {
	g := New()
	g.Add(1, "test", "a")
	g.Add(2, "test", 23)

	val1, ok := Value[string](g, 1)
	if !ok || val1 != "a" {
		t.Fatalf("Failed to get 1: %v", val1)
	}

	val2, ok := Value[int](g, 2)
	if !ok || val2 != 23 {
		t.Fatalf("Failed to get 2: %v", val2)
	}

	val3, ok := Value[bool](g, 1)
	if ok || val3 != false {
		t.Fatalf("Failed to get 3: %v, %v", ok, val3)
	}

	val4, ok := Value[*Vertex](g, 1)
	if ok || val4 != nil {
		t.Fatalf("Failed to get 4: %v, %v", ok, val4)
	}

	val5, ok := Value[string](g, 5)
	if ok || val5 != "" {
		t.Fatalf("Failed to get 5: %v, %v", ok, val5)
	}
}
