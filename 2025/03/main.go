package main

import (
	"fmt"
	"math"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type parsedInputT []string

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	parsedInput := parseInput(input)

	res1 := part1(parsedInput)
	res2 := part2(parsedInput)

	return res1, res2
}

func part1(parsedInput parsedInputT) int {
	res := 0

	for _, bank := range parsedInput {
		topOrderedN := getTopOrderedN(bank, 2)
		res += digitsArrToInt(topOrderedN)
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0
	for _, bank := range parsedInput {
		topOrderedN := getTopOrderedN(bank, 12)
		res += digitsArrToInt(topOrderedN)
	}
	return res
}

func getTopOrderedN(str string, n int) []int {
	res := make([]int, n)

	lastIdx := -1

	for i := range n {
		top := 0
		for j := lastIdx + 1; j < len(str)-(n-i-1); j += 1 {
			num := int(str[j] - '0')
			if num > top {
				top = num
				lastIdx = j
			}
		}
		res[i] = top
	}
	return res
}

func digitsArrToInt(digits []int) int {
	res := 0

	for i, d := range digits {
		res += (d * int(math.Pow(10, float64(len(digits)-i-1))))
	}

	return res
}

func parseInput(input string) parsedInputT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	return lines
}
