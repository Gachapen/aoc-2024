package d14

import (
	"aoc-2024/grd"
	. "aoc-2024/vert"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SolveDayPart1(inputPath string, gridSize Vertex) int {
	robots := parse(inputPath)
	iterations := 100

	divisionPoint := gridSize.Divide(2)

	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0

	for _, robot := range robots {
		position := findPositionAfterIterations(robot.position, robot.velocity, gridSize, iterations)

		if position.X < divisionPoint.X && position.Y < divisionPoint.Y {
			topLeft += 1
		} else if position.X < divisionPoint.X && position.Y > divisionPoint.Y {
			bottomLeft += 1
		} else if position.X > divisionPoint.X && position.Y < divisionPoint.Y {
			topRight += 1
		} else if position.X > divisionPoint.X && position.Y > divisionPoint.Y {
			bottomRight += 1
		}
	}

	return topLeft * topRight * bottomLeft * bottomRight
}

func SolveDayPart2(inputPath string, gridSize Vertex) int {
	robots := parse(inputPath)

	// divisionPoint := gridSize.Divide(2)

	grid := grd.MakeGrid(gridSize.X, gridSize.Y)

	for i := 0; true; i++ {
		for j := 0; j < len(grid.Data); j++ {
			grid.Data[j] = '.'
		}

		for _, robot := range robots {
			position := findPositionAfterIterations(robot.position, robot.velocity, gridSize, i)
			grid.SetCellValue(position, 'X')
		}

		if (i-27)%103 == 0 && (i-52)%101 == 0 {
			grid.Print()
			// fmt.Println(i)
			// input := bufio.NewScanner(os.Stdin)
			// input.Scan()
			return i
		}
	}

	return 0
}

func parse(inputPath string) []Robot {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := make([]Robot, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")

		robots = append(robots, Robot{
			position: parseVertex(parts[0][2:]),
			velocity: parseVertex(parts[1][2:]),
		})
	}

	return robots
}

func parseVertex(text string) Vertex {
	parts := strings.Split(text, ",")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return Vertex{x, y}
}

func findPositionAfterIterations(position Vertex, velocity Vertex, gridSize Vertex, iterations int) Vertex {
	x := position.X + velocity.X*iterations
	y := position.Y + velocity.Y*iterations

	x = x % gridSize.X
	y = y % gridSize.Y

	if x < 0 {
		x = gridSize.X + x
	}

	if y < 0 {
		y = gridSize.Y + y
	}

	return Vertex{x, y}
}

type Robot struct {
	position Vertex
	velocity Vertex
}
