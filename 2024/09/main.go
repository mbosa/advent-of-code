package main

import (
	"fmt"

	_ "embed"
)

//go:embed input.txt
var input []byte

type diskMapT []int

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input []byte) (int, int) {
	diskMap := parseInput(input)

	res1 := part1(diskMap)
	res2 := part2(diskMap)

	return res1, res2
}

func part1(diskMapSrc diskMapT) int {
	res := 0

	diskMap := make(diskMapT, len(diskMapSrc))
	copy(diskMap, diskMapSrc)

	i := 0
	// last even index
	j := len(diskMap) - 2 + len(diskMap)%2
	diskIndex := 0

	for i <= j {
		if diskMap[i] == 0 {
			i += 1
			continue
		}
		if diskMap[j] == 0 {
			j -= 2
			continue
		}

		if i%2 == 0 {
			id := i / 2

			res += id * diskIndex
			diskIndex += 1

			diskMap[i] -= 1
		} else {
			id := j / 2

			res += id * diskIndex
			diskIndex += 1

			diskMap[i] -= 1
			diskMap[j] -= 1
		}
	}

	return res
}

func part2(diskMapSrc diskMapT) int {
	res := 0

	diskMap := make(diskMapT, len(diskMapSrc))
	copy(diskMap, diskMapSrc)

	// start index of each block
	indexes := make([]int, len(diskMap))
	indexes[0] = 0
	for i := 1; i < len(diskMap); i++ {

		indexes[i] = indexes[i-1] + diskMap[i-1]
	}

	// last even index
	lastEvenIdx := len(diskMap) - 2 + len(diskMap)%2

	// add checksum of files that are moved
	// loop files from last to first
	for i := lastEvenIdx; i > -1; i -= 2 {
		if diskMap[i] == 0 {
			continue
		}

		// loop blanks from first to last
		for j := 1; j < i; j += 2 {
			if diskMap[i] > diskMap[j] {
				continue
			}

			id := i / 2
			idx := indexes[j]

			// the file is moved, add checksum
			for range diskMap[i] {
				res += id * idx
				idx += 1
			}

			// reduce size of blank due to the file being moved
			diskMap[j] -= diskMap[i]
			// increase start index of the blank block due to the file being moved
			indexes[j] += diskMap[i]
			// increase blank space before the moved file
			diskMap[i-1] += diskMap[i]
			// remove the moved file
			diskMap[i] = 0

			break
		}
	}

	// add checksum of files that were not moved
	for i := 0; i < len(diskMap); i += 2 {
		id := i / 2
		idx := indexes[i]
		for range diskMap[i] {
			res += id * idx
			idx += 1
		}
	}

	return res
}

func parseInput(input []byte) diskMapT {
	diskMap := make([]int, len(input))
	for i, b := range input {
		diskMap[i] = int(b - '0')
	}
	return diskMap
}
