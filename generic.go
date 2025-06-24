package graph

func (g *Graph) Add(id uint64, label string, value any) {
	if id == 0 {
		id = uint64(len(g.Vertices) + 1)
	}
	g.Vertices[id] = Vertex{
		Id:    id,
		Label: label,
		Value: value,
	}
}

func (g *Graph) Edge(from uint64, to uint64, label string) error {
	return g.AddEdge(Edge{From: from, To: to, Label: label})
}

func Value[T any](g *Graph, id uint64) (T, bool) {
	var empty T

	vertex, ok := g.Vertices[id]
	if !ok {
		return empty, ok
	}

	typedValue, ok := vertex.Value.(T)
	if !ok {
		return empty, ok
	}

	return typedValue, ok
}
