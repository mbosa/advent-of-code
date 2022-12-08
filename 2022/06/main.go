package main

import (
	_ "embed"

	"fmt"
	"sync"
)

//go:embed input.txt
var input []byte

func main() {
	part1, part2 := solve(input)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(input []byte) (resPart1, resPart2 int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		resPart1 = firstSequenceOfUniqueChars(input, 4)
		wg.Done()
	}()
	go func() {
		resPart2 = firstSequenceOfUniqueChars(input, 14)
		wg.Done()
	}()

	wg.Wait()

	return resPart1, resPart2
}

func firstSequenceOfUniqueChars(sl []byte, n int) int {
	chars := [26]int{}
	dups := 0
	i, j := 0, 0

	for j < n {
		idx := byteToIdx(sl[j])
		chars[idx]++
		if chars[idx] == 2 {
			dups++
		}
		j++
	}

	if dups == 0 {
		return j
	}

	for j < len(sl) {
		in, out := sl[j], sl[i]
		idxJ, idxI := byteToIdx(in), byteToIdx(out)

		chars[idxI]--
		if chars[idxI] == 1 {
			dups--
		}
		chars[idxJ]++
		if chars[idxJ] == 2 {
			dups++
		}

		i++
		j++

		if dups == 0 {
			return j
		}
	}
	return -1
}

func byteToIdx(b byte) int {
	return int(b - 'a')
}
