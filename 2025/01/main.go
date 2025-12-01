package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type Direction int

type Rotation struct {
	dir Direction
	val int
}

const (
	Left  Direction = -1
	Right Direction = 1
)

type parsedInputT []Rotation

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

	dialVal := 50

	for _, rot := range parsedInput {
		dialVal = dialVal + int(rot.dir)*rot.val

		if dialVal%100 == 0 {
			res += 1
		}
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	dialVal := 50

	for _, rot := range parsedInput {
		res += rot.val / 100

		if rot.dir == Right {
			res += (dialVal + rot.val%100) / 100
		} else if rot.val%100 >= dialVal && dialVal != 0 {
			res += 1
		}

		dialVal = (100 + dialVal + int(rot.dir)*rot.val%100) % 100
	}

	return res
}

func parseInput(input string) parsedInputT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	parsed := make([]Rotation, len(lines))

	for i, line := range lines {
		dirStr := line[0]
		dir := Left
		if dirStr == 'R' {
			dir = Right
		}

		valStr := line[1:]
		val, _ := strconv.Atoi(valStr)

		rot := Rotation{dir, val}

		parsed[i] = rot
	}
	return parsed
}
