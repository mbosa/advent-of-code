package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1 int, resPart2 int) {
	max, max2, max3, acc := 0, 0, 0, 0

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			calories, _ := strconv.Atoi(line)
			acc += calories
			continue
		}

		if acc > max {
			max3, max2, max = max2, max, acc
		} else if acc > max2 {
			max3, max2 = max2, acc
		} else if acc > max3 {
			max3 = acc
		}
		acc = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	if acc > max {
		max3, max2, max = max2, max, acc
	} else if acc > max2 {
		max3, max2 = max2, acc
	} else if acc > max3 {
		max3 = acc
	}

	return max, max + max2 + max3
}
