package main

import (
	_ "embed"

	"fmt"
	"slices"
)

//go:embed input.txt
var input []byte

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input []byte) (int, int) {
	parsedInput := parseInput(input)

	res1 := part1(parsedInput)
	res2 := part2(parsedInput)

	return res1, res2
}

func part1(parsedInput [][]byte) int {
	res := 0

	directions := [8][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	height := len(parsedInput)
	width := len(parsedInput[0])

	for i, row := range parsedInput {
		for j, el := range row {
			if el != 'X' {
				continue
			}

			for _, dir := range directions {
				dr, dc := dir[0], dir[1]

				r1, c1 := i+dr, j+dc
				r2, c2 := i+2*dr, j+2*dc
				r3, c3 := i+3*dr, j+3*dc

				if r3 < 0 || r3 >= height || c3 < 0 || c3 >= width {
					continue
				}

				if parsedInput[r1][c1] == 'M' && parsedInput[r2][c2] == 'A' && parsedInput[r3][c3] == 'S' {
					res += 1
				}
			}
		}
	}

	return res
}

func part2(parsedInput [][]byte) int {
	res := 0

	valid := [4][4]byte{{'M', 'M', 'S', 'S'}, {'M', 'S', 'S', 'M'}, {'S', 'S', 'M', 'M'}, {'S', 'M', 'M', 'S'}}

	height := len(parsedInput)
	width := len(parsedInput[0])

	for i := 1; i < height-1; i++ {
		row := parsedInput[i]

		for j := 1; j < width-1; j++ {
			el := row[j]

			if el != 'A' {
				continue
			}

			topLeft := parsedInput[i-1][j-1]
			topRight := parsedInput[i-1][j+1]
			bottomRight := parsedInput[i+1][j+1]
			bottomLeft := parsedInput[i+1][j-1]

			round := [4]byte{topLeft, topRight, bottomRight, bottomLeft}

			for _, v := range valid {
				if v == round {
					res += 1
					break
				}
			}
		}
	}
	return res
}

func parseInput(input []byte) [][]byte {
	l := len(input) - 1 // last byte can be a newline

	width := slices.Index(input, '\n')
	height := l / width

	parsed := make([][]byte, height)

	for i, j := 0, 0; i < l; i, j = i+width+1, j+1 {
		parsed[j] = input[i : i+width]
	}

	return parsed
}
