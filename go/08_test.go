package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test08Part1(t *testing.T) {
	result := Run08Part1("08_example.txt")
	assert.Equal(t, 14, result)
}

func Test08Part2(t *testing.T) {
	result := Run08Part2("08_example.txt")
	assert.Equal(t, 34, result)
}
