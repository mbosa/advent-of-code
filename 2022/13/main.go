package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var packets [][]any
	var prev []any

	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			i++
			continue
		}

		parsed := parsePacket(line)
		packets = append(packets, parsed)

		if i%3 == 0 {
			prev = parsed
		} else {
			if compare(prev, parsed) < 0 {
				resPart1 += i/3 + 1
			}
		}
		i++
	}

	resPart2 = part2(packets)

	return resPart1, resPart2
}

func parsePacket(packet []byte) []any {
	var parsed []any

	err := json.Unmarshal(packet, &parsed)
	if err != nil {
		panic(err)
	}
	return parsed
}

// < 0: a before b; > 0: b before a; == 0: a equals b
func compare(a any, b any) int {
	aSlice, aOk := a.([]any)
	bSlice, bOk := b.([]any)

	if !aOk && !bOk {
		return int(a.(float64) - b.(float64))
	}
	if !aOk {
		return compare([]any{a}, bSlice)
	}
	if !bOk {
		return compare(aSlice, []any{b})
	}

	for i := 0; i < min(len(aSlice), len(bSlice)); i++ {
		c := compare(aSlice[i], bSlice[i])
		if c == 0 {
			continue
		}
		return c
	}

	return len(aSlice) - len(bSlice)
}

func part2(packets [][]any) int {
	res := 1

	dividers := [][]any{{[]any{2.0}}, {[]any{6.0}}}
	packets = append(packets, dividers...)

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	for i, p := range packets {
		marshalled, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}

		s := string(marshalled)
		if s == "[[2]]" || s == "[[6]]" {
			res *= (i + 1)
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
