package main

import (
	_ "embed"

	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type reportT []int
type reportsT [][]int

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

func parseInput(input string) reportsT {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	parsed := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		nums := make([]int, len(fields))

		for j, field := range fields {
			num, _ := strconv.Atoi(field)
			nums[j] = num
		}

		parsed[i] = nums
	}

	return parsed
}

func part1(parsedInput reportsT) int {
	res := 0

	for _, report := range parsedInput {
		if isReportValid(report, -1) {
			res += 1
		}
	}

	return res
}

func part2(parsedInput reportsT) int {
	res := 0

	for _, report := range parsedInput {
		if isReportValid(report, -1) {
			res += 1
		} else {
			for i := range report {
				if isReportValid(report, i) {
					res += 1
					break
				}
			}
		}
	}

	return res
}

func isReportValid(report reportT, skipIdx int) bool {
	firstIdx := 0
	secondsIdx := 1

	if skipIdx == 0 {
		firstIdx = 1
		secondsIdx = 2
	} else if skipIdx == 1 {
		secondsIdx = 2
	}

	increasing := false
	if report[secondsIdx] > report[firstIdx] {
		increasing = true
	}

	prev := report[firstIdx]

	for i := secondsIdx; i < len(report); i++ {
		if i == skipIdx {
			continue
		}

		curr := report[i]

		diff := max(curr, prev) - min(curr, prev)
		if diff < 1 || diff > 3 {
			return false
		}

		if increasing && curr <= prev {
			return false
		}
		if !increasing && curr >= prev {
			return false
		}

		prev = curr
	}

	return true
}
