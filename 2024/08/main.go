package main

import (
	_ "embed"

	"fmt"
	"slices"
)

//go:embed input.txt
var input []byte

type positionT struct {
	r int
	c int
}
type antennaT struct {
	pos  positionT
	freq byte
}

type gridT [][]byte

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

func part1(grid gridT) int {
	return findAndCountAntinodes(grid, false)
}

func part2(grid gridT) int {
	return findAndCountAntinodes(grid, true)
}

func parseInput(input []byte) gridT {
	l := len(input) - 1 // last byte can be a newline

	width := slices.Index(input, '\n')
	height := l / width

	parsed := make([][]byte, height)

	for i, j := 0, 0; i < l; i, j = i+width+1, j+1 {
		parsed[j] = input[i : i+width]
	}

	return parsed
}

func makeGrid(rows, cols int) gridT {
	grid := make([][]byte, rows)

	for i := range rows {
		grid[i] = make([]byte, cols)
	}
	return grid
}

func findAntinodes(grid gridT, part2 bool) gridT {
	rows := len(grid)
	cols := len(grid[0])

	antennas := findAntennas(grid)
	antinodes := makeGrid(rows, cols)

	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			antennaA := antennas[i]
			antennaB := antennas[j]

			if antennaA.freq != antennaB.freq {
				continue
			}

			dr := antennaB.pos.r - antennaA.pos.r
			dc := antennaB.pos.c - antennaA.pos.c

			if !part2 {
				nextPos := positionT{r: antennaB.pos.r + dr, c: antennaB.pos.c + dc}
				prevPos := positionT{r: antennaA.pos.r - dr, c: antennaA.pos.c - dc}

				if nextPos.r > -1 && nextPos.r < rows && nextPos.c > -1 && nextPos.c < cols {
					antinodes[nextPos.r][nextPos.c] = '#'
				}
				if prevPos.r > -1 && prevPos.r < rows && prevPos.c > -1 && prevPos.c < cols {
					antinodes[prevPos.r][prevPos.c] = '#'
				}
			} else {
				// each antenna is also an antinode
				nextPos := positionT{r: antennaB.pos.r, c: antennaB.pos.c}
				prevPos := positionT{r: antennaA.pos.r, c: antennaA.pos.c}

				for nextPos.r > -1 && nextPos.r < rows && nextPos.c > -1 && nextPos.c < cols {
					antinodes[nextPos.r][nextPos.c] = '#'

					nextPos.r += dr
					nextPos.c += dc
				}
				for prevPos.r > -1 && prevPos.r < rows && prevPos.c > -1 && prevPos.c < cols {
					antinodes[prevPos.r][prevPos.c] = '#'

					prevPos.r -= dr
					prevPos.c -= dc
				}
			}
		}
	}

	return antinodes
}

func findAntennas(grid gridT) []antennaT {
	var antennas []antennaT

	for r, row := range grid {
		for c, b := range row {
			if b != '.' {
				antenna := antennaT{
					pos:  positionT{r, c},
					freq: b,
				}
				antennas = append(antennas, antenna)
			}
		}
	}
	return antennas
}

func countAntinodes(grid gridT) int {
	res := 0

	for _, row := range grid {
		for _, b := range row {
			if b == '#' {
				res += 1
			}
		}
	}
	return res
}

func findAndCountAntinodes(grid gridT, part2 bool) int {
	antinodes := findAntinodes(grid, part2)
	return countAntinodes(antinodes)
}
