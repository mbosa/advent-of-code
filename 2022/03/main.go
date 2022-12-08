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

func solve(inputFile string) (resPart1, resPart2 int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	common := []byte{}
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		n := len(line)

		// part 1
		duplicateInBag := commonChars(line[:n/2], line[n/2:])
		resPart1 += calcPriority(duplicateInBag[0])

		// part 2
		if i%3 == 0 { // first of the group of 3
			common = line
		} else {
			common = commonChars(line, common)
		}
		if i%3 == 2 { // last of the group of 3
			resPart2 += calcPriority(common[0])
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	return resPart1, resPart2
}

func charToValue(char byte) int {
	if char >= 'a' {
		return int(char - 'a')
	} else {
		return int(char-'A') + 26
	}
}

func calcPriority(char byte) int {
	return charToValue(char) + 1
}

func bitmask(b []byte) uint64 {
	var mask uint64
	for _, char := range b {
		mask |= 1 << charToValue(char)
	}
	return mask
}

func commonChars(chars1, chars2 []byte) []byte {
	res := []byte{}

	mask1 := bitmask(chars1)

	for _, char := range chars2 {
		if mask1&(1<<charToValue(char)) > 0 {
			res = append(res, char)
			mask1 &= ^(1 << charToValue(char))
		}
	}

	return res
}
