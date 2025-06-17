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
