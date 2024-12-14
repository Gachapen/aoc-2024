package d14

import (
	. "aoc-2024/vert"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveDay14Part1(t *testing.T) {
	assert.Equal(t, 12, SolveDayPart1("example.txt", Vertex{11, 7}))
}

func TestFindPositionAfterIterations(t *testing.T) {
	assert.Equal(t, Vertex{4, 1}, findPositionAfterIterations(Vertex{2, 4}, Vertex{2, -3}, Vertex{11, 7}, 1))
	assert.Equal(t, Vertex{6, 5}, findPositionAfterIterations(Vertex{2, 4}, Vertex{2, -3}, Vertex{11, 7}, 2))
	assert.Equal(t, Vertex{1, 3}, findPositionAfterIterations(Vertex{2, 4}, Vertex{2, -3}, Vertex{11, 7}, 5))
}
