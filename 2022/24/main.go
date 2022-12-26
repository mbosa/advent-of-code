package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

const inputFile = "input.txt"

var deltas = map[Position]bool{
	{-1, 0}: true, {1, 0}: true, {0, -1}: true, {0, 1}: true, {0, 0}: true,
}

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)
	valley := parse(rawInput)

	minToEnd := valley.bfs(0, valley.entrance, valley.exit)
	minToStart := valley.bfs(minToEnd, valley.exit, valley.entrance)
	minToEndAgain := valley.bfs(minToStart, valley.entrance, valley.exit)

	return minToEnd, minToEndAgain
}

type Position struct {
	x int
	y int
}

type State struct {
	pos Position
	min int
}

type Item struct {
	value    State
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

type Valley struct {
	width, height  int
	entrance, exit Position
	elements       map[Position]byte
}

func (v *Valley) bfs(min int, start, end Position) int {
	item := Item{value: State{pos: start, min: min}, priority: min + manhattanDistance(start, end)}
	q := PriorityQueue{&item}

	cycle := lcm(v.width-2, v.height-2)
	visited := map[State]bool{{pos: start, min: min}: true}

	for len(q) > 0 {
		currState := heap.Pop(&q).(*Item).value

		for d := range deltas {
			nextState := State{pos: Position{x: currState.pos.x + d.x, y: currState.pos.y + d.y}, min: currState.min + 1}

			visitedState := State{pos: nextState.pos, min: nextState.min % cycle}
			if _, ok := visited[visitedState]; ok {
				continue
			}

			if el, ok := v.elements[v.projectRight(nextState.pos, nextState.min)]; ok && el == '<' {
				continue
			}
			if el, ok := v.elements[v.projectLeft(nextState.pos, nextState.min)]; ok && el == '>' {
				continue
			}
			if el, ok := v.elements[v.projectDown(nextState.pos, nextState.min)]; ok && el == '^' {
				continue
			}
			if el, ok := v.elements[v.projectUp(nextState.pos, nextState.min)]; ok && el == 'v' {
				continue
			}

			if v.elements[nextState.pos] == '#' {
				continue
			}

			if nextState.pos.x < 0 || nextState.pos.x >= v.height || nextState.pos.y < 0 || nextState.pos.y >= v.width {
				continue
			}

			if nextState.pos == end {
				return nextState.min
			}

			item := Item{value: nextState, priority: currState.min + manhattanDistance(nextState.pos, end)}
			heap.Push(&q, &item)

			visited[State{pos: nextState.pos, min: nextState.min % cycle}] = true
		}
	}
	return -1
}

func (v *Valley) projectRight(pos Position, distance int) Position {
	yRight := (pos.y-1+distance)%(v.width-2) + 1

	return Position{x: pos.x, y: yRight}
}
func (v *Valley) projectLeft(pos Position, distance int) Position {
	yLeft := (pos.y - 1 - distance) % (v.width - 2)
	if yLeft < 0 {
		yLeft += (v.width - 2)
	}
	yLeft += 1

	return Position{x: pos.x, y: yLeft}
}
func (v *Valley) projectDown(pos Position, distance int) Position {
	xDown := (pos.x-1+distance)%(v.height-2) + 1

	return Position{x: xDown, y: pos.y}
}
func (v *Valley) projectUp(pos Position, distance int) Position {
	xUp := (pos.x - 1 - distance) % (v.height - 2)
	if xUp < 0 {
		xUp += (v.height - 2)
	}
	xUp += 1

	return Position{x: xUp, y: pos.y}
}

func parse(raw []byte) Valley {
	var width, height int
	var entrance, exit Position
	elements := map[Position]byte{}

	rows := bytes.Split(raw, []byte{'\n'})
	height, width = len(rows), len(rows[0])

	for i, row := range rows {
		for j, el := range row {
			elements[Position{x: i, y: j}] = el
		}
	}

	for i := 0; i < width; i++ {
		maybeEntrance, maybeExit := Position{x: 0, y: i}, Position{x: height - 1, y: i}
		if elements[maybeEntrance] == '.' {
			entrance = maybeEntrance
		}
		if elements[maybeExit] == '.' {
			exit = maybeExit
		}
	}

	return Valley{width: width, height: height, entrance: entrance, exit: exit, elements: elements}
}

func manhattanDistance(p1, p2 Position) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	return a * b / gcd(a, b)
}
