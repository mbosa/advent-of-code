package main

import (
	"fmt"

	_ "embed"
)

//go:embed input.txt
var input string

type parsedInputT []string

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

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	return res
}

func parseInput(input string) parsedInputT {

}
