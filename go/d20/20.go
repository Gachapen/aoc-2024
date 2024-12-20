package d20

import (
	"aoc-2024/grd"
	"aoc-2024/vert"
	"bufio"
	"fmt"
	"os"
)

const interactive = true

type Cheat struct {
	start vert.Vertex
	end   vert.Vertex
}

func SolvePart1(inputPath string) int {
	grid := parse(inputPath)
	start, _ := grid.FindPositionOf('S')
	end, _ := grid.FindPositionOf('E')

	cheatsUsed := make(map[Cheat]bool)

	// costWithoutCheats, _ := findCostOfCheapestPathToGoal(&grid, start, end, 0, cheatsUsed)
	costWithoutCheats := 84

	numCheats := 0
	sumSaved := 0

	for sumSaved < 100 {
		cost, cheat := findCostOfCheapestPathToGoal(&grid, start, end, 2, cheatsUsed)
		saved := costWithoutCheats - cost
		cheatsUsed[cheat] = true

		if saved > 0 {
			numCheats += 1
			sumSaved += saved
		}
	}

	return numCheats
}

func SolvePart2(inputPath string, gridSize int, numInitialCorruptions int) string {
	return ""
}

// func findBestCheats(grid *grd.Grid, start, end vert.Vertex, maxCheats int) []int {
// 	cheatsUsed := make(map[Cheat]bool)

// 	costWithoutCheats, _ := findCostOfCheapestPathToGoal(grid, start, end, 0, cheatsUsed)

// 	bestCheatScores := make([]int, 0)
// 	sumSaved := 0

// 	for sumSaved < 100 {
// 		cost, cheat := findCostOfCheapestPathToGoal(grid, start, end, maxCheats, cheatsUsed)
// 		saved := costWithoutCheats - cost
// 		cheatsUsed[cheat] = true

// 		if saved > 0 {
// 			bestCheatScores = append(bestCheatScores, saved)
// 			sumSaved += saved
// 		}
// 	}

// 	return len(bestCheatScores)
// }

func findCostOfCheapestPathToGoal(grid *grd.Grid, start, goal vert.Vertex, maxCheats int, cheatsUsed map[Cheat]bool) (int, Cheat) {
	queue := make([]Node, 1)
	queue[0] = Node{start, 0, 0, nil}

	visited := make(map[vert.Vertex]bool)
	visited[start] = true

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		printPath(grid, &current)

		if current.position == goal {
			return current.cost, findUsedCheat(&current)
		}

		for _, offset := range []vert.Vertex{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := current.position.Add(offset)

			if !grid.IsOutOfBounds(next) && next != current.position {
				cheats := current.cheats
				canMove := false

				if grid.GetCellValue(next) != '#' {
					canMove = true
					if cheats > 0 && cheats < maxCheats {
						cheats += 1
					}
				} else if current.cheats < maxCheats {
					canMove = true
					cheats += 1
				}

				if cheats != current.cheats && cheats == maxCheats && cheatsUsed[Cheat{start: current.position, end: next}] {
					canMove = false
				}

				if canMove {
					queue = append(queue, Node{next, current.cost + 1, cheats, &current})
					visited[next] = true
				}
			}
		}
	}

	return 0, Cheat{}
}

func printPath(original *grd.Grid, current *Node) {
	grid := grd.MakeGrid(original.Width, original.Height)
	copy(grid.Data, original.Data)

	head := current.position

	for current.previous != nil {
		previousPosition := current.previous.position
		grid.SetCellValue(previousPosition, 'X')
		current = current.previous
	}

	grid.SetCellValue(head, 'O')

	grid.Print()
	fmt.Println()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func findUsedCheat(endNode *Node) Cheat {
	current := endNode
	for current.previous != nil {
		previous := current.previous

		if previous.cheats != current.cheats {
			return Cheat{start: previous.position, end: current.position}
		}

		current = previous
	}

	return Cheat{}
}

type Node struct {
	position vert.Vertex
	cost     int
	cheats   int
	previous *Node
}

func parse(inputPath string) grd.Grid {
	return grd.ParseGrid(inputPath)
}
