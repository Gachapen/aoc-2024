package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	// part1 := Part1("input.txt")
	part2 := Part2("input.txt")
	// fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func Part1(inputPath string) int {
	grid := parse(inputPath)

	guardStartPosition, _ := findGuardPosition(&grid)
	guardStartDirection := Vertex{0, 1}

	visited := make([]Vertex, len(grid.Data))
	RunSimulation(&grid, guardStartPosition, guardStartDirection, visited)

	// printGrid(&grid)

	visitedPositions := 0
	for _, value := range grid.Data {
		if value == 'X' {
			visitedPositions += 1
		}
	}

	return visitedPositions
}

func Part2(inputPath string) int {
	start := time.Now()

	grid := parse(inputPath)

	originalGridData := make([]byte, len(grid.Data))
	copy(originalGridData, grid.Data)

	guardStartPosition, _ := findGuardPosition(&grid)
	guardStartDirection := Vertex{0, 1}
	visited := make([]Vertex, len(grid.Data))

	RunSimulation(&grid, guardStartPosition, guardStartDirection, visited)
	firstSimulationResult := make([]byte, len(grid.Data))
	copy(firstSimulationResult, grid.Data)

	obstacleLoopPositions := 0

	for i, value := range firstSimulationResult {
		if value == 'X' {
			copy(grid.Data, originalGridData)
			zeroOutVertexSlice(visited)

			// fmt.Printf("Testing %d\n", i)
			grid.Data[i] = '#'

			looped := RunSimulation(&grid, guardStartPosition, guardStartDirection, visited)

			// grid.Data[i] = 'O'
			// grid.Data[getGridIndexFromPosition(guardStartPosition, &grid)] = '^'
			// printGrid(&grid)

			if looped {
				obstacleLoopPositions += 1
				// fmt.Println("looped")
			}

			// fmt.Println()
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Part 2 time: %d\n", elapsed.Milliseconds())

	return obstacleLoopPositions
}

func zeroOutVertexSlice(vertices []Vertex) {
	for i := range vertices {
		vertices[i] = Vertex{}
	}
}

func RunSimulation(grid *Grid, guardStartPosition Vertex, guardStartDirection Vertex, visited []Vertex) bool {
	position := guardStartPosition
	direction := guardStartDirection

	for !isOutsideGrid(position, grid) {
		vistedDirection := visited[getGridIndexFromPosition(position, grid)]
		if vistedDirection.X == direction.X && vistedDirection.Y == direction.Y {
			return true
		}

		visited[getGridIndexFromPosition(position, grid)] = direction

		nextPosition := addVertices(position, direction)

		for !isOutsideGrid(nextPosition, grid) && gridPositionHasValue(nextPosition, '#', grid) {
			direction = turnGuardDirection(direction)
			nextPosition = addVertices(position, direction)
		}

		setGridCellValue(position, 'X', grid)

		position = nextPosition
	}

	return false
}

type Vertex struct {
	X int
	Y int
}

type Grid struct {
	Data   []byte
	Width  int
	Height int
}

func getGridPositionFromIndex(index int, grid *Grid) Vertex {
	return Vertex{index % grid.Width, -(index / grid.Height)}
}

func getGridIndexFromPosition(position Vertex, grid *Grid) int {
	return (-position.Y)*grid.Width + position.X
}

func gridPositionHasValue(position Vertex, value byte, grid *Grid) bool {
	return getGridCellValue(position, grid) == value
}

func getGridCellValue(position Vertex, grid *Grid) byte {
	return grid.Data[getGridIndexFromPosition(position, grid)]
}

func isOutsideGrid(position Vertex, grid *Grid) bool {
	return position.X < 0 || position.X >= grid.Width || position.Y > 0 || -position.Y >= grid.Height
}

func setGridCellValue(position Vertex, value byte, grid *Grid) {
	grid.Data[getGridIndexFromPosition(position, grid)] = value
}

func printGrid(grid *Grid) {
	for y := 0; y > -grid.Height; y-- {
		for x := 0; x < grid.Width; x++ {
			fmt.Printf("%c", getGridCellValue(Vertex{x, y}, grid))
		}
		fmt.Println()
	}
}

func addVertices(a Vertex, b Vertex) Vertex {
	return Vertex{a.X + b.X, a.Y + b.Y}
}

func equalVertices(a Vertex, b Vertex) bool {
	return a.X == b.X && a.Y == b.Y
}

func turnGuardDirection(direction Vertex) Vertex {
	if direction.X == 1 {
		return Vertex{0, -1}
	}

	if direction.Y == -1 {
		return Vertex{-1, 0}
	}

	if direction.X == -1 {
		return Vertex{0, 1}
	}

	if direction.Y == 1 {
		return Vertex{1, 0}
	}

	return Vertex{}
}

func parse(inputPath string) Grid {
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

	grid := Grid{data, width, height}

	return grid
}

func findGuardPosition(grid *Grid) (Vertex, error) {
	for i, cell := range grid.Data {
		if cell == '^' {
			return getGridPositionFromIndex(i, grid), nil
		}
	}

	return Vertex{}, errors.New("Found no guard")
}
