package d18

import (
	"aoc-2024/grd"
	"aoc-2024/vert"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	result := SolvePart1("example.txt", 7, 12)
	assert.Equal(t, 22, result)
}

func TestPart2Example(t *testing.T) {
	result := SolvePart2("example.txt", 7, 12)
	assert.Equal(t, "6,1", result)
}

func BenchmarkFindCostOfCheapestPathToGoal(b *testing.B) {
	b.StopTimer()
	corruptions := parse("input.txt")
	gridSize := 71
	grid := grd.MakeGrid(gridSize, gridSize)

	for i := 0; i < b.N; i++ {
		initGrid(&grid, 1024, corruptions)

		start := vert.Vertex{0, 0}
		goal := vert.Vertex{gridSize - 1, gridSize - 1}

		b.StartTimer()
		findCostOfCheapestPathToGoal(&grid, start, goal)
		b.StopTimer()
	}
}
