package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

const inputFile = "input.txt"

var deltas = []Position{
	{x: 0, y: 1},
	{x: 1, y: 1},
	{x: 1, y: 0},
	{x: 1, y: -1},
	{x: 0, y: -1},
	{x: -1, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
}

var directions = []Position{
	{x: -1, y: 0},
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: 0, y: 1},
}
var adjMap = map[Position][]Position{
	{x: -1, y: 0}: {{x: -1, y: 1}, {x: -1, y: 0}, {x: -1, y: -1}},
	{x: 1, y: 0}:  {{x: 1, y: 1}, {x: 1, y: 0}, {x: 1, y: -1}},
	{x: 0, y: -1}: {{x: -1, y: -1}, {x: 0, y: -1}, {x: 1, y: -1}},
	{x: 0, y: 1}:  {{x: -1, y: 1}, {x: 0, y: 1}, {x: 1, y: 1}},
}

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)
	elves := parse(rawInput)

	round := 0
	// part1
	for round < 10 {
		elves.DoRound(round)
		round++
	}
	minBottomLeft, topRight := elves.calcBoundaries()
	resPart1 = (topRight.x-minBottomLeft.x+1)*(topRight.y-minBottomLeft.y+1) - len(elves)

	// part2
	someoneMoved := true
	for someoneMoved {
		someoneMoved = elves.DoRound(round)
		round++
	}
	resPart2 = round

	return resPart1, resPart2
}

type Position struct {
	x int
	y int
}

type Elves map[Position]bool

func (e *Elves) DoRound(round int) (someoneMoved bool) {
	proposedPos := map[Position][]Position{}

	for elf := range *e {
		if !e.canElfMove(elf) {
			continue
		}

		dirs := e.calcDirectionOrderByRound(round)
		for _, dir := range dirs {
			if e.canElfMoveInDirection(elf, dir) {
				nextPos := Position{x: elf.x + dir.x, y: elf.y + dir.y}
				proposedPos[nextPos] = append(proposedPos[nextPos], elf)
				break
			}
		}
	}

	return e.resolveProposedPositions(proposedPos)
}

func (e *Elves) canElfMove(pos Position) bool {
	for _, d := range deltas {
		p := Position{x: pos.x + d.x, y: pos.y + d.y}
		if _, ok := (*e)[p]; ok {
			return true
		}
	}
	return false
}

func (e *Elves) canElfMoveInDirection(pos, dir Position) bool {
	adjs := adjMap[dir]
	for _, adj := range adjs {
		p := Position{x: pos.x + adj.x, y: pos.y + adj.y}
		if _, ok := (*e)[p]; ok {
			return false
		}
	}
	return true
}

func (e *Elves) moveElf(src, dst Position) {
	delete(*e, src)
	(*e)[dst] = true
}

func (e *Elves) resolveProposedPositions(proposedPos map[Position][]Position) (someoneMoved bool) {
	for proposed, elves := range proposedPos {
		if len(elves) == 1 {
			e.moveElf(elves[0], proposed)
			someoneMoved = true
		}
	}
	return someoneMoved
}

func (e *Elves) calcDirectionOrderByRound(round int) []Position {
	offset := round % len(directions)
	return append(directions[offset:], directions[:offset]...)
}

func (e *Elves) calcBoundaries() (bottomLeft, topRight Position) {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for elf := range *e {
		minX = min(minX, elf.x)
		maxX = max(maxX, elf.x)
		minY = min(minY, elf.y)
		maxY = max(maxY, elf.y)
	}
	min, max := Position{x: minX, y: minY}, Position{x: maxX, y: maxY}

	return min, max
}

func parse(raw []byte) Elves {
	res := Elves{}

	for i, line := range bytes.Split(raw, []byte{'\n'}) {
		for j, el := range line {
			if el == '#' {
				p := Position{x: i, y: j}
				res[p] = true
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
