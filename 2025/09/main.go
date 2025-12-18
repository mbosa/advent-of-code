package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type positionT struct {
	r int
	c int
}

type parsedInputT []positionT

type cornerT int

const (
	none cornerT = iota
	se           // south-east
	sw           // south-west
	ne           // north-east
	nw           // north-west
)

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	parsedInput := parseInput(input)

	res1 := part1(parsedInput)
	res2 := part2CoordinateCompression(parsedInput)
	// res2 := part2Edges(parsedInput)

	return res1, res2
}

func part1(parsedInput parsedInputT) int {
	maxArea := 0

	for i := 0; i < len(parsedInput)-1; i += 1 {
		a := parsedInput[i]
		for j := i + 1; j < len(parsedInput); j += 1 {
			b := parsedInput[j]

			area := calcArea(a, b)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

/*
- map the original coordinates to compressed coordinates
- draw the compressed polygon
- find the largest rectangle that, when compressed, fits inside the compressed polygon
*/
func part2CoordinateCompression(parsedInput parsedInputT) int {
	// map the original coordinates to compressed coordinates
	allR, allC := make([]int, 0, len(parsedInput)), make([]int, 0, len(parsedInput))

	for _, pos := range parsedInput {
		allR = append(allR, pos.r)
		allC = append(allC, pos.c)
	}

	slices.Sort(allR)
	slices.Sort(allC)

	ir, ic := 1, 1
	rMap := map[int]int{}
	cMap := map[int]int{}
	for _, r := range allR {
		if _, ok := rMap[r]; ok {
			continue
		}

		rMap[r] = ir
		ir += 2
	}
	for _, c := range allC {
		if _, ok := cMap[c]; ok {
			continue
		}

		cMap[c] = ic
		ic += 2
	}

	// draw the perimeter
	grid := make([][]byte, ir)
	for i := range grid {
		grid[i] = make([]byte, ic)
	}

	for i := 0; i < len(parsedInput); i += 1 {
		a := parsedInput[i]
		var b positionT
		if i == len(parsedInput)-1 {
			b = parsedInput[0]
		} else {
			b = parsedInput[i+1]
		}

		ar, br := rMap[a.r], rMap[b.r]
		ac, bc := cMap[a.c], cMap[b.c]

		grid[ar][ac] = '#'
		grid[br][bc] = '#'

		for n := min(ar, br) + 1; n < max(ar, br); n += 1 {
			grid[n][ac] = 'X'
		}
		for n := min(ac, bc) + 1; n < max(ac, bc); n += 1 {
			grid[ar][n] = 'X'
		}
	}

	// fill the polygon - ray casting
	for r, row := range grid {
		in := false
		prevCorner := none
		for c, b := range row {
			pos := positionT{r, c}

			if b == 'X' && prevCorner == none { // vertical edge
				in = !in
			} else if b == '#' { // corner
				if prevCorner == none {
					prevCorner = getCornerType(grid, pos)
					in = !in
				} else {
					if prevCorner == se && getCornerType(grid, pos) == sw {
						in = !in
					} else if prevCorner == ne && getCornerType(grid, pos) == nw {
						prevCorner = none
						in = !in
					}
					prevCorner = none
				}
			} else if b == 0 && !in { // outside point
				grid[r][c] = '.'
			}
		}
	}

	// find the valid rectangle with the largest area
	// a valid rectangle is one whose perimeter does not cross an outside point
	maxArea := 0
	for i := 0; i < len(parsedInput)-1; i += 1 {
		a := parsedInput[i]
		for j := i + 1; j < len(parsedInput); j += 1 {
			b := parsedInput[j]

			area := calcArea(a, b)

			if area <= maxArea {
				continue
			}

			ar, br := rMap[a.r], rMap[b.r]
			ac, bc := cMap[a.c], cMap[b.c]

			compressedA, compressedB := positionT{ar, ac}, positionT{br, bc}

			if !isRectangleInside(grid, compressedA, compressedB) {
				continue
			}

			maxArea = area
		}
	}

	return maxArea
}

type edgeT struct {
	a positionT
	b positionT
}

/*
- calculate all the edges
- find the largest rectangle that does not intersect or contain any edge
*/
func part2Edges(parsedInput parsedInputT) int {
	// collect all the edges
	edges := []edgeT{}

	for i := 0; i < len(parsedInput); i += 1 {
		a := parsedInput[i]
		var b positionT
		if i == len(parsedInput)-1 {
			b = parsedInput[0]
		} else {
			b = parsedInput[i+1]
		}
		edge := edgeT{a, b}
		edges = append(edges, edge)
	}

	// find the valid rectangle with the largest area
	// a valid rectangle is one that does not intersect or contain any edge (i.e. all edges are outside its perimeter)
	maxArea := 0
	for i := 0; i < len(parsedInput)-1; i += 1 {
		a := parsedInput[i]
		for j := i + 1; j < len(parsedInput); j += 1 {
			b := parsedInput[j]

			area := calcArea(a, b)

			if area <= maxArea {
				continue
			}

			if !isRectangleInsideEdges(edges, a, b) {
				continue
			}
			maxArea = area
		}
	}
	return maxArea
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		r := []string{}
		for _, b := range row {
			if b == 0 {
				r = append(r, "_")
			} else {
				r = append(r, string(b))
			}
		}
		fmt.Println(r)
	}
}

func parseInput(input string) parsedInputT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	parsed := make([]positionT, len(lines))

	for i, line := range lines {
		s := strings.Split(line, ",")
		c, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])

		pos := positionT{r, c}

		parsed[i] = pos
	}

	return parsed
}

func calcArea(a, b positionT) int {
	area := (max(a.r, b.r) - min(a.r, b.r) + 1) * (max(a.c, b.c) - min(a.c, b.c) + 1) // +1 to include the end row/col

	return area
}

func getCornerType(grid [][]byte, pos positionT) cornerT {
	r, c := pos.r, pos.c
	if r < len(grid)-1 && c < len(grid[0])-1 && grid[r+1][c] == 'X' && grid[r][c+1] == 'X' {
		return se
	}
	if r > 0 && c < len(grid[0])-1 && grid[r-1][c] == 'X' && grid[r][c+1] == 'X' {
		return ne
	}
	if r < len(grid)-1 && c > 0 && grid[r+1][c] == 'X' && grid[r][c-1] == 'X' {
		return sw
	}
	if r > 0 && c > 0 && grid[r-1][c] == 'X' && grid[r][c-1] == 'X' {
		return nw
	}

	return none
}

func isRectangleInside(grid [][]byte, a, b positionT) bool {
	minC, maxC := min(a.c, b.c), max(a.c, b.c)
	minR, maxR := min(a.r, b.r), max(a.r, b.r)

	for r := minR; r <= maxR; r += 1 {
		if grid[r][a.c] == '.' || grid[r][b.c] == '.' {
			return false
		}
	}
	for c := minC; c <= maxC; c += 1 {
		if grid[a.r][c] == '.' || grid[b.r][c] == '.' {
			return false
		}
	}
	return true
}

func isRectangleInsideEdges(edges []edgeT, a, b positionT) bool {
	minC, maxC := min(a.c, b.c), max(a.c, b.c)
	minR, maxR := min(a.r, b.r), max(a.r, b.r)

	for _, edge := range edges {
		isEdgeLeft := edge.a.c <= minC && edge.b.c <= minC
		if isEdgeLeft {
			continue
		}
		isEdgeRight := edge.a.c >= maxC && edge.b.c >= maxC
		if isEdgeRight {
			continue
		}
		isEdgeUp := edge.a.r <= minR && edge.b.r <= minR
		if isEdgeUp {
			continue
		}
		isEdgeDown := edge.a.r >= maxR && edge.b.r >= maxR
		if isEdgeDown {
			continue
		}
		return false
	}
	return true
}
