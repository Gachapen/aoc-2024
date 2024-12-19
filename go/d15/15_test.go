package d15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 10092, SolvePart1("example.txt"))
}

func TestSolvePart1Example2(t *testing.T) {
	assert.Equal(t, 9021, SolvePart1("example2.txt"))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 9021, SolvePart2("example.txt"))
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		grid, moves := parse("input.txt", true)
		b.StartTimer()
		solveParsedPart2(&grid, moves)
	}
}
