package main

import (
	"fmt"
	"slices"

	_ "embed"
)

//go:embed input.txt
var input []byte

type gridT [][]byte

type position struct {
	r int
	c int
}

type stack struct{ items []position }

func (s *stack) Push(pos position) {
	s.items = append(s.items, pos)
}
func (s *stack) Pop() position {
	lastIndex := len(s.items) - 1
	el := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return el
}
func (s stack) IsEmpty() bool {
	return len(s.items) == 0
}
func newStack(cap int) stack {
	return stack{items: make([]position, 0, cap)}
}

var (
	up    position = position{r: -1, c: 0}
	right          = position{r: 0, c: 1}
	down           = position{r: 1, c: 0}
	left           = position{r: 0, c: -1}
)

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
	res := 0

	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for r, row := range grid {
		for c, b := range row {
			pos := position{r, c}

			if visited[pos.r][pos.c] {
				continue
			}

			area := 0
			perimeter := 0

			st := newStack(0)
			st.Push(pos)

			for !st.IsEmpty() {
				currPos := st.Pop()

				if currPos.r < 0 || currPos.r >= rows || currPos.c < 0 || currPos.c >= cols {
					perimeter += 1
					continue
				}

				if grid[currPos.r][currPos.c] != b {
					perimeter += 1
					continue
				}

				if visited[currPos.r][currPos.c] {
					continue
				}

				visited[currPos.r][currPos.c] = true

				area += 1

				posUp := position{r: currPos.r + up.r, c: currPos.c + up.c}
				posRight := position{r: currPos.r + right.r, c: currPos.c + right.c}
				posDown := position{r: currPos.r + down.r, c: currPos.c + down.c}
				posLeft := position{r: currPos.r + left.r, c: currPos.c + left.c}

				st.Push(posUp)
				st.Push(posRight)
				st.Push(posDown)
				st.Push(posLeft)
			}

			price := area * perimeter
			// fmt.Printf("%s - area: %d; perimeter: %d; price: %d\n", string(b), area, perimeter, price)
			res += price
		}
	}

	return res
}

type sideRow struct {
	r         int
	direction position
}
type sideCol struct {
	c         int
	direction position
}

func part2(grid gridT) int {
	res := 0

	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for r, row := range grid {
		for c, b := range row {
			pos := position{r, c}

			if visited[pos.r][pos.c] {
				continue
			}

			area := 0
			shape := make([][]byte, rows+2)
			for i := range shape {
				shape[i] = make([]byte, cols+2)

				for j := range shape[i] {
					shape[i][j] = '.'
				}
			}
			sidesSetRow := map[sideRow]struct{}{}
			sidesSetCol := map[sideCol]struct{}{}

			st := newStack(0)
			st.Push(pos)

			for !st.IsEmpty() {
				currPos := st.Pop()

				if visited[currPos.r][currPos.c] {
					continue
				}

				visited[currPos.r][currPos.c] = true
				shape[currPos.r+1][currPos.c+1] = '#'

				area += 1

				posUp := position{r: currPos.r + up.r, c: currPos.c + up.c}
				posRight := position{r: currPos.r + right.r, c: currPos.c + right.c}
				posDown := position{r: currPos.r + down.r, c: currPos.c + down.c}
				posLeft := position{r: currPos.r + left.r, c: currPos.c + left.c}

				if posUp.r < 0 || posUp.r >= rows || posUp.c < 0 || posUp.c >= cols {
					side := sideRow{
						r:         currPos.r,
						direction: up,
					}
					sidesSetRow[side] = struct{}{}
				} else {
					if grid[posUp.r][posUp.c] != b {
						side := sideRow{
							r:         currPos.r,
							direction: up,
						}
						sidesSetRow[side] = struct{}{}
					} else {
						st.Push(posUp)
					}
				}

				if posRight.r < 0 || posRight.r >= rows || posRight.c < 0 || posRight.c >= cols {
					side := sideCol{
						c:         currPos.c,
						direction: right,
					}
					sidesSetCol[side] = struct{}{}
				} else {
					if grid[posRight.r][posRight.c] != b {
						side := sideCol{
							c:         currPos.c,
							direction: right,
						}
						sidesSetCol[side] = struct{}{}
					} else {
						st.Push(posRight)
					}
				}

				if posDown.r < 0 || posDown.r >= rows || posDown.c < 0 || posDown.c >= cols {
					side := sideRow{
						r:         currPos.r,
						direction: down,
					}
					sidesSetRow[side] = struct{}{}
				} else {
					if grid[posDown.r][posDown.c] != b {
						side := sideRow{
							r:         currPos.r,
							direction: down,
						}
						sidesSetRow[side] = struct{}{}
					} else {
						st.Push(posDown)
					}
				}

				if posLeft.r < 0 || posLeft.r >= rows || posLeft.c < 0 || posLeft.c >= cols {
					side := sideCol{
						c:         currPos.c,
						direction: left,
					}
					sidesSetCol[side] = struct{}{}
				} else {
					if grid[posLeft.r][posLeft.c] != b {
						side := sideCol{
							c:         currPos.c,
							direction: left,
						}
						sidesSetCol[side] = struct{}{}
					} else {
						st.Push(posLeft)
					}
				}

			}

			sides := 0
			for r := 1; r < len(shape)-1; r++ {
				for c := 1; c < len(shape[0])-1; c++ {
					b := shape[r][c]

					if b == '.' {
						continue
					}

					// NW
					if shape[r-1][c] != b && shape[r][c-1] != b {
						// convex
						sides += 1
					} else if shape[r-1][c-1] != b && shape[r-1][c] == b && shape[r][c-1] == b {
						// concave
						sides += 1
					}
					// NE
					if shape[r-1][c] != b && shape[r][c+1] != b {
						// convex
						sides += 1
					} else if shape[r-1][c+1] != b && shape[r-1][c] == b && shape[r][c+1] == b {
						// concave
						sides += 1
					}
					// SE
					if shape[r+1][c] != b && shape[r][c+1] != b {
						// convex
						sides += 1
					} else if shape[r+1][c+1] != b && shape[r+1][c] == b && shape[r][c+1] == b {
						// concave
						sides += 1
					}
					// SW
					if shape[r+1][c] != b && shape[r][c-1] != b {
						// convex
						sides += 1
					} else if shape[r+1][c-1] != b && shape[r+1][c] == b && shape[r][c-1] == b {
						// concave
						sides += 1
					}
				}
			}

			price := area * sides
			// fmt.Printf("%s - area: %d; sides: %d; price: %d\n", string(b), area, sides, price)
			res += price
		}
	}

	return res
}

func parseInput(input []byte) gridT {
	l := len(input) - 1 // last byte can be a newline

	width := slices.Index(input, '\n')
	height := l / width

	parsed := make(gridT, height)

	for i, j := 0, 0; i < l; i, j = i+width+1, j+1 {
		parsed[j] = input[i : i+width]
	}

	return parsed
}
