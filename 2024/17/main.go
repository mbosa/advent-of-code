package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type computer struct {
	regA    int
	regB    int
	regC    int
	program []int
	pointer int
}

func (c computer) comboOperand(operand int) int {
	if operand > -1 && operand < 4 {
		return operand
	}
	if operand == 4 {
		return c.regA
	}
	if operand == 5 {
		return c.regB
	}
	if operand == 6 {
		return c.regC
	}

	panic("operand not uspported")
}

func (c *computer) adv(operand int) {
	comboOp := c.comboOperand(operand)

	c.regA >>= comboOp // -> c.regA / 2^comboOp
}

func (c *computer) bxl(operand int) {
	c.regB ^= operand
}

func (c *computer) bst(operand int) {
	comboOp := c.comboOperand(operand)

	c.regB = comboOp % 8
}

func (c *computer) jnz(operand int) bool {
	if c.regA != 0 {
		c.pointer = operand
		return true
	}
	return false
}

func (c *computer) bxc(operand int) {
	_ = operand

	c.regB ^= c.regC
}

func (c *computer) out(operand int) int {
	comboOp := c.comboOperand(operand)

	return comboOp % 8
}

func (c *computer) bdv(operand int) {
	comboOp := c.comboOperand(operand)

	c.regB = c.regA >> comboOp // -> c.regA / 2^comboOp
}

func (c *computer) cdv(operand int) {
	comboOp := c.comboOperand(operand)

	c.regC = c.regA >> comboOp // -> c.regA / 2^comboOp
}

func (c *computer) runOnce() int {
	for c.pointer < len(c.program) {
		instruction := c.program[c.pointer]
		operand := c.program[c.pointer+1]

		switch instruction {
		case 0:
			c.adv(operand)
		case 1:
			c.bxl(operand)
		case 2:
			c.bst(operand)
		case 3:
			if c.jnz(operand) {
				continue
			}
		case 4:
			c.bxc(operand)
		case 5:
			c.pointer += 2
			return c.out(operand)
		case 6:
			c.bdv(operand)
		case 7:
			c.cdv(operand)
		}
		c.pointer += 2

	}
	return -1
}

func (c *computer) runProgram() []int {
	output := []int{}

	for c.pointer < len(c.program) {
		out := c.runOnce()
		if out != -1 {
			output = append(output, out)
		}
	}

	return output
}

func (c *computer) resetPointer() {
	c.pointer = 0
}

func (c *computer) overwriteRegisters(regA, regB, regC int) {
	c.regA, c.regB, c.regC = regA, regB, regC
}

func newComputer(regA, regB, regC int, program []int) computer {
	return computer{regA: regA, regB: regB, regC: regC, program: program}
}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (string, int) {
	computer := parseInput(input)

	res1 := part1(computer)
	res2 := part2(computer)

	return res1, res2
}

func part1(c computer) string {

	output := c.runProgram()

	resBytes := make([]byte, len(output)*2)

	i := 0
	for _, n := range output {
		resBytes[i] = byte('0' + n)
		resBytes[i+1] = byte(',')
		i += 2
	}

	return string(resBytes[:len(resBytes)-1])
}

func part2(c computer) int {
	regA, regB, regC := 0, 0, 0
	program := c.program

	var helper func(regA, i int) int
	helper = func(regA, i int) int {
		if i < 0 {
			return regA
		}

		target := program[i]

		for j := range 8 {
			nextRegA := regA<<3 | j // regA * 8 + j

			c.resetPointer()
			c.overwriteRegisters(nextRegA, regB, regC)
			output := c.runOnce()

			if output == target {
				q := helper(nextRegA, i-1)
				if q != -1 {
					return q
				}
			}
		}

		return -1
	}

	return helper(regA, len(program)-1)
}

func parseInput(input string) computer {
	spl := strings.Split(strings.TrimSpace(input), "\n\n")

	registersStr := spl[0]
	programLine := spl[1]

	regA, regB, regC := 0, 0, 0

	fmt.Sscanf(registersStr, "Register A: %d\nRegister B: %d\nRegister C: %d", &regA, &regB, &regC)

	program := []int{}

	for _, b := range []byte(programLine) {
		if b >= '0' && b <= '9' {
			program = append(program, int(b-'0'))
		}
	}

	return newComputer(regA, regB, regC, program)
}
