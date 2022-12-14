package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)
	input := parse(rawInput)

	occupiedPos, depth := calcPosWithRocks(input)

	startPos := Position{500, 0}
	floor := depth + 2
	count := 0
	prev := stack{startPos}

	for len(prev) > 0 {
		count++
		sand := prev.Peek()

		for {
			if resPart1 == 0 && sand.y == depth {
				// first sand over the edge
				resPart1 = count - 1
			}

			if sand.y == floor-1 {
				occupiedPos[sand] = true
				break
			}

			downPos := Position{x: sand.x, y: sand.y + 1}
			if _, ok := occupiedPos[downPos]; !ok {
				sand = downPos
				prev.Push(sand)
				continue
			}
			downLeftPos := Position{x: sand.x - 1, y: sand.y + 1}
			if _, ok := occupiedPos[downLeftPos]; !ok {
				sand = downLeftPos
				prev.Push(sand)
				continue
			}
			downRightPos := Position{x: sand.x + 1, y: sand.y + 1}
			if _, ok := occupiedPos[downRightPos]; !ok {
				sand = downRightPos
				prev.Push(sand)
				continue
			}

			break
		}
		occupiedPos[sand] = true
		prev.Pop()
	}

	return resPart1, count
}

type Position struct {
	x int
	y int
}

type stack []Position

func (s *stack) Push(item Position) {
	*s = append(*s, item)
}
func (s *stack) Pop() Position {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]

	return popped
}
func (s *stack) Peek() Position {
	n := len(*s)
	return (*s)[n-1]
}

func parse(raw []byte) [][]Position {
	parsed := [][]Position{}

	for _, line := range bytes.Split(raw, []byte{'\n'}) {
		inputLine := []Position{}
		split := bytes.Split(line, []byte(" -> "))

		for _, pair := range split {
			split := bytes.Split(pair, []byte{','})

			p := Position{x: bytesToInt(split[0]), y: bytesToInt(split[1])}
			inputLine = append(inputLine, p)
		}
		parsed = append(parsed, inputLine)
	}
	return parsed
}

func bytesToInt(b []byte) int {
	c, _ := strconv.Atoi(string(b))
	return c
}

func calcPosWithRocks(input [][]Position) (positions map[Position]bool, depth int) {
	positions = make(map[Position]bool)

	for _, line := range input {
		for i := 1; i < len(line); i++ {
			toPos := line[i]
			fromPos := line[i-1]

			if toPos.y > depth {
				depth = toPos.y
			} else if fromPos.y > depth {
				depth = fromPos.y
			}

			for i := min(toPos.x, fromPos.x); i <= max(toPos.x, fromPos.x); i++ {
				p := Position{x: i, y: toPos.y}
				positions[p] = true
			}
			for j := min(toPos.y, fromPos.y); j <= max(toPos.y, fromPos.y); j++ {
				p := Position{x: toPos.x, y: j}
				positions[p] = true
			}
		}
	}
	return positions, depth
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
