package grid

import (
	. "aoc-2024/vert"
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Grid struct {
	Data   []byte
	Width  int
	Height int
}

func MakeGrid(width int, height int) Grid {
	return Grid{
		Data:   make([]byte, width*height),
		Width:  width,
		Height: height,
	}
}

func (grid *Grid) GetPositionFromIndex(index int) Vertex {
	return Vertex{index % grid.Width, index / grid.Height}
}

func (grid *Grid) GetIndexFromPosition(position Vertex) int {
	return position.Y*grid.Width + position.X
}

func PositionHasValue(grid *Grid, position Vertex, value byte) bool {
	return grid.GetCellValue(position) == value
}

func (grid *Grid) GetCellValue(position Vertex) byte {
	return grid.Data[grid.GetIndexFromPosition(position)]
}

func (grid *Grid) IsOutOfBounds(position Vertex) bool {
	return position.X < 0 || position.X >= grid.Width || position.Y < 0 || position.Y >= grid.Height
}

func (grid *Grid) SetCellValue(position Vertex, value byte) {
	grid.Data[grid.GetIndexFromPosition(position)] = value
}

func (grid *Grid) Print() {
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			fmt.Printf("%c", grid.GetCellValue(Vertex{x, y}))
		}
		fmt.Println()
	}
}

func (grid *Grid) FindPositionOf(value byte) (Vertex, error) {
	for i, cell := range grid.Data {
		if cell == value {
			return grid.GetPositionFromIndex(i), nil
		}
	}

	return Vertex{}, errors.New("Not found")
}

func ParseGrid(inputPath string) Grid {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return ParseGridFromScanner(scanner)
}

func ParseGridFromScanner(scanner *bufio.Scanner) Grid {
	data := make([]byte, 0)
	width := 0

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}

		width = len(line)
		data = append(data, line...)
	}

	height := len(data) / width

	grid := Grid{data, width, height}

	return grid
}
