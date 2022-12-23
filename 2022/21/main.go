package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const inputFile = "input.txt"

var evalOperationMap = map[byte]func(operand1, operand2 int) int{
	'+': func(operand1, operand2 int) int { return operand1 + operand2 },
	'-': func(operand1, operand2 int) int { return operand1 - operand2 },
	'*': func(operand1, operand2 int) int { return operand1 * operand2 },
	'/': func(operand1, operand2 int) int { return operand1 / operand2 },
}

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)

	monkeys := parse(rawInput)

	resPart1 = monkeys.ResolveMonkey("root")

	newRoot := monkeys["root"]
	newRoot.operation.symbol = '-'
	monkeys["root"] = newRoot

	resPart2 = monkeys.ResolveForHumnToTarget("root", 0)

	return resPart1, resPart2
}

type Job struct {
	value     int
	operation Operation
}

type Operation struct {
	left   string
	right  string
	symbol byte
}

type Monkeys map[string]Job

func (m *Monkeys) ResolveMonkey(name string) int {
	if (*m)[name].value != 0 {
		return (*m)[name].value
	}

	left := m.ResolveMonkey((*m)[name].operation.left)
	right := m.ResolveMonkey((*m)[name].operation.right)

	res := evalOperationMap[(*m)[name].operation.symbol](left, right)

	return res
}

func (m *Monkeys) ResolveForHumnToTarget(name string, target int) int {
	if name == "humn" {
		return target
	}

	if (*m)[name].value != 0 {
		return (*m)[name].value
	}

	if !m.DerivesFromHumn(name) {
		return m.ResolveMonkey(name)
	}

	left, right := (*m)[name].operation.left, (*m)[name].operation.right

	if !m.DerivesFromHumn(left) {
		val := m.ResolveMonkey(left)

		switch (*m)[name].operation.symbol {
		case '+':
			return m.ResolveForHumnToTarget(right, target-val)
		case '-':
			return m.ResolveForHumnToTarget(right, val-target)
		case '*':
			return m.ResolveForHumnToTarget(right, target/val)
		case '/':
			return m.ResolveForHumnToTarget(right, val/target)
		}
	} else if !m.DerivesFromHumn(right) {
		val := m.ResolveMonkey(right)

		switch (*m)[name].operation.symbol {
		case '+':
			return m.ResolveForHumnToTarget(left, target-val)
		case '-':
			return m.ResolveForHumnToTarget(left, target+val)
		case '*':
			return m.ResolveForHumnToTarget(left, target/val)
		case '/':
			return m.ResolveForHumnToTarget(left, target*val)
		}
	}

	panic("should return before this")
}

func (m *Monkeys) DerivesFromHumn(name string) bool {
	if name == "humn" {
		return true
	}
	if (*m)[name].value != 0 {
		return false
	}

	return m.DerivesFromHumn((*m)[name].operation.left) || m.DerivesFromHumn((*m)[name].operation.right)
}

func parse(raw []byte) Monkeys {
	monkeys := Monkeys{}

	for _, line := range bytes.Split(raw, []byte{'\n'}) {
		fields := bytes.Split(line, []byte{':', ' '})
		name := string(fields[0])
		operationFields := bytes.Split(fields[1], []byte{' '})

		if len(operationFields) == 1 {
			monkeys[name] = Job{value: bytesToInt(operationFields[0])}
		} else {
			op1, op2, symbol := string(operationFields[0]), string(operationFields[2]), operationFields[1]
			monkeys[name] = Job{operation: Operation{left: op1, right: op2, symbol: symbol[0]}}
		}
	}

	return monkeys
}

func bytesToInt(b []byte) int {
	c, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return c
}
