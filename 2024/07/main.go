package main

import (
	_ "embed"

	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type equationT struct {
	result   int
	operands []int
}

type equationsT []equationT

type operatorT int

const (
	add operatorT = iota
	mul
	con
)

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

func part1(equations equationsT) int {
	res := 0

	operators := []operatorT{add, mul}
	for _, equation := range equations {
		if isEquationValid(equation, operators) {
			res += equation.result
		}
	}

	return res
}

func part2(equations equationsT) int {
	res := 0

	operators := []operatorT{add, mul, con}
	for _, equation := range equations {
		if isEquationValid(equation, operators) {
			res += equation.result
		}
	}

	return res
}

func parseInput(input string) equationsT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	parsed := make(equationsT, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		result, _ := strconv.Atoi(fields[0][:len(fields[0])-1])
		operands := make([]int, len(fields)-1)

		for j := 1; j < len(fields); j++ {
			operands[j-1], _ = strconv.Atoi(fields[j])
		}

		equation := equationT{result, operands}

		parsed[i] = equation
	}

	return parsed
}

// e.g. 190: 10 19 -> 190/19 = 10; 156: 15 6 -> 156 split 6 = 15
func isEquationValid(equation equationT, operators []operatorT) bool {
	operands := equation.operands
	result := equation.result

	var helper func(res, i int) bool
	helper = func(res, i int) bool {
		operand := operands[i]

		if i == 0 {
			return res == operand
		}

		for _, operator := range operators {
			switch operator {
			case add:
				if helper(res-operand, i-1) {
					return true
				}
			case mul:
				if res%operand == 0 && helper(res/operand, i-1) {
					return true
				}
			case con:
				if r := getPartBefore(res, operand); r > 0 && helper(r, i-1) {
					return true
				}
			}
		}

		return false
	}

	return helper(result, len(operands)-1)
}

// e.g. 2 -> 10; 15 -> 100; 123 -> 1000
func getNumFactor(n int) int {
	factor := 1
	for n > 0 {
		factor *= 10
		n /= 10
	}
	return factor
}

// e.g. 12,25 -> 1200 + 25 = 1225
func concatenateInts(a, b int) int {
	bFactor := getNumFactor(b)

	return a*bFactor + b

}

// e.g. 123,23 -> 1; 428,8 -> 42. Returns -1 if a does not end with b
func getPartBefore(a, b int) int {
	bFactor := getNumFactor(b)

	if a%bFactor == b {
		return a / bFactor
	}

	return -1
}
