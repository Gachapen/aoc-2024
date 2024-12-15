package d09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SolvePart1(inputPath string) int {
	diskMap := parse(inputPath)
	defragmented := defragmentBlocks(diskMap)
	return checksum(defragmented)
}

func SolvePart2(inputPath string) int {
	diskMap := parse(inputPath)
	defragmented := defragmentFiles(diskMap)
	return checksum(defragmented)
}

func parse(inputPath string) []int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, "")
	diskMap := make([]int, len(parts))

	for i, part := range parts {
		diskMap[i], _ = strconv.Atoi(part)
	}

	return diskMap
}

func defragmentBlocks(diskMap []int) []int {
	originalBlocks := getBlocksFromDiskMap(diskMap)
	defragmentedBlocks := make([]int, 0)

	endIndex := len(originalBlocks) - 1

	for i, id := range originalBlocks {
		if i > endIndex {
			break
		}

		if id != -1 {
			defragmentedBlocks = append(defragmentedBlocks, originalBlocks[i])
		} else {
			defragmentedBlocks = append(defragmentedBlocks, originalBlocks[endIndex])
			endIndex -= 1
			for originalBlocks[endIndex] == -1 {
				endIndex -= 1
			}
		}
	}

	return defragmentedBlocks
}

func defragmentFiles(diskMap []int) []int {
	blocks := getBlocksFromDiskMap(diskMap)
	freeSpaces := getFreeSpaces(diskMap)

	blockIndex := len(blocks)

	for diskMapIndex := len(diskMap) - 1; diskMapIndex >= 0; diskMapIndex-- {
		free := diskMapIndex%2 != 0
		fileLength := diskMap[diskMapIndex]
		blockIndex -= fileLength

		if !free {
			freeSpace := findFreeSpace(freeSpaces, fileLength)
			if freeSpace != nil && freeSpace.start < blockIndex {
				moveFile(blocks, blockIndex, freeSpace.start, fileLength)
				freeSpace.start += fileLength
				freeSpace.length -= fileLength
			}
		}
	}

	end := len(blocks)
	for i := end - 1; i >= 0 && blocks[i] == -1; i-- {
		end = i
	}

	return blocks[:end]
}

func moveFile(blocks []int, from int, to int, length int) {
	for i := 0; i < length; i++ {
		blocks[to+i] = blocks[from+i]
		blocks[from+i] = -1
	}
}

func findFreeSpace(freeSpaces []FreeSpace, requiredLength int) *FreeSpace {
	for i, space := range freeSpaces {
		if space.length >= requiredLength {
			return &freeSpaces[i]
		}
	}

	return nil
}

func getBlocksFromDiskMap(diskMap []int) []int {
	blocks := make([]int, 0, len(diskMap))

	id := 0

	for i, length := range diskMap {
		free := i%2 != 0

		for j := 0; j < length; j++ {
			if free {
				blocks = append(blocks, -1)
			} else {
				blocks = append(blocks, id)
			}
		}

		if free {
			id += 1
		}
	}

	return blocks
}

func getFreeSpaces(diskMap []int) []FreeSpace {
	freeSpaces := make([]FreeSpace, 0)
	blockIndex := 0

	for i, length := range diskMap {
		free := i%2 != 0

		if free && length != 0 {
			freeSpace := FreeSpace{start: blockIndex, length: length}
			freeSpaces = append(freeSpaces, freeSpace)
		}

		blockIndex += length
	}

	return freeSpaces
}

func checksum(blocks []int) int {
	sum := 0
	for i, id := range blocks {
		if id != -1 {
			sum += i * id
		}
	}
	return sum
}

type FreeSpace struct {
	start  int
	length int
}

func pop[T any](slice *[]T) T {
	original := *slice
	value := original[len(original)-1]
	*slice = original[:len(original)-1]

	return value
}

func swapRemove[T any](slice []T, index int) []T {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
