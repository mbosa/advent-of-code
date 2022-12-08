package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (score1 int, score2 int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		score1 += calcScore1(line)
		score2 += calcScore2(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	return score1, score2
}

func calcScore1(match string) int {
	lookup := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	return lookup[match]
}
func calcScore2(match string) int {
	lookup := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}
	return lookup[match]
}
