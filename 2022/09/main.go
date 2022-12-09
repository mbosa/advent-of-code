package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input.txt"

const (
	right = 'R'
	left  = 'L'
	up    = 'U'
	down  = 'D'
)

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rope1, rope2 := newRope(2), newRope(10)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		dir := line[0]
		val := bytesToInt(line[2:])

		for i := 0; i < val; i++ {
			rope1.move(dir)
			rope2.move(dir)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	resPart1, resPart2 = len(rope1.tailPositionsRegistry), len(rope2.tailPositionsRegistry)

	return resPart1, resPart2
}

func newRope(n int) *rope {
	r := &rope{
		knots:                 make([]position, n),
		tailPositionsRegistry: make(map[position]bool),
	}
	r.tailPositionsRegistry[r.knots[n-1]] = true

	return r
}

type position struct {
	x int
	y int
}

type rope struct {
	knots                 []position
	tailPositionsRegistry map[position]bool
}

func (r *rope) move(dir byte) {
	head := &r.knots[0]
	if dir == right {
		head.x++
	} else if dir == left {
		head.x--
	} else if dir == up {
		head.y++
	} else if dir == down {
		head.y--
	}

	r.updateBody()
}
func (r *rope) updateBody() {
	n := len(r.knots)
	for i := 1; i < n; i++ {
		prev := r.knots[i-1]
		curr := &r.knots[i]

		deltaX := prev.x - curr.x
		deltaY := prev.y - curr.y

		if abs(deltaX) > 1 || abs(deltaY) > 1 {
			curr.x += sign(deltaX)
			curr.y += sign(deltaY)
		}
	}

	r.recordTailPos()
}
func (r *rope) recordTailPos() {
	n := len(r.knots)
	r.tailPositionsRegistry[r.knots[n-1]] = true
}

func bytesToInt(b []byte) int {
	c, _ := strconv.Atoi(string(b))
	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}
