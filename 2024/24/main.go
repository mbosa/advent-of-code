package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type operatorT string

type parsedInputT struct {
	wireValues map[string]func() int
	operations []operationT
}

type operationT struct {
	operand1 string
	operand2 string
	operator operatorT
	result   string
}

type wireAndOperator struct {
	wire     string
	operator operatorT
}

const (
	AND operatorT = "AND"
	OR            = "OR"
	XOR           = "XOR"
)

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, string) {
	m := parseInput(input)

	res1 := part1(m)
	res2 := part2(m)

	return res1, res2
}

func part1(parsedInput parsedInputT) int {
	res := 0
	wireValues := parsedInput.wireValues

	i := 0
	for {
		z := fmt.Sprintf("z%02d", i)

		if value, ok := wireValues[z]; ok {
			res += value() * (1 << i)

			i += 1
		} else {
			break
		}
	}

	return res
}

/*
The device is a Ripple-carry adder.
Each adder receives 3 inputs: A, B, and a carry C_in from the previous adder,
and outputs a sum of the inputs, S, and a carry C_out

S = A XOR B XOR C_in
C_out = (A AND B) OR (C_in AND (A XOR B))

Notes:
- The first adder doesn't receive a C_in, since there is no adder before it (C_in = 0)
- The last S is the C_out of the last adder (A = 0, B = 0)

e.g.
input x10, y10, dgr; output z10, cjv
where z10 = (x10 XOR y10) XOR dgr
cjv = (x10 AND y10) OR (dgr AND (x10 XOR y10))
cjv will be the C_in for the next adder, that will receive x11, y11, cjv

Rules to check to find the misplaced gates:
- first z: x00 XOR y00 -> z00
- last z: z45 must be the outcome of an OR operation
- any other z must be the outcome of XOR between non-x and non-y values
- x__ XOR y__ and x__ AND y__ should not outcome z__, except for z00
- the result of each AND operation must be used in an OR operation, except for x00 AND y00

The problem says some output wires were swapped. I will assume the operations are all valid
*/
func part2(parsedInput parsedInputT) string {
	badWires := make([]string, 0, 8)

	lookup := make(map[wireAndOperator]struct{}, len(parsedInput.operations)*2)

	for _, operation := range parsedInput.operations {
		operand1, operand2, operator := operation.operand1, operation.operand2, operation.operator
		wo1 := wireAndOperator{operand1, operator}
		wo2 := wireAndOperator{operand2, operator}

		lookup[wo1] = struct{}{}
		lookup[wo2] = struct{}{}
	}

	for _, operation := range parsedInput.operations {
		operator, result := operation.operator, operation.result

		switch operator {
		case XOR:
			if !isValidXOR(operation, lookup) {
				badWires = append(badWires, result)
			}
		case AND:
			if !isValidAND(operation, lookup) {
				badWires = append(badWires, result)
			}
		case OR:
			if !isValidOR(operation, lookup) {
				badWires = append(badWires, result)
			}
		}
	}

	sort.Strings(badWires)

	return strings.Join(badWires, ",")
}

func parseInput(input string) parsedInputT {
	spl := strings.Split(strings.TrimSpace(input), "\n\n")
	operationsLines := strings.Split(spl[1], "\n")

	wireValues := map[string]func() int{}
	operations := make([]operationT, len(operationsLines))

	for _, line := range strings.Split(spl[0], "\n") {
		wire := line[:3]
		val := int(line[5] - '0')

		wireValues[wire] = func() int { return val }
	}

	for i, line := range operationsLines {
		spl := strings.Split(line, " ")

		operand1, operator, operand2, result := spl[0], operatorT(spl[1]), spl[2], spl[4]

		wireValues[result] = func() int {
			return calc(wireValues[operand1](), wireValues[operand2](), operator)
		}

		operation := operationT{operand1, operand2, operator, result}
		operations[i] = operation
	}

	return parsedInputT{wireValues, operations}
}

func calc(operand1, operand2 int, operator operatorT) int {
	switch operator {
	case AND:
		return operand1 & operand2
	case OR:
		return operand1 | operand2
	case XOR:
		return operand1 ^ operand2
	}

	panic("operator not supported")
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(s + " is not a number")
	}

	return n
}

func isX(s string) bool { return s[0] == 'x' }
func isY(s string) bool { return s[0] == 'y' }
func isZ(s string) bool { return s[0] == 'z' }

// Not exhaustive for a ripple-carry adder in general. It checks just what's needed for this problem
func isValidXOR(operation operationT, lookup map[wireAndOperator]struct{}) bool {
	operand1, operand2, operator, result := operation.operand1, operation.operand2, operation.operator, operation.result

	if operator != XOR {
		return false
	}

	// first adder
	if operand1 == "x00" || operand2 == "x00" {
		// there is no C_in for the first adder, so it must output z
		return result == "z00"
	}

	// x__ XOR y__, or y__ XOR x__
	if isX(operand1) || isX(operand2) {
		// except for the first adder, this cannot output z, since it must consider C_in
		if isZ(result) {
			return false
		}

		// the result must be used in a following XOR
		expected := wireAndOperator{result, XOR}
		_, ok := lookup[expected]

		return ok
	}

	// ___ XOR ___ must output z
	return isZ(result)
}

// Not exhaustive for a ripple-carry adder in general. It checks just what's needed for this problem
func isValidAND(operation operationT, lookup map[wireAndOperator]struct{}) bool {
	operand1, operand2, operator, result := operation.operand1, operation.operand2, operation.operator, operation.result

	if operator != AND {
		return false
	}

	// z cannot be the outcome of AND
	if isZ(result) {
		return false
	}

	// first adder
	if operand1 == "x00" || operand2 == "x00" {
		// there is no C_in for the first adder, so AND outputs C_out, that must be used in a following XOR
		expected := wireAndOperator{result, XOR}
		_, ok := lookup[expected]

		return ok
	}

	// except for the first adder, the outcome of each AND is used in a OR
	expected := wireAndOperator{result, OR}
	_, ok := lookup[expected]

	return ok
}

// Not exhaustive for a ripple-carry adder in general. It checks just what's needed for this problem
func isValidOR(operation operationT, lookup map[wireAndOperator]struct{}) bool {
	_ = lookup

	operator, result := operation.operator, operation.result

	if operator != OR {
		return false
	}

	// OR can only output z for the last bit
	if isZ(result) {
		return result == "z45"
	}

	return true
}
