package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var operationsMap = map[string]func(int, int) int{
	"+": func(v1, v2 int) int { return v1 + v2 },
	"-": func(v1, v2 int) int { return v1 - v2 },
	"*": func(v1, v2 int) int { return v1 * v2 },
	"/": func(v1, v2 int) int { return v1 / v2 },
}

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	input, _ := os.ReadFile(inputFile)
	monkeysConfig := parseInput(input)

	lcm := 1
	for _, c := range monkeysConfig {
		lcm *= c.throwConfig.testValue
	}

	worryControlFn1 := func(item int) int {
		return item / 3
	}
	worryControlFn2 := func(item int) int {
		return item % lcm
	}

	monkeys1 := newMonkeysGroup(monkeysConfig, worryControlFn1)
	monkeys2 := newMonkeysGroup(monkeysConfig, worryControlFn2)

	for round := 0; round < 20; round++ {
		monkeys1.doRound()
	}
	for round := 0; round < 10000; round++ {
		monkeys2.doRound()
	}

	return monkeys1.calcMonkeysBusiness(), monkeys2.calcMonkeysBusiness()
}

func parseInput(input []byte) []monkeyConfig {
	monkeys := strings.Split((string(input)), "\n\n")
	res := make([]monkeyConfig, len(monkeys))

	for i, monkey := range monkeys {
		lines := strings.Split(monkey, "\n")

		items := []int{}
		for _, item := range strings.Split(lines[1][18:], ", ") {
			items = append(items, mustAtoi(item))
		}

		operationFields := strings.Split(lines[2][19:], " ")
		operation := operationConfig{symbol: operationFields[1], val1: operationFields[0], val2: operationFields[2]}

		testValue := mustAtoi(lines[3][21:])
		throwConfig := throwConfig{
			testValue: testValue,
			ifTrue:    mustAtoi(lines[4][29:]),
			ifFalse:   mustAtoi(lines[5][30:]),
		}

		monkeyConfig := monkeyConfig{
			items:       items,
			operation:   operation,
			throwConfig: throwConfig,
		}

		res[i] = monkeyConfig
	}
	return res
}

func mustAtoi(num string) int {
	c, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return c
}

type monkeyConfig struct {
	items       []int
	operation   operationConfig
	throwConfig throwConfig
}

type operationConfig struct {
	symbol string
	val1   string
	val2   string
}

type throwConfig struct {
	testValue int
	ifTrue    int
	ifFalse   int
}

type monkeysGroup []*monkey

func newMonkeysGroup(config []monkeyConfig, worryControlFn func(int) int) *monkeysGroup {
	group := make(monkeysGroup, len(config))

	for i, monkeyConfig := range config {
		group[i] = newMonkey(monkeyConfig, &group, worryControlFn)
	}

	return &group
}

func (m *monkeysGroup) doRound() {
	for _, m := range *m {
		m.doRound()
	}
}
func (m *monkeysGroup) calcMonkeysBusiness() int {
	max1, max2 := 0, 0

	for _, monkey := range *m {
		if monkey.inspections > max1 {
			max2 = max1
			max1 = monkey.inspections
		} else if monkey.inspections > max2 {
			max2 = monkey.inspections
		}
	}
	return max1 * max2
}

type queue []int

func (q *queue) enqueue(el int) {
	*q = append(*q, el)
}
func (q *queue) dequeue() int {
	el := (*q)[0]
	*q = (*q)[1:]
	return el
}

type monkey struct {
	group        *monkeysGroup
	items        queue
	operation    operationConfig
	throwConfig  throwConfig
	worryControl func(int) int
	inspections  int
}

func newMonkey(config monkeyConfig, group *monkeysGroup, worryControlFn func(int) int) *monkey {
	return &monkey{
		group:        group,
		items:        config.items,
		operation:    config.operation,
		throwConfig:  config.throwConfig,
		worryControl: worryControlFn,
	}
}
func (m *monkey) doRound() {
	for len(m.items) > 0 {
		item := m.items.dequeue()
		m.inspectItem(item)
	}
}
func (m *monkey) inspectItem(item int) {
	m.inspections++
	value := m.do(item)

	if m.test(value) {
		dst := (*m.group)[m.throwConfig.ifTrue]
		m.throw(value, dst)
	} else {
		dst := (*m.group)[m.throwConfig.ifFalse]
		m.throw(value, dst)
	}
}
func (m *monkey) do(item int) int {
	var v1, v2 int
	if m.operation.val1 == "old" {
		v1 = item
	} else {
		v1 = mustAtoi(m.operation.val1)
	}
	if m.operation.val2 == "old" {
		v2 = item
	} else {
		v2 = mustAtoi(m.operation.val2)
	}

	opResult := operationsMap[m.operation.symbol](v1, v2)
	return m.worryControl(opResult)
}
func (m *monkey) test(item int) bool {
	return item%m.throwConfig.testValue == 0
}
func (m *monkey) throw(item int, dst *monkey) {
	dst.addItem(item)
}
func (m *monkey) addItem(item int) {
	m.items.enqueue(item)
}
