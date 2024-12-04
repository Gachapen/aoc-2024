package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1 := Part1("input.txt")
	part2 := Part2("input.txt")
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	data, width, height := Parse(inputPath)
	matches := 0

	for index, value := range data {
		if value == 'X' {
			matches += FindXmasCount(data, index, width, height)
		}
	}

	return matches
}

func Part2(inputPath string) int {
	data, width, height := Parse(inputPath)
	matches := 0

	for index, value := range data {
		if value == 'A' {
			if IsMasCross(data, index, width, height) {
				matches += 1
			}
		}
	}

	return matches
}

func Parse(inputPath string) ([]byte, int, int) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	data := make([]byte, 0)
	width := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		if len(line) != 0 {
			width = len(line)
			data = append(data, line...)
		}
	}

	height := len(data) / width

	return data, width, height
}

func FindXmasCount(data []byte, startIndex int, width int, height int) int {
	startX := startIndex % width
	startY := startIndex / width

	matches := 0

	for dirX := -1; dirX <= 1; dirX++ {
		for dirY := -1; dirY <= 1; dirY++ {
			x := startX + dirX
			y := startY + dirY
			index := y*width + x

			if MatchAtPosition('M', x, y, data, width, height) && MatchRemainder(data, index, dirX, dirY, width, height) {
				fmt.Println(startX, startY, dirX, dirY)
				matches += 1
			}
		}
	}

	return matches
}

func MatchRemainder(data []byte, startIndex int, dirX int, dirY int, width int, height int) bool {
	toMatch := [2]byte{'A', 'S'}

	startX := startIndex % width
	startY := startIndex / width

	for i := 0; i < len(toMatch); i++ {
		x := startX + dirX*(i+1)
		y := startY + dirY*(i+1)

		if !MatchAtPosition(toMatch[i], x, y, data, width, height) {
			return false
		}
	}

	return true
}

func IsMasCross(data []byte, startIndex int, width int, height int) bool {
	startX := startIndex % width
	startY := startIndex / width

	if startX < 1 || startX >= width-1 || startY < 1 || startY >= height-1 {
		return false
	}

	topLeft := data[(startY-1)*width+(startX-1)]
	bottomRight := data[(startY+1)*width+(startX+1)]
	topRight := data[(startY-1)*width+(startX+1)]
	bottomLeft := data[(startY+1)*width+(startX-1)]

	return IsMasMatch(topLeft, bottomRight) && IsMasMatch(topRight, bottomLeft)
}

func MatchAtPosition(element byte, x int, y int, data []byte, width int, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height && data[y*width+x] == element
}

func IsMasMatch(first byte, second byte) bool {
	return (first == 'M' && second == 'S') || (first == 'S' && second == 'M')
}
