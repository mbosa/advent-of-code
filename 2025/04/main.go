package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type parsedInputT [][]byte

var adj = [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

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

	for r := 1; r < len(parsedInput)-1; r += 1 {
		for c := 1; c < len(parsedInput[0])-1; c += 1 {
			if parsedInput[r][c] != '@' {
				continue
			}

			adjRolls := 0

			for _, a := range adj {
				if parsedInput[r+a[0]][c+a[1]] == '@' {
					adjRolls += 1
				}
			}
			if adjRolls < 4 {
				res += 1
			}
		}
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	for {
		found := false

		for r := 1; r < len(parsedInput)-1; r += 1 {
			for c := 1; c < len(parsedInput[0])-1; c += 1 {
				if parsedInput[r][c] != '@' {
					continue
				}

				adjRolls := 0

				for _, a := range adj {
					if parsedInput[r+a[0]][c+a[1]] == '@' {
						adjRolls += 1
					}
				}
				if adjRolls < 4 {
					found = true
					res += 1
					parsedInput[r][c] = '.'
				}
			}
		}

		if !found {
			break
		}
	}

	return res
}

func parseInput(input string) parsedInputT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	parsed := make([][]byte, len(lines))

	for i, line := range lines {
		parsed[i] = append(append([]byte{'.'}, []byte(line)...), '.')
	}

	emptyLine := make([]byte, len(lines[0])+2)
	for i := range emptyLine {
		emptyLine[i] = '.'
	}

	parsed = append(append([][]byte{emptyLine}, parsed...), emptyLine)

	return parsed
}
