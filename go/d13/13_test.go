package d13

import (
	. "aoc-2024/vert"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveDay3Part1(t *testing.T) {
	assert.Equal(t, 480, SolveDay13Part1("example.txt"))
}

func TestFindLowestCost(t *testing.T) {
	assert.Equal(t, 280, findLowestCost(Vertex{94, 34}, Vertex{22, 67}, Vertex{8400, 5400}))
	assert.Equal(t, 0, findLowestCost(Vertex{26, 66}, Vertex{67, 21}, Vertex{12748, 12176}))
	assert.Equal(t, 200, findLowestCost(Vertex{17, 86}, Vertex{84, 37}, Vertex{7870, 6450}))
	assert.Equal(t, 0, findLowestCost(Vertex{69, 23}, Vertex{27, 71}, Vertex{18641, 10279}))
}
