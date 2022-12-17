package main

import (
	"bytes"
	"fmt"
	"os"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)
	input := bytes.Split(rawInput, []byte{'\n'})

	var start, end Position

	for i, row := range input {
		for j, el := range row {
			if el == 'S' {
				start = Position{col: j, row: i}
			} else if el == 'E' {
				end = Position{col: j, row: i}
			}
		}
	}
	input[start.row][start.col] = 'a'
	input[end.row][end.col] = 'z'

	// bfs from end to start, keep track of the first 'a' encountered
	var shortestA *Position

	visited := make([][]int, len(input))
	for i := range visited {
		visited[i] = make([]int, len(input[0]))
	}
	visited[end.row][end.col] = 0

	q := queue{}
	q.push(end)

	for len(q) > 0 {
		current := q.pop()

		if shortestA == nil && input[current.row][current.col] == 'a' {
			shortestA = &current
		}

		neighbors := []Position{}
		if current.row > 0 {
			neighbors = append(neighbors, Position{col: current.col, row: current.row - 1})
		}
		if current.row < len(input)-1 {
			neighbors = append(neighbors, Position{col: current.col, row: current.row + 1})
		}
		if current.col > 0 {
			neighbors = append(neighbors, Position{col: current.col - 1, row: current.row})
		}
		if current.col < len(input[0])-1 {
			neighbors = append(neighbors, Position{col: current.col + 1, row: current.row})
		}

		for _, next := range neighbors {
			if int(input[current.row][current.col])-int(input[next.row][next.col]) > 1 {
				continue
			}

			costToNext := visited[current.row][current.col] + 1

			if next != end && (visited[next.row][next.col] == 0 || costToNext < visited[next.row][next.col]) {
				visited[next.row][next.col] = costToNext

				q.push(next)
			}
		}
	}

	return visited[start.row][start.col], visited[shortestA.row][shortestA.col]
}

type Position struct {
	row int
	col int
}

type queue []Position

func (q *queue) push(p Position) {
	*q = append(*q, p)
}
func (q *queue) pop() Position {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}
