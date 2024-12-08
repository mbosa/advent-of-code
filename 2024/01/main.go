package main

import (
	_ "embed"

	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type parsedInputT [2][]int

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

	aList := parsedInput[0]
	bList := parsedInput[1]

	for i := range aList {
		res += max(aList[i], bList[i]) - min(aList[i], bList[i])
	}

	return res
}

func part2(parsedInput parsedInputT) int {
	res := 0

	aList := parsedInput[0]
	bList := parsedInput[1]

	bListMap := make(map[int]int)

	for _, b := range bList {
		bListMap[b] += 1
	}

	for _, a := range aList {
		res += a * bListMap[a]
	}

	return res
}

func parseInput(input string) parsedInputT {
	ids := strings.Fields(input)

	listsLen := len(ids) / 2

	aList := make([]int, 0, listsLen)
	bList := make([]int, 0, listsLen)

	for i, id := range ids {
		idNum, _ := strconv.Atoi(id)

		if i%2 == 0 {
			aList = append(aList, idNum)
		} else {
			bList = append(bList, idNum)
		}
	}

	sort.Ints(aList)
	sort.Ints(bList)

	return parsedInputT{aList, bList}
}
