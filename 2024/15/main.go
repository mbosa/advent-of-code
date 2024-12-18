package main

import (
	"fmt"
	"slices"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type parsedInputT struct {
	grid  [][]byte
	moves []byte
}

type position struct {
	r int
	c int
}

type directionT struct {
	dr int
	dc int
}

var (
	up    directionT = directionT{dr: -1, dc: 0}
	right            = directionT{dr: 0, dc: 1}
	down             = directionT{dr: 1, dc: 0}
	left             = directionT{dr: 0, dc: -1}
)

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

	// cloning the grid to not modify the input
	grid := cloneGrid(parsedInput.grid)
	moves := parsedInput.moves
	robotPos := findRobotPos(grid)

	for _, m := range moves {
		dir := moveToDirection(m)

		targets := make([]position, 0, 16)
		targets = append(targets, robotPos)
		moveTargets := true

		for i := 0; i < len(targets); i++ {
			t := targets[i]
			nextPos := position{r: t.r + dir.dr, c: t.c + dir.dc}
			b := grid[nextPos.r][nextPos.c]

			if b == '#' {
				moveTargets = false
				break
			} else if b == '.' {
				break
			} else {
				targets = append(targets, nextPos)
			}
		}

		if moveTargets {
			for i := len(targets) - 1; i > -1; i-- {
				t := targets[i]
				nextPos := position{r: t.r + dir.dr, c: t.c + dir.dc}

				grid[nextPos.r][nextPos.c], grid[t.r][t.c] = grid[t.r][t.c], grid[nextPos.r][nextPos.c]
			}

			robotPos.r += dir.dr
			robotPos.c += dir.dc
		}
	}

	for r, line := range grid {
		for c, b := range line {
			if b == 'O' {
				res += r*100 + c
			}
		}
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	firstWarehouseGrid, moves := parsedInput.grid, parsedInput.moves
	grid := createSecondWarehouseGrid(firstWarehouseGrid)
	robotPos := findRobotPos(grid)

	for _, m := range moves {
		dir := moveToDirection(m)

		targets := make([]position, 0, 16)
		targets = append(targets, robotPos)
		moveTargets := true

		for i := 0; i < len(targets); i++ {
			t := targets[i]

			nextPos := position{r: t.r + dir.dr, c: t.c + dir.dc}
			b := grid[nextPos.r][nextPos.c]

			if b == '#' {
				moveTargets = false
				break
			} else if b == '[' {
				rightSide := position{r: nextPos.r, c: nextPos.c + 1}

				if !slices.Contains[[]position](targets, nextPos) {
					targets = append(targets, nextPos)
				}
				if !slices.Contains[[]position](targets, rightSide) {
					targets = append(targets, rightSide)
				}
			} else if b == ']' {
				leftSide := position{r: nextPos.r, c: nextPos.c - 1}

				if !slices.Contains[[]position](targets, nextPos) {
					targets = append(targets, nextPos)
				}
				if !slices.Contains[[]position](targets, leftSide) {
					targets = append(targets, leftSide)
				}
			}
		}

		if moveTargets {
			for i := len(targets) - 1; i > -1; i-- {
				t := targets[i]
				nextPos := position{r: t.r + dir.dr, c: t.c + dir.dc}

				grid[nextPos.r][nextPos.c], grid[t.r][t.c] = grid[t.r][t.c], grid[nextPos.r][nextPos.c]
			}

			robotPos.r += dir.dr
			robotPos.c += dir.dc
		}
	}

	for r, line := range grid {
		for c, b := range line {
			if b == '[' {
				res += r*100 + c
			}
		}
	}

	return res
}

func parseInput(input string) parsedInputT {
	spl := strings.Split(strings.TrimSpace(input), "\n\n")

	gridStr, movesStr := spl[0], spl[1]

	gridLines := strings.Split(gridStr, "\n")

	grid := make([][]byte, len(gridLines))

	for i, lineStr := range gridLines {
		grid[i] = []byte(lineStr)
	}

	movesStr = strings.Join(strings.Split(movesStr, "\n"), "")

	moves := []byte(movesStr)

	return parsedInputT{grid, moves}
}

func cloneGrid(grid [][]byte) [][]byte {
	clone := make([][]byte, len(grid))
	for i := range clone {
		clone[i] = make([]byte, len(grid[0]))

		copy(clone[i], grid[i])
	}

	return clone
}

func findRobotPos(grid [][]byte) position {
	for r, line := range grid {
		for c, b := range line {
			if b == '@' {
				return position{r, c}
			}
		}
	}

	panic("Robot not found")
}

func moveToDirection(move byte) directionT {
	switch move {
	case '^':
		return up
	case '>':
		return right
	case 'v':
		return down
	case '<':
		return left
	}

	panic(fmt.Sprintf("moveToDirection: not a valid move: %q", move))
}

func createSecondWarehouseGrid(firstGrid [][]byte) [][]byte {
	secondGrid := make([][]byte, len(firstGrid))
	for r := range secondGrid {
		secondGrid[r] = make([]byte, 0, len(firstGrid[0])*2)
	}

	for r, line := range firstGrid {
		for _, b := range line {
			if b == '#' {
				secondGrid[r] = append(secondGrid[r], '#', '#')
			} else if b == 'O' {
				secondGrid[r] = append(secondGrid[r], '[', ']')
			} else if b == '.' {
				secondGrid[r] = append(secondGrid[r], '.', '.')
			} else if b == '@' {
				secondGrid[r] = append(secondGrid[r], '@', '.')
			}
		}
	}

	return secondGrid
}
