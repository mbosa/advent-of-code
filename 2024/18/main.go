package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type position struct {
	r int
	c int
}

type direction struct {
	dr int
	dc int
}

type qItem struct {
	pos  position
	dist int
}

var (
	up    direction = direction{dr: -1, dc: 0}
	right           = direction{dr: 0, dc: 1}
	down            = direction{dr: 1, dc: 0}
	left            = direction{dr: 0, dc: -1}
)

const rows = 71
const cols = 71
const part1Bytes = 1024

func main() {
	res1, res2 := solve(input, rows, cols, part1Bytes)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string, rows, cols, part1Bytes int) (int, string) {
	bytes := parseInput(input)

	res1 := part1(bytes, rows, cols, part1Bytes)
	res2 := part2(bytes, rows, cols, part1Bytes)

	return res1, res2
}

func part1(bytes []position, rows, cols, part1Bytes int) int {
	grid := buildMemoryGrid(bytes, rows, cols, part1Bytes)

	start, end := position{0, 0}, position{rows - 1, cols - 1}

	return bfs(grid, start, end, part1Bytes)
}

func part2(bytes []position, rows, cols, part1Bytes int) string {
	grid := buildMemoryGrid(bytes, rows, cols, len(bytes))

	start, end := position{0, 0}, position{rows - 1, cols - 1}

	// Binary search to find the first byte that prevents exit
	// If exit is possible, the target byte will be after mid
	// If exit is not possible, the target byte will be before mid
	low, high := part1Bytes, len(bytes)-1

	for low <= high {
		mid := (low + high) / 2

		if bfs(grid, start, end, mid) != -1 {
			// end reached
			low = mid + 1
		} else {
			// end not reached
			high = mid - 1
		}
	}

	return fmt.Sprintf("%d,%d", bytes[low].c, bytes[low].r)
}

func parseInput(input string) []position {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	bytes := make([]position, len(lines))

	for i, line := range lines {
		spl := strings.Split(line, ",")

		p := position{r: mustAtoi(spl[1]), c: mustAtoi(spl[0])}

		bytes[i] = p
	}

	return bytes
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(s + " cannot be converted to int")
	}

	return n
}

func buildMemoryGrid(bytes []position, rows, cols, time int) [][]int {
	grid := make([][]int, rows)
	for r := range grid {
		grid[r] = make([]int, cols)
		for c := range grid[r] {
			grid[r][c] = math.MaxInt
		}
	}

	for i := range time {
		b := bytes[i]
		grid[b.r][b.c] = i
	}

	return grid
}

func makeSeenGrid(rows, cols int) [][]bool {
	g := make([][]bool, rows)
	for i := range g {
		g[i] = make([]bool, cols)
	}
	return g
}

// Breadth first search
// Each position in the grid contains the time when a byte falls on it
// When traversing, all positions with n > current time are accessible,
// since a byte didn't fall on it yet
func bfs(grid [][]int, start, end position, time int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [4]direction{up, right, down, left}

	startItem := qItem{pos: start, dist: 0}
	q := []qItem{startItem}

	seen := makeSeenGrid(rows, cols)

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		if item.pos == end {
			return item.dist
		}

		for _, d := range directions {
			nextPos := position{r: item.pos.r + d.dr, c: item.pos.c + d.dc}

			if nextPos.r < 0 || nextPos.r >= rows || nextPos.c < 0 || nextPos.c >= cols {
				continue
			}
			if !seen[nextPos.r][nextPos.c] && time < grid[nextPos.r][nextPos.c] {
				seen[nextPos.r][nextPos.c] = true
				nextItem := qItem{pos: nextPos, dist: item.dist + 1}
				q = append(q, nextItem)
			}
		}

	}
	return -1
}
