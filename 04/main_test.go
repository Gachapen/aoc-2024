package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1("example.txt")
	assert.Equal(t, 18, result)
}

func TestPart2(t *testing.T) {
	result := Part2("example.txt")
	assert.Equal(t, 9, result)
}
