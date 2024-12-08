package main

import "fmt"

func Run08Part1(inputPath string) int {
	grid := ParseGrid(inputPath)

	numFrequencyTypes := 'z' - '0' + 1
	frequencyPositions := make([][]Vertex, numFrequencyTypes)

	for i, value := range grid.Data {
		if value != '.' {
			frequencyIndex := getFrequencyIndexFromValue(value)

			savedPositions := frequencyPositions[frequencyIndex]
			if savedPositions == nil {
				savedPositions = make([]Vertex, 0, 1)
			}

			position := grid.GetPositionFromIndex(i)
			savedPositions = append(savedPositions, position)
			frequencyPositions[frequencyIndex] = savedPositions
		}
	}

	antinodes := make([]bool, len(grid.Data))

	for _, positions := range frequencyPositions {
		if positions == nil {
			continue
		}

		for firstIndex, firstPosition := range positions {
			for _, secondPosition := range positions[firstIndex+1:] {
				diff := firstPosition.Sub(secondPosition)

				firstAntinode := firstPosition.Add(diff)
				firstGridIndex := grid.GetIndexFromPosition(firstAntinode)
				if !grid.IsOutOfBounds(firstAntinode) && !antinodes[firstGridIndex] {
					antinodes[firstGridIndex] = true
				}

				secondAntinode := secondPosition.Sub(diff)
				secondGridIndex := grid.GetIndexFromPosition(secondAntinode)
				if !grid.IsOutOfBounds(secondAntinode) && !antinodes[secondGridIndex] {
					antinodes[secondGridIndex] = true
				}
			}
		}
	}

	count := 0
	for _, antinode := range antinodes {
		if antinode {
			count += 1
		}
	}

	fmt.Println("08 Part 1:", count)
	return count
}

func Run08Part2(inputPath string) int {
	grid := ParseGrid(inputPath)

	numFrequencyTypes := 'z' - '0' + 1
	frequencyPositions := make([][]Vertex, numFrequencyTypes)

	for i, value := range grid.Data {
		if value != '.' {
			frequencyIndex := getFrequencyIndexFromValue(value)

			savedPositions := frequencyPositions[frequencyIndex]
			if savedPositions == nil {
				savedPositions = make([]Vertex, 0, 1)
			}

			position := grid.GetPositionFromIndex(i)
			savedPositions = append(savedPositions, position)
			frequencyPositions[frequencyIndex] = savedPositions
		}
	}

	antinodes := make([]bool, len(grid.Data))

	for _, positions := range frequencyPositions {
		if positions == nil {
			continue
		}

		for firstIndex, firstPosition := range positions {
			for _, secondPosition := range positions[firstIndex+1:] {
				diff := firstPosition.Sub(secondPosition)

				firstAntinode := firstPosition
				for !grid.IsOutOfBounds(firstAntinode) {
					firstGridIndex := grid.GetIndexFromPosition(firstAntinode)
					if !antinodes[firstGridIndex] {
						antinodes[firstGridIndex] = true
					}
					firstAntinode = firstAntinode.Add(diff)
				}

				secondAntinode := secondPosition
				for !grid.IsOutOfBounds(secondAntinode) {
					secondGridIndex := grid.GetIndexFromPosition(secondAntinode)
					if !antinodes[secondGridIndex] {
						antinodes[secondGridIndex] = true
					}
					secondAntinode = secondAntinode.Sub(diff)
				}
			}
		}
	}

	count := 0
	for i, antinode := range antinodes {
		if antinode {
			count += 1
			grid.Data[i] = '#'
		}
	}

	fmt.Println("08 Part 2:", count)
	return count
}

func getFrequencyIndexFromValue(value byte) int {
	return (int)(value - '0')
}
