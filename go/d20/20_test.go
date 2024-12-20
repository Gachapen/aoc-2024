package d20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	result := SolvePart1("example.txt")
	assert.Equal(t, 44, result)
}

func TestPart1ExampleWitoutCheats(t *testing.T) {
	grid := parse("example.txt")
	start, _ := grid.FindPositionOf('S')
	end, _ := grid.FindPositionOf('E')

	cheatsUsed := make(map[Cheat]bool)

	cost, _ := findCostOfCheapestPathToGoal(&grid, start, end, 0, cheatsUsed)
	assert.Equal(t, 84, cost)
}

// func TestPart1ExampleBestCheats(t *testing.T) {
// 	grid := parse("example.txt")
// 	start, _ := grid.FindPositionOf('S')
// 	end, _ := grid.FindPositionOf('E')

// 	bestCheats := findBestCheats(&grid, start, end, 2)

// 	slices.Sort(bestCheats)
// 	assert.Equal(t, 2, bestCheats[0])
// 	assert.Equal(t, 64, bestCheats[len(bestCheats)-1])
// 	assert.Equal(t, 44, len(bestCheats))
// }

func TestPart2Example(t *testing.T) {
	result := SolvePart2("example.txt", 7, 12)
	assert.Equal(t, "6,1", result)
}
