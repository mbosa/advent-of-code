package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type machine struct {
	a     position
	b     position
	prize position
}

type position struct {
	x int
	y int
}

var machineRegex = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	machines := parseInput(input)

	res1 := part1(machines)
	res2 := part2(machines)

	return res1, res2
}

func part1(machines []machine) int {
	res := 0

	for _, m := range machines {
		res += costToPrize(m, 0, 100)
	}

	return res
}

func part2(machines []machine) int {
	modifier := 10000000000000

	res := 0

	for _, m := range machines {
		res += costToPrize(m, modifier, -1)
	}

	return res
}

func parseInput(input string) []machine {
	machineStrs := strings.Split(strings.TrimSpace(input), "\n\n")

	machines := make([]machine, len(machineStrs))

	for i, machineStr := range machineStrs {
		matches := machineRegex.FindStringSubmatch(machineStr)

		a := position{x: mustAtoi(matches[1]), y: mustAtoi(matches[2])}
		b := position{x: mustAtoi(matches[3]), y: mustAtoi(matches[4])}
		prize := position{x: mustAtoi(matches[5]), y: mustAtoi(matches[6])}

		m := machine{a, b, prize}

		machines[i] = m
	}

	return machines
}

func costToPrize(m machine, modifier, pressLimit int) int {
	ax, ay, bx, by := m.a.x, m.a.y, m.b.x, m.b.y
	px, py := m.prize.x+modifier, m.prize.y+modifier

	bPresses := (py*ax - ay*px) / (ax*by - ay*bx)
	aPresses := (px - bx*bPresses) / ax

	if aPresses*ax+bPresses*bx != px {
		return 0
	}

	if aPresses*ay+bPresses*by != py {
		return 0
	}

	if pressLimit > 0 {
		if aPresses > 100 || bPresses > 100 {
			return 0
		}
	}

	return aPresses*3 + bPresses
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("Can't convert, received: " + s)
	}
	return n
}
