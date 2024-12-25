package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	secrets := parseInput(input)

	res1, res2 := 0, 0

	sequences := make(map[[4]int]int, 1<<17)
	seen := make(map[[4]int]struct{}, 1<<13)

	for _, s := range secrets {
		// clean the existing set instead of creating a new one to save allocations
		cleanSeen(seen)

		prevLastDigit := s % 10
		seq := [4]int{}

		// invalid sequences
		for range 3 {
			s = nextSecret(s)

			lastDigit := s % 10
			change := lastDigit - prevLastDigit
			prevLastDigit = lastDigit

			seq[0], seq[1], seq[2], seq[3] = seq[1], seq[2], seq[3], change
		}

		// valid sequences
		for range 2000 - 3 {
			s = nextSecret(s)

			lastDigit := s % 10
			change := lastDigit - prevLastDigit
			prevLastDigit = lastDigit

			seq[0], seq[1], seq[2], seq[3] = seq[1], seq[2], seq[3], change

			if _, ok := seen[seq]; !ok {
				seen[seq] = struct{}{}
				sequences[seq] += lastDigit
			}
		}

		res1 += s
	}

	for _, v := range sequences {
		if v > res2 {
			res2 = v
		}
	}

	return res1, res2
}

func parseInput(input string) []int {
	spl := strings.Split(strings.TrimSpace(input), "\n")

	secrets := make([]int, len(spl))

	for i, line := range spl {
		secrets[i] = mustAtoi(line)
	}

	return secrets
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		panic("not a valid number: " + s)
	}

	return n
}

func nextSecret(secret int) int {
	t := secret << 6 // secret * 64
	secret ^= t
	secret &= 16777215 // secret % 16777216 - 16777216 == 2^24

	t = secret >> 5 // secret / 32
	secret ^= t
	secret &= 16777215 // secret % 16777216

	t = secret << 11 // secret * 2048
	secret ^= t
	secret &= 16777215 // secret % 16777216

	return secret
}

func cleanSeen(seen map[[4]int]struct{}) {
	for k := range seen {
		delete(seen, k)
	}
}
