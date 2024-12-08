package main

import (
	_ "embed"
	"slices"
	"sync"
	"sync/atomic"

	"fmt"
)

//go:embed input.txt
var input []byte

type gridT [][]byte
type positionT struct {
	r int
	c int
}
type directionT struct {
	r int
	c int
}

var (
	up    directionT = directionT{r: -1, c: 0}
	right            = directionT{r: 0, c: 1}
	down             = directionT{r: 1, c: 0}
	left             = directionT{r: 0, c: -1}
)

var directionByteMap = map[directionT]byte{up: '^', right: '>', down: 'v', left: '<'}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input []byte) (int, int) {
	grid := parseInput(input)

	res1 := part1(grid)
	res2 := part2(grid)

	return res1, res2
}

// calculate part1 result. It modifies the input by painting the path with 'x', needed in part2
func part1(grid gridT) int {
	res := 0

	rows, cols := len(grid), len(grid[0])

	startPosition := positionT{}
	for r, row := range grid {
		for c, b := range row {
			if b == '^' {
				startPosition.r, startPosition.c = r, c
			}
		}
	}

	direction := up
	position := startPosition

	for {
		nextR := position.r + direction.r
		nextC := position.c + direction.c

		if nextR < 0 || nextR > rows-1 {
			grid[position.r][position.c] = 'x'
			break
		}
		if nextC < 0 || nextC > cols-1 {
			grid[position.r][position.c] = 'x'
			break
		}

		nextB := grid[nextR][nextC]

		if nextB == '#' {
			direction = rotatedDir(direction)
		} else {
			grid[position.r][position.c] = 'x'
			position.r, position.c = nextR, nextC
		}
	}

	for _, row := range grid {
		for _, b := range row {
			if b == 'x' {
				res += 1
			}
		}
	}

	grid[startPosition.r][startPosition.c] = '^'

	return res
}

func part2(grid gridT) int {
	var q atomic.Int64
	var wg sync.WaitGroup

	startPosition := positionT{}
	for r, row := range grid {
		for c, b := range row {
			if b == '^' {
				startPosition.r, startPosition.c = r, c
			}
		}
	}

	for r, row := range grid {
		for c, b := range row {
			if b != 'x' {
				continue
			}

			wg.Add(1)
			go func() {
				g := clone(grid)
				g[r][c] = '#'
				if loops(g, startPosition) {
					q.Add(1)
				}
				wg.Done()
			}()
		}
	}

	wg.Wait()

	return int(q.Load())
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

func clone(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))

	for i, row := range src {
		dst[i] = make([]byte, len(row))

		for j, b := range row {
			dst[i][j] = b
		}
	}
	return dst
}

func rotatedDir(dir directionT) directionT {
	return directionT{r: dir.c, c: dir.r * -1}
}

// return true if there is a loop. It modifies the input grid
func loops(grid gridT, startPosition positionT) bool {
	rows, cols := len(grid), len(grid[0])
	direction := up
	position := startPosition

	for {
		if position != startPosition && grid[position.r][position.c] == directionByteMap[direction] {
			return true
		}

		nextR := position.r + direction.r
		nextC := position.c + direction.c

		if nextR < 0 || nextR > rows-1 {
			break
		}
		if nextC < 0 || nextC > cols-1 {
			break
		}

		nextB := grid[nextR][nextC]

		if nextB == '#' {
			direction = rotatedDir(direction)
		} else {
			grid[position.r][position.c] = directionByteMap[direction]
			position.r, position.c = nextR, nextC
		}
	}

	return false
}
