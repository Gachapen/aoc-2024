package d17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	registers := Registers{C: 9}
	runProgram(&registers, []int{2, 6})
	assert.Equal(t, 1, registers.B)
}

func Test2(t *testing.T) {
	registers := Registers{A: 10}
	output := runProgram(&registers, []int{5, 0, 5, 1, 5, 4})
	assert.Equal(t, []int{0, 1, 2}, output)
}

func Test3(t *testing.T) {
	registers := Registers{A: 2024}
	output := runProgram(&registers, []int{0, 1, 5, 4, 3, 0})
	assert.Equal(t, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, output)
	assert.Equal(t, 0, registers.A)
}

func Test4(t *testing.T) {
	registers := Registers{B: 29}
	runProgram(&registers, []int{1, 7})
	assert.Equal(t, 26, registers.B)
}

func Test5(t *testing.T) {
	registers := Registers{B: 2024, C: 43690}
	runProgram(&registers, []int{4, 0})
	assert.Equal(t, 44354, registers.B)
}

func TestPart1Example(t *testing.T) {
	result := SolvePart1("example.txt")
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", result)
}

func TestPart2Example(t *testing.T) {
	result := SolvePart2("example2.txt")
	assert.Equal(t, 117440, result)
}
