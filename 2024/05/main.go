package main

import (
	_ "embed"

	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type rulesMatrixT [100][100]int
type updateT []int

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	rules, updates := parseInput(input)

	res1 := part1(rules, updates)
	res2 := part2(rules, updates)

	return res1, res2
}

func part1(rules rulesMatrixT, updates []updateT) int {
	res := 0

	for _, update := range updates {
		if isUpdateCorrect(update, rules) {
			res += update[len(update)/2]
		}
	}

	return res
}

func part2(rules rulesMatrixT, updates []updateT) int {
	res := 0

	for _, update := range updates {
		if !isUpdateCorrect(update, rules) {
			slices.SortFunc(update, func(a, b int) int {
				return rules[a][b]
			})
			res += update[len(update)/2]
		}
	}

	return res
}

func parseInput(input string) (rulesMatrixT, []updateT) {
	var updates []updateT
	rules := [100][100]int{}

	spl := strings.Split(strings.TrimSpace(input), "\n\n")

	for _, line := range strings.Split(spl[0], "\n") {
		l := strings.Split(line, "|")
		r1, _ := strconv.Atoi(l[0])
		r2, _ := strconv.Atoi(l[1])

		rules[r1][r2] = -1 // correct order
		rules[r2][r1] = 1  // wrong order
	}

	for _, line := range strings.Split(spl[1], "\n") {
		l := strings.Split(line, ",")

		update := make(updateT, len(l))
		for i, ll := range l {
			update[i], _ = strconv.Atoi(ll)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func isUpdateCorrect(update updateT, rules rulesMatrixT) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			r1 := update[i]
			r2 := update[j]

			if rules[r1][r2] == 1 {
				return false
			}
		}
	}

	return true
}
