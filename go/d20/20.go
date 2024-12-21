package d20

import (
	"aoc-2024/grd"
	"aoc-2024/math"
	"aoc-2024/vert"
	"bufio"
	"fmt"
	"os"
)

type Cheat struct {
	start vert.Vertex
	end   vert.Vertex
}

func Solve(inputPath string, minSaved int, maxCheatMoves int) int {
	grid := parse(inputPath)
	start, _ := grid.FindPositionOf('S')
	end, _ := grid.FindPositionOf('E')

	gridScores := make([]int, len(grid.Data))

	findCostWithoutCheats(&grid, start, end, gridScores)
	// fmt.Println(costWithoutCheats)

	cheatsUsed := make(map[Cheat]bool)

	numCheats := 0
	// saves := make([]int, 0)

	saved, moves, start, previous := findCheat(&grid, start, end, cheatsUsed, gridScores, minSaved, maxCheatMoves, 0, start)

	for saved >= minSaved {
		// fmt.Println(saved)
		numCheats += 1
		// fmt.Println(numCheats)

		saved, moves, start, previous = findCheat(&grid, start, end, cheatsUsed, gridScores, minSaved, maxCheatMoves, moves, previous)
	}

	// slices.Sort(saves)
	// fmt.Println(saves)

	return numCheats
}

func SolvePart2(inputPath string, gridSize int, numInitialCorruptions int) string {
	return ""
}

func findCostWithoutCheats(grid *grd.Grid, start, goal vert.Vertex, gridScores []int) int {
	previous := start
	current := start

	moves := 0

	for current != goal {
		var next vert.Vertex

		for _, offset := range []vert.Vertex{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next = current.Add(offset)
			if next != previous && grid.GetCellValue(next) != '#' {
				break
			}
		}

		previous = current
		current = next

		moves += 1

		gridIndex := grid.GetIndexFromPosition(next)
		gridScores[gridIndex] = moves
	}

	return moves
}

func findCheat(
	grid *grd.Grid,
	start,
	goal vert.Vertex,
	cheatsUsed map[Cheat]bool,
	gridScores []int,
	minSave int,
	maxCheatMoves int,
	moves int,
	previous vert.Vertex,
) (int, int, vert.Vertex, vert.Vertex) {
	current := start

	for current != goal {
		bestMove := findBestMove(grid, current, previous, cheatsUsed, gridScores, moves, minSave, maxCheatMoves)

		if bestMove.saved > 0 {
			cheatsUsed[bestMove.cheat] = true
			return bestMove.saved, moves, current, previous
		}

		moves += bestMove.distance
		previous = current
		current = bestMove.next
	}

	return 0, moves, current, previous
}

func findNextWithoutCheats(grid *grd.Grid, current vert.Vertex, previous vert.Vertex, gridScores []int) vert.Vertex {
	for _, offset := range []vert.Vertex{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		next := current.Add(offset)
		if next != previous && grid.GetCellValue(next) != '#' {
			scoreChange := gridScores[grid.GetIndexFromPosition(next)] - gridScores[grid.GetIndexFromPosition(current)]
			if scoreChange > 0 {
				return next
			}
		}
	}

	panic("no next")
}

func findBestMove(
	grid *grd.Grid,
	current vert.Vertex,
	previous vert.Vertex,
	cheatsUsed map[Cheat]bool,
	gridScores []int,
	currentMoves int,
	minSave int,
	maxCheatDistance int,
) Move {
	bestMove := Move{}

	for offsetY := -maxCheatDistance; offsetY <= maxCheatDistance; offsetY++ {
		absOffsetY := offsetY
		if absOffsetY < 0 {
			absOffsetY = -absOffsetY
		}

		lookWidth := maxCheatDistance - absOffsetY

		for offsetX := -lookWidth; offsetX <= lookWidth; offsetX++ {
			if offsetX == 0 && offsetY == 0 {
				continue
			}

			offset := vert.Vertex{offsetX, offsetY}
			next := current.Add(offset)

			if next == previous || grid.IsOutOfBounds(next) || grid.GetCellValue(next) == '#' {
				continue
			}

			distance := math.AbsInt(offsetX) + math.AbsInt(offsetY)

			if distance == 1 {
				if bestMove.saved == 0 {
					bestMove = Move{next, distance, 0, Cheat{}}
				}
			} else {
				cheat := Cheat{current, next}
				if !cheatsUsed[cheat] {
					moves := currentMoves + distance
					saved := gridScores[grid.GetIndexFromPosition(next)] - moves
					if saved >= minSave && saved > bestMove.saved {
						bestMove = Move{next, distance, saved, cheat}
					}
				}
			}
		}
	}

	return bestMove
}

type Move struct {
	next     vert.Vertex
	distance int
	saved    int
	cheat    Cheat
}

func printPosition(original *grd.Grid, position vert.Vertex) {
	grid := grd.MakeGrid(original.Width, original.Height)
	copy(grid.Data, original.Data)

	head := position
	grid.SetCellValue(head, 'O')

	grid.Print()
	fmt.Println()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
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
