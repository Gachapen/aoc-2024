package d09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1928, SolvePart1("example.txt"))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 2858, SolvePart2("example.txt"))
}

func TestDefragment(t *testing.T) {
	assert.Equal(t, []int{0, 2, 2, 1, 1, 1, 2, 2, 2}, defragmentBlocks([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6}, defragmentBlocks([]int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}))
}
