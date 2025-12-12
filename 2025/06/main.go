package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type problemsT struct {
	nums       [][]int
	operations []string
}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	parsedInput := parseInput(input)
	parsedInputPart2 := parseInputPart2(input)

	res1 := part1(parsedInput)
	res2 := part2(parsedInputPart2)

	return res1, res2
}

func part1(parsedInput problemsT) int {
	return solveProblems(parsedInput)
}

func part2(parsedInput problemsT) int {
	return solveProblems(parsedInput)
}

func parseInput(input string) problemsT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	operations := strings.Fields(lines[len(lines)-1])

	nums := make([][]int, len(operations))
	for i := range nums {
		nums[i] = make([]int, 0)
	}

	for i := 0; i < len(lines)-1; i += 1 {
		line := lines[i]
		fields := strings.Fields(line)

		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			nums[i] = append(nums[i], n)
		}
	}

	return problemsT{nums, operations}
}

func parseInputPart2(input string) problemsT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	operations := strings.Fields(lines[len(lines)-1])

	numLines := lines[:len(lines)-1]
	nums := make([][]int, len(operations))
	for i := range nums {
		nums[i] = make([]int, 0)
	}

	q := len(operations) - 1

	for i := len(numLines[0]) - 1; i >= 0; i -= 1 {
		numBytes := []byte{}

		for _, l := range numLines {
			c := l[i]

			if c != ' ' {
				numBytes = append(numBytes, c)
			}
		}

		if len(numBytes) > 0 {
			n, _ := strconv.Atoi(string(numBytes))
			nums[q] = append(nums[q], n)

		} else {
			q -= 1
		}

	}

	return problemsT{nums, operations}
}

func solveProblems(problems problemsT) int {
	res := 0

	for i, op := range problems.operations {
		nn := problems.nums[i]

		partialRes := 0
		if op == "*" {
			partialRes = 1
		}

		for _, n := range nn {
			switch op {
			case "+":
				partialRes += n
			case "*":
				partialRes *= n
			}

		}

		res += partialRes
	}

	return res
}
