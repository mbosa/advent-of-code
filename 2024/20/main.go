package main

import (
	"fmt"
	"slices"

	_ "embed"
)

//go:embed input.txt
var input []byte

type position struct {
	r int
	c int
}

type direction struct {
	dr int
	dc int
}

type cheat struct {
	start position
	end   position
}

var (
	up    direction = direction{dr: -1, dc: 0}
	right           = direction{dr: 0, dc: 1}
	down            = direction{dr: 1, dc: 0}
	left            = direction{dr: 0, dc: -1}
)

const targetTimeSave = 100

func main() {
	res1, res2 := solve(input, targetTimeSave)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input []byte, targetTimeSave int) (int, int) {
	grid := parseInput(input)

	rows, cols := len(grid), len(grid[0])

	directions := [4]direction{up, right, down, left}

	start, end := position{}, position{}
	for r, row := range grid {
		for c, b := range row {
			if b == 'S' {
				start.r, start.c = r, c
			} else if b == 'E' {
				end.r, end.c = r, c
			}
		}
	}

	// DFS
	stack := []position{start}
	path := []position{}

	seen := make([][]bool, rows)
	for r := range seen {
		seen[r] = make([]bool, cols)
	}

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		seen[pos.r][pos.c] = true
		path = append(path, pos)

		if pos == end {
			break
		}

		for _, d := range directions {
			nextPos := position{pos.r + d.dr, pos.c + d.dc}

			if grid[nextPos.r][nextPos.c] == '#' {
				continue
			}

			if seen[nextPos.r][nextPos.c] {
				continue
			}

			if nextPos == start {
				continue
			}

			stack = append(stack, nextPos)
		}
	}

	res1, res2 := make(chan int), make(chan int)

	go func() {
		defer close(res1)
		res1 <- cheatsToTarget(path, 2, targetTimeSave)
	}()

	go func() {
		defer close(res2)
		res2 <- cheatsToTarget(path, 20, targetTimeSave)
	}()

	return <-res1, <-res2
}

func cheatsToTarget(path []position, cheatTime, targetTimeSave int) int {
	res := 0

	for i := 0; i < len(path)-targetTimeSave; i++ {
		startCheat := path[i]

		for j := i + targetTimeSave; j < len(path); j++ {
			endCheat := path[j]

			dist := manhattanDistance(endCheat, startCheat)
			if dist <= cheatTime {
				timeSave := j - i - dist

				if timeSave >= targetTimeSave {
					res += 1
				}
			}
		}
	}
	return res
}

func parseInput(input []byte) [][]byte {
	l := len(input) - 1 // last byte can be a newline

	cols := slices.Index(input, '\n')
	rows := l / cols

	grid := make([][]byte, rows)

	for i, j := 0, 0; i < l; i, j = i+cols+1, j+1 {
		grid[j] = input[i : i+cols]
	}

	return grid
}

func manhattanDistance(p1, p2 position) int {
	dr := p1.r - p2.r
	dc := p1.c - p2.c
	if dr < 0 {
		dr = -dr
	}
	if dc < 0 {
		dc = -dc
	}
	return dr + dc
}
