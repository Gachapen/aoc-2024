package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test10Part1(t *testing.T) {
	result := Run10Part1("10_example.txt")
	assert.Equal(t, 36, result)
}

func Test10Part2(t *testing.T) {
	result := Run10Part2("10_example.txt")
	assert.Equal(t, 81, result)
}

func BenchmarkRateAllTrailHeads(b *testing.B) {
	grid := ParseGrid("10_kb.txt")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RateAllTrailHeads(&grid)
	}
}
