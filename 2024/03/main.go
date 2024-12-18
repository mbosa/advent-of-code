package main

import (
	"fmt"
	"regexp"
	"strconv"

	_ "embed"
)

//go:embed input.txt
var input string

var mulRegexPart1 = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var mulRegexPart2 = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	res1 := part1(input)
	res2 := part2(input)

	return res1, res2
}

func part1(input string) int {
	res := 0

	matches := mulRegexPart1.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		res += (n1 * n2)
	}

	return res
}

func part2(input string) int {
	res := 0

	matches := mulRegexPart2.FindAllStringSubmatch(input, -1)

	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else {
			if enabled {
				n1, _ := strconv.Atoi(match[1])
				n2, _ := strconv.Atoi(match[2])
				res += (n1 * n2)

			}
		}
	}

	return res
}
