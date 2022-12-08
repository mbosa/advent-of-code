package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pair1, pair2 := parseLine(line)
		if fullyContained(pair1, pair2) || fullyContained(pair2, pair1) {
			resPart1++
		}
		if overlap(pair1, pair2) {
			resPart2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	return resPart1, resPart2
}

func parseLine(line string) (pair1, pair2 [2]int) {
	pairs := [2][2]int{}
	for i, pairStr := range strings.Split(line, ",") {
		pair := [2]int{}
		for j, v := range strings.Split(pairStr, "-") {
			pair[j], _ = strconv.Atoi(v)
		}
		pairs[i] = pair
	}
	return pairs[0], pairs[1]
}

// small is fully contained in big
func fullyContained(big, small [2]int) bool {
	return small[0] >= big[0] && small[1] <= big[1]
}

func overlap(pair1, pair2 [2]int) bool {
	return (pair1[0] >= pair2[0] && pair1[0] <= pair2[1]) || (pair2[0] >= pair1[0] && pair2[0] <= pair1[1])
}
