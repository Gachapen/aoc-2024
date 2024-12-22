package d22

import (
	"aoc-2024/vert"
	"bufio"
	"os"
	"strconv"
)

type Cheat struct {
	start vert.Vertex
	end   vert.Vertex
}

func SolvePart1(inputPath string) int {
	secrets := parse(inputPath)
	sum := 0

	for _, secret := range secrets {
		sum += evolveTimes(secret, 2000)
	}

	return sum
}

func mix(secret int, value int) int {
	return value ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}

func evolve(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))
	return secret
}

func evolveTimes(secret int, times int) int {
	for i := 0; i < times; i++ {
		secret = evolve(secret)
	}

	return secret
}

func parse(inputPath string) []int {
	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	secrets := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			secret, _ := strconv.Atoi(line)
			secrets = append(secrets, secret)
		}
	}

	return secrets
}
