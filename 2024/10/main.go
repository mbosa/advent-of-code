package main

import (
	"fmt"
	"slices"

	_ "embed"
)

//go:embed input.txt
var input []byte

type gridT [][]byte

type positionT struct {
	r int
	c int
}

var directions = []positionT{{r: -1, c: 0}, {r: 0, c: 1}, {r: 1, c: 0}, {r: 0, c: -1}}

type stackT struct{ items []positionT }

func (s *stackT) Push(pos positionT) {
	s.items = append(s.items, pos)
}
func (s *stackT) Pop() positionT {
	lastIndex := len(s.items) - 1
	el := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return el
}
func (s stackT) IsEmpty() bool {
	return len(s.items) == 0
}
func NewStack(cap int) stackT {
	return stackT{items: make([]positionT, 0, cap)}
}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input []byte) (int, int) {
	parsedInput := parseInput(input)

	res1, res2 := part1And2(parsedInput)

	return res1, res2
}

func part1And2(grid gridT) (int, int) {
	resPart1, resPart2 := 0, 0

	height := len(grid)
	width := len(grid[0])

	for r, row := range grid {
		for c, b := range row {
			if b != '0' {
				continue
			}

			stack := NewStack(8)
			stack.Push(positionT{r, c})

			score := map[positionT]struct{}{}

			for !stack.IsEmpty() {
				pos := stack.Pop()

				bb := grid[pos.r][pos.c]

				if bb == '9' {
					score[pos] = struct{}{}
					resPart2 += 1
				} else {
					for _, d := range directions {
						nextPos := positionT{r: pos.r + d.r, c: pos.c + d.c}

						if nextPos.r > -1 && nextPos.r < height && nextPos.c > -1 && nextPos.c < width {
							if grid[nextPos.r][nextPos.c]-bb == 1 {
								stack.Push(nextPos)
							}
						}
					}
				}
			}
			resPart1 += len(score)
		}
	}

	return resPart1, resPart2
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
