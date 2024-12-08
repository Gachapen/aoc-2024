package main

type Vertex struct {
	X int
	Y int
}

func (a Vertex) Add(b Vertex) Vertex {
	return Vertex{a.X + b.X, a.Y + b.Y}
}

func (a Vertex) Sub(b Vertex) Vertex {
	return Vertex{a.X - b.X, a.Y - b.Y}
}

func Equals(a Vertex, b Vertex) bool {
	return a.X == b.X && a.Y == b.Y
}

func ZeroOutVertexSlice(vertices []Vertex) {
	for i := range vertices {
		vertices[i] = Vertex{}
	}
}
