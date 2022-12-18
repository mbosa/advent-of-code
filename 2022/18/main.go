package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

const inputFile = "input.txt"

var delta = []Position{
	{x: 1, y: 0, z: 0},
	{x: -1, y: 0, z: 0},
	{x: 0, y: 1, z: 0},
	{x: 0, y: -1, z: 0},
	{x: 0, y: 0, z: 1},
	{x: 0, y: 0, z: -1},
}

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)
	lines := bytes.Split(rawInput, []byte{'\n'})

	droplets := make(map[Position]bool)

	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt

	for _, line := range lines {
		coords := bytes.Split(line, []byte{','})
		x, y, z := bytesToInt(coords[0]), bytesToInt(coords[1]), bytesToInt(coords[2])

		minX, minY, minZ = min(minX, x), min(minY, y), min(minZ, z)
		maxX, maxY, maxZ = max(maxX, x), max(maxY, y), max(maxZ, z)

		pos := Position{x, y, z}
		droplets[pos] = true
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		resPart1 = part1(droplets)
		wg.Done()
	}()

	go func() {
		minPos := Position{minX - 1, minY - 1, minZ - 1}
		maxPos := Position{maxX + 1, maxY + 1, maxZ + 1}

		resPart2 = part2(droplets, minPos, maxPos)
		wg.Done()
	}()

	wg.Wait()
	return resPart1, resPart2
}

type Position struct {
	x int
	y int
	z int
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

func part1(droplets map[Position]bool) int {
	res := 0
	for droplet := range droplets {
		neighbors := getNeighbors(droplet)

		for _, n := range neighbors {
			if _, ok := droplets[n]; !ok {
				res++
			}
		}
	}
	return res
}

func part2(droplets map[Position]bool, minPos, maxPos Position) int {
	res := 0

	// bfs
	q := queue{minPos}
	visited := map[Position]bool{}
	for len(q) > 0 {
		curr := q.pop()

		neighbors := getNeighbors(curr)
		for _, next := range neighbors {
			if next.x < minPos.x || next.y < minPos.y || next.z < minPos.z || next.x > maxPos.x || next.y > maxPos.y || next.z > maxPos.z {
				continue
			}

			if _, ok := visited[next]; ok {
				continue
			}

			if _, ok := droplets[next]; ok {
				res++
				continue
			}

			visited[next] = true
			q.push(next)
		}
	}

	return res
}

func getNeighbors(pos Position) []Position {
	res := make([]Position, 6)

	for i, d := range delta {
		p := Position{
			x: pos.x + d.x,
			y: pos.y + d.y,
			z: pos.z + d.z,
		}
		res[i] = p
	}
	return res
}

func bytesToInt(b []byte) int {
	c, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
