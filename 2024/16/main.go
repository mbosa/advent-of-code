package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"sync"

	_ "embed"
)

//go:embed input.txt
var input []byte

type position struct {
	r int
	c int
}

type direction struct {
	dr int
	dc int
}

type posDir struct {
	pos position
	dir direction
}

type pqItem struct {
	pos   position
	dir   direction
	score int
}

type priorityQ []*pqItem

func (pq priorityQ) Len() int { return len(pq) }

func (pq priorityQ) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq priorityQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQ) Push(x any) {
	item := x.(*pqItem)
	*pq = append(*pq, item)
}

func (pq *priorityQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[:n-1]
	return item
}

var (
	up    direction = direction{dr: -1, dc: 0}
	right           = direction{dr: 0, dc: 1}
	down            = direction{dr: 1, dc: 0}
	left            = direction{dr: 0, dc: -1}
)

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input []byte) (int, int) {
	grid := parseInput(input)

	res1, res2 := 0, 0

	startPos := position{}
	endPos := position{}
	for r, row := range grid {
		for c, b := range row {
			if b == 'S' {
				startPos.r = r
				startPos.c = c
				break
			}
		}

	}

	// pool of pqItem to reduce allocations
	pqItemPool := sync.Pool{
		New: func() any {
			return &pqItem{}
		},
	}

	startItem := pqItemPool.Get().(*pqItem)
	startItem.pos = startPos
	startItem.dir = right
	startItem.score = 0

	pq := priorityQ{startItem}

	// seen holds the best score for each direction
	seen := make([][][4]int, len(grid))
	for i := range grid {
		seen[i] = make([][4]int, len(grid[0]))
		for j := range seen[i] {
			seen[i][j] = [4]int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}
		}
	}

	seen[startPos.r][startPos.c][dirToIdx(right)] = 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)

		if grid[item.pos.r][item.pos.c] == 'E' {
			res1 = item.score

			endPos = item.pos
			break
		}

		// step
		nextPos := position{item.pos.r + item.dir.dr, item.pos.c + item.dir.dc}
		nextScore := item.score + 1
		if grid[nextPos.r][nextPos.c] != '#' {
			prevScore := seen[nextPos.r][nextPos.c][dirToIdx(item.dir)]
			if nextScore < prevScore {
				stepItem := pqItemPool.Get().(*pqItem)
				stepItem.pos = nextPos
				stepItem.dir = item.dir
				stepItem.score = nextScore

				seen[stepItem.pos.r][stepItem.pos.c][dirToIdx(stepItem.dir)] = stepItem.score

				heap.Push(&pq, stepItem)
			}
		}

		// rotate right
		rightRotationDir := direction{dr: item.dir.dc, dc: -item.dir.dr}
		prevScoreRightRot := seen[item.pos.r][item.pos.c][dirToIdx(rightRotationDir)]
		nextScoreRightRot := item.score + 1000
		if nextScoreRightRot < prevScoreRightRot {
			rightRotationItem := pqItemPool.Get().(*pqItem)
			rightRotationItem.pos = item.pos
			rightRotationItem.dir = rightRotationDir
			rightRotationItem.score = nextScoreRightRot

			seen[rightRotationItem.pos.r][rightRotationItem.pos.c][dirToIdx(rightRotationItem.dir)] = rightRotationItem.score

			heap.Push(&pq, rightRotationItem)
		}

		// rotate left
		leftRotationDir := direction{dr: -item.dir.dc, dc: item.dir.dr}
		prevScoreLeftRot := seen[item.pos.r][item.pos.c][dirToIdx(leftRotationDir)]
		nextScoreLeftRot := item.score + 1000
		if nextScoreLeftRot < prevScoreLeftRot {
			leftRotationItem := pqItemPool.Get().(*pqItem)
			leftRotationItem.pos = item.pos
			leftRotationItem.dir = leftRotationDir
			leftRotationItem.score = nextScoreLeftRot

			seen[leftRotationItem.pos.r][leftRotationItem.pos.c][dirToIdx(leftRotationItem.dir)] = leftRotationItem.score

			heap.Push(&pq, leftRotationItem)
		}

		pqItemPool.Put(item)
	}

	res2 = backtrack(seen, endPos, res1)

	return res1, res2
}

// Once the end is reached using dijkstra, there's not gonna be any other path
// with a better score, and all the paths with the same score will already be explored.
// I can backtrack from the end, moving through all the positions with a score lower
// the the previous position. This way I will explore all the paths with the best score
func backtrack(seen [][][4]int, end position, score int) int {
	type posScore struct {
		pos   position
		score int
	}

	uniquePos := map[position]struct{}{}

	endItem := posScore{pos: end, score: score}

	// DFS to explore all the best paths
	stack := []posScore{endItem}

	directions := [4]direction{up, right, down, left}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		uniquePos[item.pos] = struct{}{}

		for _, d := range directions {
			nextPos := position{
				r: item.pos.r - d.dr,
				c: item.pos.c - d.dc,
			}

			prevScore := seen[nextPos.r][nextPos.c][dirToIdx(d)]
			if prevScore < item.score {
				nextItem := posScore{
					pos:   nextPos,
					score: prevScore,
				}

				stack = append(stack, nextItem)
			}
		}
	}

	return len(uniquePos)
}

func parseInput(input []byte) [][]byte {
	l := len(input) - 1 // last byte can be a newline

	width := slices.Index(input, '\n')
	height := l / width

	grid := make([][]byte, height)

	for i, j := 0, 0; i < l; i, j = i+width+1, j+1 {
		grid[j] = input[i : i+width]
	}

	return grid
}

func dirToIdx(dir direction) int {
	switch dir {
	case up:
		return 0
	case right:
		return 1
	case down:
		return 2
	case left:
		return 3
	}

	panic("invalid direction")
}
