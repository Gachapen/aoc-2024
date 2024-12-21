package d20

import (
	"aoc-2024/grd"
	"aoc-2024/pq"
	"aoc-2024/vert"
	"bufio"
	"fmt"
	"os"
)

type Cheat struct {
	start vert.Vertex
	end   vert.Vertex
}

func SolvePart1(inputPath string) int {
	grid := parse(inputPath)
	start, _ := grid.FindPositionOf('S')
	end, _ := grid.FindPositionOf('E')

	cheatsUsed := make(map[Cheat]bool)

	costWithoutCheats, _ := findCostOfCheapestPathToGoal(&grid, start, end, 0, cheatsUsed)
	// costWithoutCheats := 84

	// const minSaved = 100
	const minSaved = 1
	numCheats := 0

	cost, cheat := findCostOfCheapestPathToGoal(&grid, start, end, 2, cheatsUsed)
	saved := costWithoutCheats - cost
	cheatsUsed[cheat] = true

	for cost != 0 && saved >= minSaved {
		fmt.Println(saved)
		numCheats += 1
		cost, cheat = findCostOfCheapestPathToGoal(&grid, start, end, 2, cheatsUsed)

		saved = costWithoutCheats - cost
		cheatsUsed[cheat] = true
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
	queue := pq.MakePriorityQueue[Visit, int]()
	queue.PushItem(Visit{start, 0, 0, nil}, 0)

	for len(queue) != 0 {
		item := queue.PopItem()

		current := item.Value
		previous := current.previous

		// printPath(grid, &current)

		if current.position == goal {
			// printPath(grid, &current)
			return current.cost, findUsedCheat(&current)
		}

		for _, offset := range []vert.Vertex{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := current.position.Add(offset)

			if !grid.IsOutOfBounds(next) && (previous == nil || next != previous.position) {
				cheatTime := current.cheatTime
				canMove := false

				if grid.GetCellValue(next) != '#' {
					canMove = true
					if cheatTime > 0 && cheatTime < maxCheats {
						cheatTime += 1
					}
				} else if cheatTime < maxCheats {
					if cheatTime < maxCheats-1 {
						canMove = true
					}
					cheatTime += 1
				}

				if cheatTime != current.cheatTime && cheatTime == maxCheats && cheatsUsed[Cheat{start: current.position, end: next}] {
					canMove = false
				}

				if canMove {
					cost := current.cost + 1
					heuristic := current.position.ManhattanDistanceTo(goal)
					priority := cost + heuristic

					queue.PushItem(Visit{next, cost, cheatTime, &current}, priority)
				}
			}
		}
	}

	return 0, Cheat{}
}

func printPath(original *grd.Grid, current *Visit) {
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

func findUsedCheat(endNode *Visit) Cheat {
	current := endNode
	for current.previous != nil {
		previous := current.previous

		if previous.cheatTime != current.cheatTime {
			return Cheat{start: previous.position, end: current.position}
		}

		current = previous
	}

	return Cheat{}
}

type Visit struct {
	position  vert.Vertex
	cost      int
	cheatTime int
	previous  *Visit
}

func parse(inputPath string) grd.Grid {
	return grd.ParseGrid(inputPath)
}
