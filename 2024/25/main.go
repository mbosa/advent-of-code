package main

import (
	"fmt"
	"slices"

	_ "embed"
)

//go:embed input.txt
var input []byte

func main() {
	res1 := solve(input)

	fmt.Println("part1:", res1)
}

func solve(input []byte) int {
	locks, keys := parseInput(input)

	res1 := part1(locks, keys)

	return res1
}

func part1(locks, keys []uint32) int {
	res := 0

	for _, lock := range locks {
		for _, key := range keys {
			// if the bitmasks of lock and key don't overlap
			if lock&key == 0 {
				res += 1
			}
		}
	}

	return res
}

// Return two lists of bitmasks, one for locks, and one for keys
// One bitmask fits in an uint32
func parseInput(input []byte) ([]uint32, []uint32) {
	lockMasks, keyMasks := make([]uint32, 0), make([]uint32, 0)

	// a lock or key is 41 bytes long, and they are divided by 2 new lines
	for chunk := range slices.Chunk(input, 43) {
		isLock := chunk[0] == '#'

		var bitmask uint32 = 0

		// the top and bottom row of each lock and key are constant. They can be ignored
		for _, b := range chunk[6:35] {
			bitmask = (bitmask << 1) | (uint32(b) & 1)
		}

		if isLock {
			lockMasks = append(lockMasks, bitmask)
		} else {
			keyMasks = append(keyMasks, bitmask)
		}
	}

	return lockMasks, keyMasks
}
