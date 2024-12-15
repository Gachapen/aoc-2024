package d15

import (
	"aoc-2024/grid"
	"aoc-2024/vert"
	"bufio"
	"fmt"
	"os"
)

const interactive = false

func SolvePart1(inputPath string) int {

	grid, moves := parse(inputPath, false)
	position := findStartPosition(&grid)

	if interactive {
		grid.Print()
	}

	for _, move := range moves {
		if interactive {
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

		if interactive {
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

func SolvePart2(inputPath string) int {
	grid, moves := parse(inputPath, true)
	position := findStartPosition(&grid)

	if interactive {
		grid.Print()
	}

	for _, move := range moves {
		if interactive {
			grid.SetCellValue(position, '.')
		}

		direction := getDirectionFromMove(move)
		next := position.Add(direction)

		cellValue := grid.GetCellValue(next)
		if cellValue == '[' || cellValue == ']' {
			if moveWideBoxes(&grid, next, direction) {
				position = next
			}
		} else if cellValue != '#' {
			position = next
		}

		if interactive {
			fmt.Println("Move: ", string(move))
			grid.SetCellValue(position, '@')
			grid.Print()

			input := bufio.NewScanner(os.Stdin)
			input.Scan()
		}
	}

	sum := 0

	for i, value := range grid.Data {
		if value == '[' {
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

func moveWideBoxes(grid *grid.Grid, startPosition vert.Vertex, direction vert.Vertex) bool {
	if direction.Y == 0 {
		position := startPosition
		currentValue := grid.GetCellValue(position)

		for currentValue == '[' || currentValue == ']' {
			position = position.Add(direction)
			currentValue = grid.GetCellValue(position)
		}

		if currentValue == '#' {
			return false
		}

		distance := position.X - startPosition.X
		if distance < 0 {
			distance = -distance
		}

		for i := 0; i < distance; i++ {
			currentX := position.X + i*-direction.X
			previousX := currentX - direction.X
			previousValue := grid.GetCellValue(vert.Vertex{previousX, position.Y})
			grid.SetCellValue(vert.Vertex{currentX, position.Y}, previousValue)
		}

		grid.SetCellValue(startPosition, '.')
	} else {
		return moveWideBoxesVertical(startPosition, direction.Y, grid)
	}

	return true
}

func moveWideBoxesVertical(startPosition vert.Vertex, direction int, grid *grid.Grid) bool {
	queue := make([]vert.Vertex, 1)
	queue[0] = startPosition
	visited := make([]vert.Vertex, 0)

	for len(queue) != 0 {
		position := queue[0]
		queue = queue[1:]

		value := grid.GetCellValue(position)
		if value == ']' {
			position.X -= 1
		}

		visited = append(visited, position)

		nextPositions := []vert.Vertex{{position.X, position.Y + direction}, {position.X + 1, position.Y + direction}}
		nextValues := []byte{0, 0}

		for i, nextPosition := range nextPositions {
			nextValue := grid.GetCellValue(nextPosition)
			nextValues[i] = nextValue

			if nextValue == '#' {
				return false
			} else if nextValue == '[' || nextValue == ']' {
				queue = append(queue, nextPosition)
			}
		}
	}

	for i := len(visited) - 1; i >= 0; i-- {
		position := visited[i]

		grid.SetCellValue(position.Add(vert.Vertex{0, direction}), '[')
		grid.SetCellValue(position.Add(vert.Vertex{1, direction}), ']')

		grid.SetCellValue(position, '.')
		grid.SetCellValue(position.Add(vert.Vertex{1, 0}), '.')
	}

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

func parse(inputPath string, double bool) (grid.Grid, []byte) {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	originalGrid := grid.ParseGridFromScanner(scanner)

	moves := make([]byte, 0)

	for scanner.Scan() {
		line := scanner.Bytes()
		moves = append(moves, line...)
	}

	if double {
		doubledGrid := grid.MakeGrid(originalGrid.Width*2, originalGrid.Height)
		for y := 0; y < originalGrid.Height; y++ {
			for x := 0; x < originalGrid.Width; x++ {
				originalPosition := vert.Vertex{x, y}
				doubledPosition := vert.Vertex{x * 2, y}

				value := originalGrid.GetCellValue(originalPosition)

				if value == '#' || value == '.' {
					doubledGrid.SetCellValue(doubledPosition, value)
					doubledGrid.SetCellValue(doubledPosition.Add(vert.Vertex{1, 0}), value)
				} else if value == '@' {
					doubledGrid.SetCellValue(doubledPosition, value)
					doubledGrid.SetCellValue(doubledPosition.Add(vert.Vertex{1, 0}), '.')
				} else {
					doubledGrid.SetCellValue(doubledPosition, '[')
					doubledGrid.SetCellValue(doubledPosition.Add(vert.Vertex{1, 0}), ']')
				}
			}
		}
		return doubledGrid, moves
	}

	return originalGrid, moves
}

func findStartPosition(grid *grid.Grid) vert.Vertex {
	position, _ := grid.FindPositionOf('@')
	return position
}
