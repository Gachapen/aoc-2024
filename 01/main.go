package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	leftList, rightList := parse(inputPath)
	sort.Ints(leftList)
	sort.Ints(rightList)

	sum := 0
	for i, left := range leftList {
		right := rightList[i]
		distance := left - right
		if distance < 0 {
			distance = -distance
		}
		sum += distance
	}

	return sum
}

func Part2(inputPath string) int {
	leftList, rightList := parse(inputPath)

	sort.Ints(leftList)
	sort.Ints(rightList)

	sum := 0
	for _, left := range leftList {
		indexInRight := sort.SearchInts(rightList, left)
		if indexInRight != len(rightList) && rightList[indexInRight] == left {
			count := 0
			for j := indexInRight; j < len(rightList) && rightList[j] == left; j++ {
				count += 1
			}

			sum += left * count
		}
	}

	return sum
}

func parse(inputPath string) ([]int, []int) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			splitLine := strings.Split(line, "   ")
			first, _ := strconv.Atoi(splitLine[0])
			second, _ := strconv.Atoi(splitLine[1])

			firstList = append(firstList, first)
			secondList = append(secondList, second)
		}
	}

	return firstList, secondList
}
