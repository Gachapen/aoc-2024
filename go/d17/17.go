package d17

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const interactive = true

type Registers struct {
	A int
	B int
	C int
}

const (
	adv int = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func SolvePart1(inputPath string) string {
	registers, program := parse(inputPath)
	output := runProgram(&registers, program)

	result := strconv.Itoa(output[0])
	for _, value := range output[1:] {
		result += "," + strconv.Itoa(value)
	}
	return result
}

func SolvePart2(inputPath string) int {
	registers, program := parse(inputPath)

	initialRegisters := registers
	initialRegisters.A = 0

	testValue := 215000000000000
	output := runProgram(&registers, program)

	for !slices.Equal(output, program) {
		registers = initialRegisters

		testValue += 1
		registers.A = testValue

		output = runProgram(&registers, program)

		if interactive {
			fmt.Println(testValue)
			fmt.Println(output)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
		}
	}

	return testValue
}

func runProgram(registers *Registers, program []int) []int {
	instructionPointer := 0
	output := make([]int, 0)

	for instructionPointer < len(program) {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]

		nextInstruction := runInstruction(opcode, operand, registers, &output)

		if nextInstruction != -1 {
			instructionPointer = nextInstruction
		} else {
			instructionPointer += 2
		}
	}

	return output
}

func runInstruction(opcode, operand int, registers *Registers, output *[]int) int {
	switch opcode {
	case adv:
		numerator := registers.A
		comboOperandValue := getComboOperandValue(operand, registers)
		denominator := twoToThePowerOf(comboOperandValue)
		registers.A = numerator / denominator
	case bxl:
		registers.B = registers.B ^ operand
	case bst:
		registers.B = getComboOperandValue(operand, registers) % 8
	case jnz:
		if registers.A != 0 {
			return operand
		}
	case bxc:
		registers.B ^= registers.C
	case out:
		value := getComboOperandValue(operand, registers) % 8
		*output = append(*output, value)
	case bdv:
		numerator := registers.A
		comboOperandValue := getComboOperandValue(operand, registers)
		denominator := twoToThePowerOf(comboOperandValue)
		registers.B = numerator / denominator
	case cdv:
		numerator := registers.A
		comboOperandValue := getComboOperandValue(operand, registers)
		denominator := twoToThePowerOf(comboOperandValue)
		registers.C = numerator / denominator
	}

	return -1
}

func twoToThePowerOf(power int) int {
	if power < 0 {
		panic("negative power")
	}

	if power == 0 || power == 1 {
		return 2
	}

	return 2 << (power - 1)
}

func getComboOperandValue(operand int, registers *Registers) int {
	if operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return registers.A
	case 5:
		return registers.B
	case 6:
		return registers.C
	}

	panic("Invalid combo operand")
}

func parse(inputPath string) (Registers, []int) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	registers := Registers{}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	registers.A = parseRegisterValue(scanner.Text())

	scanner.Scan()
	registers.B = parseRegisterValue(scanner.Text())

	scanner.Scan()
	registers.C = parseRegisterValue(scanner.Text())

	scanner.Scan()
	scanner.Scan()
	instructions := parseInstructions(scanner.Text())

	return registers, instructions
}

func parseRegisterValue(line string) int {
	colonIndex := strings.Index(line, ": ")
	valuePart := line[colonIndex+2:]
	value, _ := strconv.Atoi(valuePart)
	return value
}

func parseInstructions(line string) []int {
	colonIndex := strings.Index(line, ": ")
	instructionsPart := line[colonIndex+2:]
	instructionParts := strings.Split(instructionsPart, ",")

	instructions := make([]int, len(instructionParts))
	for i, part := range instructionParts {
		instruction, _ := strconv.Atoi(part)
		instructions[i] = instruction
	}

	return instructions
}
