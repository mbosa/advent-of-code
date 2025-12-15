package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type boxT struct {
	x int
	y int
	z int
}

type parsedInputT []boxT

type pairWithDistanceT struct {
	boxA     boxT
	boxB     boxT
	distance float64
}

// disjoint set union - https://cp-algorithms.com/data_structures/disjoint_set_union.html
type DSU struct {
	parent map[boxT]boxT
	size   map[boxT]int
}

func (dsu *DSU) makeSet(el boxT) {
	dsu.parent[el] = el
	dsu.size[el] = 1
}
func (dsu DSU) findSet(el boxT) boxT {
	if dsu.parent[el] == el {
		return el
	}

	return dsu.findSet(dsu.parent[el])
}
func (dsu *DSU) unionSets(a, b boxT) {
	setA := dsu.findSet(a)
	setB := dsu.findSet(b)

	if setA != setB {
		if dsu.size[setA] < dsu.size[setB] {
			dsu.parent[setA] = setB
			dsu.size[setB] += dsu.size[setA]
		} else {
			dsu.parent[setB] = setA
			dsu.size[setA] += dsu.size[setB]
		}
	}
}

func newDSU(items []boxT) DSU {
	dsu := DSU{parent: map[boxT]boxT{}, size: map[boxT]int{}}

	for _, item := range items {
		dsu.makeSet(item)
	}
	return dsu
}

func main() {
	res1, res2 := solve(input, 1000)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string, connections int) (int, int) {
	parsedInput := parseInput(input)

	res1 := part1(parsedInput, connections)
	res2 := part2(parsedInput)

	return res1, res2
}

func part1(parsedInput parsedInputT, connections int) int {
	// calculate the distance between each pair of boxes
	distances := make([]pairWithDistanceT, 0)

	for i := 0; i < len(parsedInput)-1; i += 1 {
		boxA := parsedInput[i]
		for j := i + 1; j < len(parsedInput); j += 1 {
			boxB := parsedInput[j]
			distance := math.Sqrt(math.Pow(float64(boxA.x-boxB.x), 2) + math.Pow(float64(boxA.y-boxB.y), 2) + math.Pow(float64(boxA.z-boxB.z), 2))

			pairWithDistance := pairWithDistanceT{boxA, boxB, distance}
			distances = append(distances, pairWithDistance)
		}
	}

	// sort the distances
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	// create a DSU
	dsu := newDSU(parsedInput)

	// join the N closest pairs
	for i := range connections {
		dsu.unionSets(distances[i].boxA, distances[i].boxB)
	}

	// find the 3 largest sets in the DSU
	largest, secondLargest, thirdLargest := 0, 0, 0
	for _, v := range dsu.size {
		if v > largest {
			thirdLargest = secondLargest
			secondLargest = largest
			largest = v
		} else if v > secondLargest {
			thirdLargest = secondLargest
			secondLargest = v
		} else if v > thirdLargest {
			thirdLargest = v
		}
	}

	return largest * secondLargest * thirdLargest
}

func part2(parsedInput parsedInputT) int {
	// calculate the distance between each pair of boxes
	distances := make([]pairWithDistanceT, 0)

	for i := 0; i < len(parsedInput)-1; i += 1 {
		boxA := parsedInput[i]
		for j := i + 1; j < len(parsedInput); j += 1 {
			boxB := parsedInput[j]
			distance := math.Sqrt(math.Pow(float64(boxA.x-boxB.x), 2) + math.Pow(float64(boxA.y-boxB.y), 2) + math.Pow(float64(boxA.z-boxB.z), 2))

			pairWithDistance := pairWithDistanceT{boxA, boxB, distance}
			distances = append(distances, pairWithDistance)
		}
	}

	// sort the distances
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	// create a DSU
	dsu := newDSU(parsedInput)

	// join the closest pairs until a union set includes all the boxes
	for i := range distances {
		dsu.unionSets(distances[i].boxA, distances[i].boxB)

		for _, v := range dsu.size {
			if v == len(parsedInput) {
				return distances[i].boxA.x * distances[i].boxB.x
			}

		}
	}
	return 0
}

func parseInput(input string) parsedInputT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	parsed := make([]boxT, len(lines))

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		box := boxT{x, y, z}

		parsed[i] = box
	}

	return parsed
}
