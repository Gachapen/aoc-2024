package d15

import (
	"aoc-2024/grid"
	"aoc-2024/vert"
	"bufio"
	"fmt"
	"os"
)

func SolvePart1(inputPath string) int {
	const print = true

	grid, moves := parse(inputPath)
	position := findStartPosition(&grid)

	if print {
		grid.Print()
	}

	for _, move := range moves {
		if print {
			grid.SetCellValue(position, '.')
		}

		direction := getDirectionFromMove(move)
		next := position.Add(direction)

		cellValue := grid.GetCellValue(next)
		if cellValue == 'O' {
			if moveBoxes(&grid, next, direction) {
				position = next
			}
		} else if cellValue != '#' {
			position = next
		}

		if print {
			fmt.Println("Move: ", string(move))
			grid.SetCellValue(position, '@')
			grid.Print()

			input := bufio.NewScanner(os.Stdin)
			input.Scan()
		}
	}

	sum := 0

	for i, value := range grid.Data {
		if value == 'O' {
			pos := grid.GetPositionFromIndex(i)
			sum += 100*pos.Y + pos.X
		}
	}

	return sum
}

func moveBoxes(grid *grid.Grid, startPosition vert.Vertex, direction vert.Vertex) bool {
	position := startPosition
	currentValue := grid.GetCellValue(position)

	for currentValue == 'O' {
		position = position.Add(direction)
		currentValue = grid.GetCellValue(position)
	}

	if currentValue == '#' {
		return false
	}

	grid.SetCellValue(position, 'O')
	grid.SetCellValue(startPosition, '.')

	return true
}

func getDirectionFromMove(move byte) vert.Vertex {
	switch move {
	case '<':
		return vert.Vertex{X: -1, Y: 0}
	case '>':
		return vert.Vertex{X: 1, Y: 0}
	case 'v':
		return vert.Vertex{X: 0, Y: 1}
	case '^':
		return vert.Vertex{X: 0, Y: -1}
	default:
		panic("Unknown move")
	}
}

func SolvePart2(inputPath string) int {
	return 0
}

func parse(inputPath string) (grid.Grid, []byte) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := grid.ParseGridFromScanner(scanner)
	moves := make([]byte, 0)

	for scanner.Scan() {
		line := scanner.Bytes()
		moves = append(moves, line...)
	}

	return grid, moves
}

func findStartPosition(grid *grid.Grid) vert.Vertex {
	position, _ := grid.FindPositionOf('@')
	return position
}
