package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type parsedInputT [][]byte

type positionT struct {
	r int
	c int
}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	parsedInput := parseInput(input)

	res1 := part1(parsedInput)
	// res2 := part2Recursive(parsedInput)
	res2 := part2Iter(parsedInput)

	return res1, res2
}

func part1(parsedInput parsedInputT) int {
	res := 0

	startPos := positionT{}
	for r, row := range parsedInput {
		for c, el := range row {
			if el == 'S' {
				startPos = positionT{r, c}
			}
		}
	}

	seen := make([][]bool, len(parsedInput))
	for i := range seen {
		seen[i] = make([]bool, len(parsedInput[0]))
	}

	posStack := []positionT{startPos}

	for len(posStack) > 0 {
		pos := posStack[len(posStack)-1]
		posStack = posStack[:len(posStack)-1]

		if pos.r >= len(parsedInput) {
			continue
		}

		if seen[pos.r][pos.c] {
			continue
		}
		seen[pos.r][pos.c] = true

		if parsedInput[pos.r][pos.c] == '^' {
			// split beam
			res += 1

			posL := positionT{r: pos.r, c: pos.c - 1}
			posR := positionT{r: pos.r, c: pos.c + 1}

			posStack = append(posStack, posL, posR)
			continue
		}

		// go down
		nextPos := positionT{r: pos.r + 1, c: pos.c}
		posStack = append(posStack, nextPos)
	}

	return res
}

func part2Recursive(parsedInput parsedInputT) int {
	startPos := positionT{}
	for r, row := range parsedInput {
		for c, el := range row {
			if el == 'S' {
				startPos = positionT{r, c}
			}
		}
	}

	cache := map[positionT]int{}

	var helper func(pos positionT) int
	helper = func(pos positionT) int {
		if res, ok := cache[pos]; ok {
			return res
		}

		if pos.r >= len(parsedInput) {
			return 1
		}

		if parsedInput[pos.r][pos.c] == '^' {
			posL := positionT{r: pos.r, c: pos.c - 1}
			posR := positionT{r: pos.r, c: pos.c + 1}

			res := helper(posL) + helper(posR)
			cache[pos] = res
			return res
		} else {
			nextPos := positionT{r: pos.r + 1, c: pos.c}
			res := helper(nextPos)
			cache[pos] = res
			return res
		}
	}

	return helper(startPos)
}

func part2Iter(parsedInput parsedInputT) int {
	res := 0

	startPos := positionT{}
	for r, row := range parsedInput {
		for c, el := range row {
			if el == 'S' {
				startPos = positionT{r, c}
			}
		}
	}

	vals := make([][]int, len(parsedInput))
	for i := range vals {
		vals[i] = make([]int, len(parsedInput[0]))
	}

	vals[startPos.r][startPos.c] = 1

	for r := 0; r < len(vals)-1; r += 1 {
		row := vals[r]

		for c, v := range row {
			if v == 0 {
				continue
			}

			// move down
			if parsedInput[r+1][c] == '^' {
				vals[r+1][c-1] += v
				vals[r+1][c+1] += v
			} else {
				vals[r+1][c] += v
			}
		}
	}

	for _, v := range vals[len(vals)-1] {
		res += v
	}

	return res
}

func parseInput(input string) parsedInputT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	parsed := make([][]byte, len(lines))

	for i, l := range lines {
		parsed[i] = []byte(l)
	}

	return parsed
}
