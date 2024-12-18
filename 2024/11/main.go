package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type cacheItem struct {
	num    int
	blinks int
}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	parsedInput := parseInput(input)

	res1 := part1(parsedInput)
	res2 := part2(parsedInput)

	return res1, res2
}

func part1(stones []int) int {
	return count(stones, 25)
}

func part2(stones []int) int {
	return count(stones, 75)
}

func solveMemo(input string) (int, int) {
	parsedInput := parseInput(input)

	cache := map[cacheItem]int{}

	res1 := part1Memo(parsedInput, cache)
	res2 := part2Memo(parsedInput, cache)

	return res1, res2
}

func part1Memo(stones []int, cache map[cacheItem]int) int {
	res := 0

	for _, s := range stones {
		res += countPerStoneMemo(s, 25, cache)
	}

	return res
}

func part2Memo(stones []int, cache map[cacheItem]int) int {
	res := 0

	for _, s := range stones {
		res += countPerStoneMemo(s, 75, cache)
	}

	return res
}

func parseInput(input string) []int {
	fields := strings.Fields(input)

	parsed := make([]int, len(fields))

	for i, f := range fields {
		parsed[i], _ = strconv.Atoi(f)
	}
	return parsed
}

func countDigits(n int) int {
	count := 0

	for n > 0 {
		count += 1
		n /= 10
	}
	return count
}

func splitInt(n int) (int, int) {
	digits := countDigits(n)

	divisor := int(math.Pow10(digits / 2))

	n1 := n / divisor
	n2 := n % divisor

	return n1, n2
}

func count(stones []int, blinks int) int {
	// map[stone number]index
	indexMap := make(map[int]int, 4096)
	// quantity of each stone number
	quantities := make([]int, 4096)

	for _, s := range stones {
		index := len(indexMap)
		indexMap[s] = index
		quantities[index] = 1
	}

	for range blinks {
		// quantities after the blink
		nextQuantities := make([]int, 4096)

		// loop over all the stone numbers found
		// introducing a new set in order to reduce the iterations made the function slower
		for num, index := range indexMap {
			q := quantities[index]

			if q == 0 {
				continue
			}

			if num == 0 {
				nextNum := 1
				nextIndex, ok := indexMap[nextNum]

				if !ok {
					nextIndex = len(indexMap)
					indexMap[nextNum] = nextIndex
				}

				nextQuantities[nextIndex] += q
			} else if countDigits(num)%2 == 0 {
				nextNum1, nextNum2 := splitInt(num)

				nextIndex1, ok1 := indexMap[nextNum1]
				if !ok1 {
					nextIndex1 = len(indexMap)
					indexMap[nextNum1] = nextIndex1
				}

				nextIndex2, ok2 := indexMap[nextNum2]
				if !ok2 {
					nextIndex2 = len(indexMap)
					indexMap[nextNum2] = nextIndex2
				}

				nextQuantities[nextIndex1] += q
				nextQuantities[nextIndex2] += q
			} else {
				nextNum := num * 2024
				nextIndex, ok := indexMap[nextNum]

				if !ok {
					nextIndex = len(indexMap)
					indexMap[nextNum] = nextIndex
				}

				nextQuantities[nextIndex] += q
			}
		}

		quantities = nextQuantities
	}

	res := 0
	for _, q := range quantities {
		res += q
	}

	return res
}

func countPerStoneMemo(num, blinks int, cache map[cacheItem]int) int {
	var helper func(num, blinks int) int
	helper = func(num, blinks int) int {
		if blinks == 0 {
			return 1
		}

		c := cacheItem{num, blinks}
		if v, ok := cache[c]; ok {
			return v
		}

		res := 0
		if num == 0 {
			res = helper(1, blinks-1)
		} else if countDigits(num)%2 == 0 {
			s1, s2 := splitInt(num)
			res = helper(s1, blinks-1) + helper(s2, blinks-1)
		} else {
			res = helper(num*2024, blinks-1)
		}

		cache[c] = res
		return res
	}

	return helper(num, blinks)
}
