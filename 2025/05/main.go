package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type rangeT struct {
	start int
	end   int
}

type parsedInputT struct {
	ranges      []rangeT
	ingredients []int
}

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

ingLoop:
	for _, ingredient := range parsedInput.ingredients {
		for _, r := range parsedInput.ranges {
			if ingredient >= r.start && ingredient <= r.end {
				res += 1
				continue ingLoop
			}
		}
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	ranges := parsedInput.ranges

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	maxEnd := 0
	for i := 0; i < len(ranges); i += 1 {
		r := ranges[i]

		if r.end <= maxEnd {
			continue
		}

		res += (r.end - max(maxEnd+1, r.start) + 1)

		maxEnd = r.end
	}

	return res
}

func parseInput(input string) parsedInputT {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")

	rangesStr := strings.Split(sections[0], "\n")
	ranges := make([]rangeT, len(rangesStr))

	ingredientsStr := strings.Split(sections[1], "\n")
	ingredients := make([]int, len(ingredientsStr))

	for i, rangeStr := range rangesStr {
		fields := strings.Split(rangeStr, "-")

		s, _ := strconv.Atoi(fields[0])
		e, _ := strconv.Atoi(fields[1])

		r := rangeT{start: s, end: e}
		ranges[i] = r
	}

	for i, ingretientStr := range ingredientsStr {
		ing, _ := strconv.Atoi(ingretientStr)
		ingredients[i] = ing
	}

	return parsedInputT{ranges, ingredients}
}
