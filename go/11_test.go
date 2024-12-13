package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvolveStonesBF(t *testing.T) {
	result1 := EvolveStonesBF([]int{0, 1, 10, 99, 999}, 1)
	assert.Equal(t, 7, result1)

	result2 := EvolveStonesBF([]int{125, 17}, 6)
	assert.Equal(t, 22, result2)

	result3 := EvolveStonesBF([]int{125, 17}, 25)
	assert.Equal(t, 55312, result3)
}

func TestEvolveStonesRecursive(t *testing.T) {
	result1 := EvolveStonesRecursive([]int{0, 1, 10, 99, 999}, 1)
	assert.Equal(t, 7, result1)

	result2 := EvolveStonesRecursive([]int{125, 17}, 6)
	assert.Equal(t, 22, result2)

	result3 := EvolveStonesRecursive([]int{125, 17}, 25)
	assert.Equal(t, 55312, result3)
}

func TestEvolveStonesDF2(t *testing.T) {
	result1 := EvolveStonesDF2([]int{0, 1, 10, 99, 999}, 1)
	assert.Equal(t, 7, result1)

	result2 := EvolveStonesDF2([]int{125, 17}, 6)
	assert.Equal(t, 22, result2)

	result3 := EvolveStonesDF2([]int{125, 17}, 25)
	assert.Equal(t, 55312, result3)
}

func BenchmarkEvolveStonesRecursivePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvolveStonesRecursive([]int{28, 4, 3179, 96938, 0, 6617406, 490, 816207}, 75)
	}
}

func BenchmarkEvolveStonesRecursivePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvolveStonesRecursive([]int{28, 4, 3179, 96938, 0, 6617406, 490, 816207}, 25)
	}
}

func BenchmarkEvolveStonesBFPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvolveStonesBF([]int{28, 4, 3179, 96938, 0, 6617406, 490, 816207}, 25)
	}
}
