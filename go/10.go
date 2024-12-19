package main

import (
	. "aoc-2024/grd"
	. "aoc-2024/vert"
	"fmt"
	"time"
)

func Run10Part1(inputPath string) int {
	grid := ParseGrid(inputPath)
	score, _ := ScoreAllTrailHeads(&grid)

	fmt.Println("10 Part 1:", score)
	return score
}

func Run10Part2(inputPath string) int {
	grid := ParseGrid(inputPath)

	start := time.Now()
	_, rating := ScoreAllTrailHeads(&grid)
	end := time.Now()

	elapsed := end.Sub(start)
	fmt.Println("Time:", elapsed.Microseconds())

	fmt.Println("10 Part 2:", rating)
	return rating
}

func ScoreAllTrailHeads(grid *Grid) (int, int) {
	score := 0
	rating := 0

	for i, value := range grid.Data {
		if value == '0' {
			trailHeadScore, trailHeadRating := scoreTrailHead(grid, i)
			score += trailHeadScore
			rating += trailHeadRating
		}
	}

	return score, rating
}

func RateAllTrailHeads(grid *Grid) int {
	rating := 0
	stack := make([]int, 1, 8)

	for i, value := range grid.Data {
		if value == '0' {
			trailHeadRating := rateTrailHead(grid, i, stack)
			rating += trailHeadRating
		}
	}

	return rating
}

func scoreTrailHead(grid *Grid, startIndex int) (int, int) {
	stack := make([]int, 1)
	stack[0] = startIndex

	topVisited := make([]bool, len(grid.Data))
	numTrails := 0

	for len(stack) != 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentHeight := grid.Data[index]
		nextHeight := currentHeight + 1
		position := grid.GetPositionFromIndex(index)

		for _, direction := range []Vertex{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nextPosition := position.Add(direction)
			if !grid.IsOutOfBounds(nextPosition) {
				index = grid.GetIndexFromPosition(nextPosition)
				if grid.Data[index] == nextHeight {
					// fmt.Println(nextPosition, (string)([]byte{nextHeight}))
					if nextHeight == '9' {
						topVisited[index] = true
						numTrails += 1
					} else {
						stack = append(stack, index)
					}
				}
			}
		}
	}

	numUniqueTops := 0
	for _, visited := range topVisited {
		if visited {
			numUniqueTops += 1
		}
	}

	return numUniqueTops, numTrails
}

func rateTrailHead(grid *Grid, startIndex int, stack []int) int {
	stack = stack[:1]
	stack[0] = startIndex

	numTrails := 0

	for len(stack) != 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentHeight := grid.Data[index]
		nextHeight := currentHeight + 1
		position := grid.GetPositionFromIndex(index)

		for _, direction := range []Vertex{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nextPosition := position.Add(direction)
			if !grid.IsOutOfBounds(nextPosition) {
				index = grid.GetIndexFromPosition(nextPosition)
				if grid.Data[index] == nextHeight {
					if nextHeight == '9' {
						numTrails += 1
					} else {
						stack = append(stack, index)
					}
				}
			}
		}
	}

	return numTrails
}
