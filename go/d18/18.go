package d18

import (
	"aoc-2024/grd"
	"aoc-2024/vert"
	"bufio"
	"fmt"
	"os"
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

func SolvePart1(inputPath string, gridSize int, numCorruptions int) int {
	corruptions := parse(inputPath)

	grid := grd.MakeGrid(gridSize, gridSize)
	initGrid(&grid, numCorruptions, corruptions)

	start := vert.Vertex{0, 0}
	goal := vert.Vertex{gridSize - 1, gridSize - 1}

	return findCostOfCheapestPathToGoal(&grid, start, goal)
}

func SolvePart2(inputPath string, gridSize int, numInitialCorruptions int) string {
	corruptions := parse(inputPath)

	grid := grd.MakeGrid(gridSize, gridSize)
	initGrid(&grid, numInitialCorruptions, corruptions)

	start := vert.Vertex{0, 0}
	goal := vert.Vertex{gridSize - 1, gridSize - 1}

	for _, corruption := range corruptions[numInitialCorruptions:] {
		grid.SetCellValue(corruption, '#')
		cost := findCostOfCheapestPathToGoal(&grid, start, goal)
		if cost == 0 {
			return fmt.Sprintf("%d,%d", corruption.X, corruption.Y)
		}
	}

	return ""
}

func initGrid(grid *grd.Grid, numCorruptions int, corruptions []vert.Vertex) {
	for i := 0; i < len(grid.Data); i++ {
		grid.Data[i] = '.'
	}

	for i := 0; i < numCorruptions; i++ {
		grid.SetCellValue(corruptions[i], '#')
	}
}

func findCostOfCheapestPathToGoal(grid *grd.Grid, start, goal vert.Vertex) int {
	queue := make([]Node, 1)
	queue[0] = Node{start, 0}

	visited := make(map[vert.Vertex]bool)
	visited[start] = true

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		// grid.SetCellValue(current.position, 'O')
		// grid.Print()
		// grid.SetCellValue(current.position, '.')
		// fmt.Println()
		// input := bufio.NewScanner(os.Stdin)
		// input.Scan()

		if current.position == goal {
			return current.cost
		}

		for _, offset := range []vert.Vertex{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := current.position.Add(offset)

			if !visited[next] && !grid.IsOutOfBounds(next) && grid.GetCellValue(next) != '#' {
				queue = append(queue, Node{next, current.cost + 1})
				visited[next] = true
			}
		}
	}

	return 0
}

func findCutoffByte() {

}

type Node struct {
	position vert.Vertex
	cost     int
}

func parse(inputPath string) []vert.Vertex {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	positions := make([]vert.Vertex, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		positions = append(positions, vert.Vertex{X: x, Y: y})
	}

	return positions
}
