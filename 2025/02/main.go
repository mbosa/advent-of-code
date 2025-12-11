package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type parsedInputT [][2]int

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

func part1(parsedInput parsedInputT) int {
	res := 0

	for _, el := range parsedInput {
		for n := el[0]; n <= el[1]; n += 1 {
			s := strconv.Itoa(n)

			if len(s)%2 != 0 {
				continue
			}

			if s[:len(s)/2] == s[len(s)/2:] {
				res += n
			}
		}
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	for _, el := range parsedInput {
		for n := el[0]; n <= el[1]; n += 1 {
			s := strconv.Itoa(n)

		foo:
			for i := len(s) / 2; i > 0; i -= 1 {

				if len(s)%i != 0 {
					continue
				}

				for j := i; j < len(s); j += 1 {
					if s[j] != s[j-i] {
						continue foo
					}
				}

				res += n

				break
			}
		}
	}

	return res
}

func parseInput(input string) parsedInputT {
	ranges := strings.Split(strings.TrimSpace(input), ",")

	parsed := make([][2]int, len(ranges))

	for i, el := range ranges {
		limits := strings.Split(el, "-")

		a, _ := strconv.Atoi(limits[0])
		b, _ := strconv.Atoi(limits[1])

		parsed[i] = [2]int{a, b}
	}
	return parsed
}
