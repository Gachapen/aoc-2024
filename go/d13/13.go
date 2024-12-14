package d13

import (
	"aoc-2024/pq"
	. "aoc-2024/vert"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SolveDay13Part1(inputPath string) int {
	file, _ := os.Open(inputPath)

	scanner := bufio.NewScanner(file)

	part := 0
	machine := Machine{}

	machines := make([]Machine, 0, 1)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		switch part {
		case 0:
			machine.buttonA = parseButtonCoords(line)
		case 1:
			machine.buttonB = parseButtonCoords(line)
		case 2:
			machine.goal = parseButtonCoords(line)
			machines = append(machines, machine)
		}

		part = (part + 1) % 3
	}

	file.Close()

	sum := 0
	for _, machine := range machines {
		sum += findLowestCost(machine.buttonA, machine.buttonB, machine.goal)
	}
	return sum
}

func parseButtonCoords(text string) Vertex {
	colonIndex := strings.Index(text, ":")

	xStart := colonIndex + 4
	xEnd := strings.Index(text[xStart:], ",") + xStart

	yStart := xEnd + 4

	xString := text[xStart:xEnd]
	yString := text[yStart:]
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)

	return Vertex{x, y}
}

func findLowestCost(a Vertex, b Vertex, goal Vertex) int {
	openSetQueue := pq.MakePriorityQueue[Visit, int]()
	openSet := make(map[Vertex]struct{})

	start := Vertex{X: 0, Y: 0}
	openSetQueue.PushItem(Visit{start, [2]int{0, 0}}, 0)
	openSet[start] = struct{}{}

	gScores := make(map[Vertex]int, 0)
	// fScores := make(map[Vertex]int, 0)

	for len(openSetQueue) != 0 {
		current := openSetQueue.PopItem()
		currentPos := current.Value.pos

		delete(openSet, current.Value.pos)

		currentCost := gScores[currentPos]

		if currentPos.Equals(goal) {
			return currentCost
		}

		neighbors := []Neighbor{
			{currentPos.Add(a), 0, currentCost + 3},
			{currentPos.Add(b), 1, currentCost + 1},
		}

		for _, neighbor := range neighbors {
			if gScores[neighbor.pos] == 0 || neighbor.cost < gScores[neighbor.pos] {
				gScores[neighbor.pos] = neighbor.cost
				fScore := neighbor.cost + 0

				_, inOpenSet := openSet[neighbor.pos]
				presses := current.Value.presses
				presses[neighbor.button] += 1

				if !inOpenSet && presses[neighbor.button] <= 100 {
					openSetQueue.PushItem(Visit{neighbor.pos, presses}, fScore)
					openSet[neighbor.pos] = struct{}{}
				}
			}
		}
	}

	return 0
}

type Neighbor struct {
	pos    Vertex
	button int
	cost   int
}

type Visit struct {
	pos     Vertex
	presses [2]int
}

type Machine struct {
	buttonA Vertex
	buttonB Vertex
	goal    Vertex
}
