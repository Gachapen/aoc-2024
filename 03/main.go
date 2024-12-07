package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1 := Part1("input.txt")
	fmt.Println("Part 1", part1)
	part2 := Part2("input.txt")
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()

		num1 := 0
		num2 := 0

		for {
			match := false

			match, text = scanMulStart(text)
			if !match {
				break
			}

			match, num1, text = scanNumber(text)
			if !match {
				continue
			}

			if text[0] != ',' {
				continue
			}

			text = text[1:]

			match, num2, text = scanNumber(text)
			if !match {
				continue
			}

			if text[0] != ')' {
				continue
			}

			mul := num1 * num2
			fmt.Println(num1, num2, mul)

			sum += mul
		}
	}

	return sum
}

func scanMulStart(text string) (bool, string) {
	mulStartIndex := strings.Index(text, "mul(")
	if mulStartIndex == -1 {
		return false, ""
	}

	return true, text[mulStartIndex+4:]
}

type Instruction int

const (
	None Instruction = -1
	Do   Instruction = iota
	Dont
	Mul
)

func scanForInstruction(text string) (Instruction, string) {
	const MulLen = 4
	const DoLen = 4
	const DontLen = 7

	for i := 0; i < len(text); i++ {
		if text[i] == 'm' {
			if len(text)-i >= MulLen && text[i:i+MulLen] == "mul(" {
				return Mul, text[i+MulLen:]
			}
		} else if text[i] == 'd' {
			remainingLength := len(text) - i
			if remainingLength >= DoLen && text[i:i+DoLen] == "do()" {
				return Do, text[i+DoLen:]
			}
			if remainingLength >= DontLen && text[i:i+DontLen] == "don't()" {
				return Dont, text[i+DontLen:]
			}
		}
	}

	return None, ""
}

func scanNumber(text string) (bool, int, string) {
	digits := make([]byte, 0)

	index := 0
	for text[index] >= 48 && text[index] <= 57 {
		digits = append(digits, text[index])
		index++
	}

	if index == 0 {
		return false, 0, ""
	}

	number, _ := strconv.Atoi(string(digits))
	return true, number, text[index:]
}

func Part2(inputPath string) int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	enabled := true

	for scanner.Scan() {
		text := scanner.Text()

		num1 := 0
		num2 := 0

		for {
			match := false
			instruction := None

			instruction, text = scanForInstruction(text)
			if instruction == None {
				break
			}
			if instruction == Do {
				enabled = true
				continue
			}
			if instruction == Dont {
				enabled = false
				continue
			}

			match, num1, text = scanNumber(text)
			if !match {
				continue
			}

			if text[0] != ',' {
				continue
			}

			text = text[1:]

			match, num2, text = scanNumber(text)
			if !match {
				continue
			}

			if text[0] != ')' {
				continue
			}

			text = text[1:]

			mul := num1 * num2

			if enabled {
				fmt.Println(num1, num2, mul)
				sum += mul
			}
		}
	}

	return sum
}
