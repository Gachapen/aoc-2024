package d22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMix(t *testing.T) {
	assert.Equal(t, 37, mix(42, 15))
}

func TestPrune(t *testing.T) {
	assert.Equal(t, 16113920, prune(100000000))
}

func TestEvolve(t *testing.T) {
	assert.Equal(t, 15887950, evolve(123))
}

func TestEvolveTimes(t *testing.T) {
	const times = 2000
	assert.Equal(t, 8685429, evolveTimes(1, times))
	assert.Equal(t, 4700978, evolveTimes(10, times))
	assert.Equal(t, 15273692, evolveTimes(100, times))
	assert.Equal(t, 8667524, evolveTimes(2024, times))
}

func TestPart1Example(t *testing.T) {
	result := SolvePart1("example.txt")
	assert.Equal(t, 37327623, result)
}
