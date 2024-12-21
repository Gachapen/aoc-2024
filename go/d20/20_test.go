package d20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	result := Solve("example.txt", 1, 2)
	assert.Equal(t, 44, result)
}

func TestPart2Example(t *testing.T) {
	result := Solve("example.txt", 50, 20)
	assert.Equal(t, 285, result)
}

func TestPart1ExampleWithoutCheats(t *testing.T) {
	grid := parse("example.txt")
	start, _ := grid.FindPositionOf('S')
	end, _ := grid.FindPositionOf('E')

	gridScores := make([]int, len(grid.Data))

	cost := findCostWithoutCheats(&grid, start, end, gridScores)
	assert.Equal(t, 84, cost)
}
