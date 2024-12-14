package vert

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

func (a Vertex) Equals(b Vertex) bool {
	return a.X == b.X && a.Y == b.Y
}

func (self Vertex) Negate() Vertex {
	return Vertex{-self.X, -self.Y}
}

func (self Vertex) EuclideanLength() int {
	return self.X + self.Y
}

func (self Vertex) Divide(by int) Vertex {
	return Vertex{self.X / by, self.Y / by}
}

func ZeroOutVertexSlice(vertices []Vertex) {
	for i := range vertices {
		vertices[i] = Vertex{}
	}
}
