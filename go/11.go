package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run11Part1(inputPath string) {
	result := run(inputPath, 25)
	fmt.Println("11 Part 1: ", result)
}

func Run11Part2(inputPath string) {
	result := run(inputPath, 75)
	fmt.Println("11 Part 2: ", result)
}

func run(inputPath string, iterations int) int {
	file, _ := os.Open(inputPath)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	file.Close()

	values := strings.Split(line, " ")
	stones := make([]int, len(values))

	for i, value := range values {
		stones[i], _ = strconv.Atoi(value)
	}

	return EvolveStonesRecursive(stones, iterations)
}

func EvolveStonesBF(stones []int, iterations int) int {
	stagingStones := make([]int, 0, len(stones))

	for iteration := 0; iteration < iterations; iteration++ {
		for _, value := range stones {
			if value == 0 {
				stagingStones = append(stagingStones, 1)
			} else {
				digits := strconv.Itoa(value)
				numDigits := len(digits)

				if numDigits%2 == 0 {
					leftDigits := digits[:numDigits/2]
					rightDigits := digits[numDigits/2:]

					leftValue, _ := strconv.Atoi(leftDigits)
					rightValue, _ := strconv.Atoi(rightDigits)
					stagingStones = append(stagingStones, leftValue, rightValue)
				} else {
					stagingStones = append(stagingStones, value*2024)
				}
			}
		}

		temp := stones
		stones = stagingStones
		stagingStones = temp[:0]
	}

	return len(stones)
}

func EvolveStonesRecursive(stones []int, iterations int) int {
	sum := 0
	cache := make(map[string]int)

	for i := 0; i < len(stones); i++ {
		sum += EvolveStone(stones[i], 0, iterations, cache)
	}

	return sum
}

func EvolveStone(value int, iteration int, iterations int, cache map[string]int) int {
	if iteration == iterations {
		return 1
	}

	cacheKey := fmt.Sprintf("%d_%d", value, iterations-iteration)
	cached := cache[cacheKey]
	if cached != 0 {
		return cached
	}

	iteration += 1
	result := 0

	if value == 0 {
		result = EvolveStone(1, iteration, iterations, cache)
	} else {
		digits := strconv.Itoa(value)
		numDigits := len(digits)

		if numDigits%2 == 0 {
			leftDigits := digits[:numDigits/2]
			rightDigits := digits[numDigits/2:]

			leftValue, _ := strconv.Atoi(leftDigits)
			rightValue, _ := strconv.Atoi(rightDigits)
			result = EvolveStone(leftValue, iteration, iterations, cache) + EvolveStone(rightValue, iteration, iterations, cache)
		} else {
			result = EvolveStone(value*2024, iteration, iterations, cache)
		}
	}

	cache[cacheKey] = result

	return result
}

func EvolveStonesDF2(stones []int, iterations int) int {
	sum := 0

	for i := 0; i < len(stones); i++ {
		sum += EvolveStone2(stones[i], iterations)
	}

	return sum
}

func EvolveStone2(value int, iterations int) int {
	stones := make([]State, 1)
	stones[0] = State{value, 0}
	sum := 0

	for len(stones) != 0 {
		state := stones[len(stones)-1]
		stones = stones[:len(stones)-1]

		value = state.value

		for i := state.startIteration; i < iterations; i++ {
			if value == 0 {
				value = 1
			} else {
				digits := strconv.Itoa(value)
				numDigits := len(digits)

				if numDigits%2 == 0 {
					strings.Split(digits, "")
					leftDigits := digits[:numDigits/2]
					rightDigits := digits[numDigits/2:]

					leftValue, _ := strconv.Atoi(leftDigits)
					rightValue, _ := strconv.Atoi(rightDigits)

					value = leftValue
					stones = append(stones, State{rightValue, i + 1})
				} else {
					value *= 2024
				}
			}
		}

		sum += 1
	}

	return sum
}

type State struct {
	value          int
	startIteration int
}
