package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 string) {
	crateLines := [][]byte{}
	crateStacks1 := []stack{}
	crateStacks2 := []stack{}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for len(crateStacks1) == 0 && scanner.Scan() {
		line := scanner.Bytes()

		if line[1] == '1' {
			crateStacks1 = parseCrateLines(crateLines)
			crateStacks2 = parseCrateLines(crateLines)
		} else {
			crateLines = append(crateLines, line)
		}
	}

	scanner.Scan() // empty line

	for scanner.Scan() {
		line := scanner.Text()

		instruction := parseInstruction(line)

		applyInstructionPart1(crateStacks1, instruction)
		applyInstructionPart2(crateStacks2, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	for i := 0; i < len(crateStacks1); i++ {
		s1 := crateStacks1[i]
		s2 := crateStacks2[i]
		resPart1 += string(s1.Pop())
		resPart2 += string(s2.Pop())
	}

	return resPart1, resPart2
}

type stack []byte

func (s *stack) Push(item byte) {
	*s = append(*s, item)
}
func (s *stack) PushBulk(items []byte) {
	*s = append(*s, items...)
}
func (s *stack) Pop() byte {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]

	return popped
}
func (s *stack) PopBulk(size int) []byte {
	n := len(*s)
	popped := (*s)[n-size:]
	*s = (*s)[:n-size]

	return popped
}

type instruction struct {
	move int
	from int
	to   int
}

func parseCrateLines(crateLines [][]byte) []stack {
	numOfLines := len(crateLines)
	lenOfLine := len(crateLines[0])
	numOfStacks := (lenOfLine + 1) / 4

	crateStacks := make([]stack, numOfStacks)

	for i := numOfLines - 1; i >= 0; i-- {
		line := crateLines[i]

		for j := 1; j < lenOfLine; j += 4 {
			stackIdx := (j - 1) / 4
			crate := line[j]
			if crate != ' ' {
				crateStacks[stackIdx].Push(crate)
			}
		}
	}

	return crateStacks
}

func parseInstruction(line string) *instruction {
	fields := strings.Split(line, " ")

	return &instruction{
		move: mustAtoi(fields[1]),
		from: mustAtoi(fields[3]) - 1,
		to:   mustAtoi(fields[5]) - 1,
	}
}

func applyInstructionPart1(crateStacks []stack, instruction *instruction) {
	for i := 0; i < instruction.move; i++ {
		crate := crateStacks[instruction.from].Pop()
		crateStacks[instruction.to].Push(crate)
	}
}
func applyInstructionPart2(crateStacks []stack, instruction *instruction) {
	crates := crateStacks[instruction.from].PopBulk(instruction.move)
	crateStacks[instruction.to].PushBulk(crates)
}

func mustAtoi(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}
