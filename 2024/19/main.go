package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type patternsT map[string]struct{}

type designsT []string

const patternMaxLen = 8

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	patterns, designs := parseInput(input)

	res1, res2 := 0, 0

	cache := make(map[string]int, 64)

	for _, design := range designs {
		// DP memo
		clearCache(cache)
		options := countOptionsMemo(design, patterns, cache)

		// DP bottom up
		// options := countOptionsBottomUp(design, patterns)

		if options > 0 {
			res1 += 1
		}

		res2 += options

	}

	return res1, res2
}

func parseInput(input string) (patternsT, designsT) {
	spl := strings.Split(strings.TrimSpace(input), "\n\n")

	patterns := patternsT{}
	for _, p := range strings.Split(spl[0], ", ") {
		patterns[p] = struct{}{}
	}

	designs := strings.Split(spl[1], "\n")

	return patterns, designs
}

// DP memoization
func countOptionsMemo(design string, patterns patternsT, cache map[string]int) int {
	if cached, ok := cache[design]; ok {
		return cached
	}

	res := 0

	if _, ok := patterns[design]; ok {
		res += 1
	}

	for i := 1; i < min(len(design), patternMaxLen+1); i++ {
		pattern := design[:i]
		if _, ok := patterns[pattern]; ok {
			r := countOptionsMemo(design[i:], patterns, cache)

			cache[design[i:]] = r

			res += r
		}
	}

	return res
}

// DP bottom up
// Consider substrings design[i:]
// For each substring, check if it starts with a pattern
// If yes, the number of options is equal of the number of options of design[i+len(pattern):]
// example: design: rrbgbr, patterns: r, wr, b, g, bwu, rb, gb, br
// options per length of the substring
// 6 | 5 | 4 | 3 | 2 | 1 | 0
// 6 | 6 | 3 | 3 | 2 | 1 | 1
// substring of length 0 has 1 option (no pattern)
// substring of length 1 (r) has 1 option: pattern r + options of substring with len 0
// substring of length 2 (br) has 2 options:
// - pattern r + options of substring with len 1 (-> 1)
// - pattern tb + options of substring with len 0 (-> 1)
// -> tot 2 options
// substring of length 3 (gbr) has 3 options:
// - pattern g + options of substring with len 2 (-> 2)
// - pattern gb + options of substring with len 1 (-> 1)
// -> tot 3 options
// etc...
func countOptionsBottomUp(design string, patterns patternsT) int {
	options := make([]int, len(design)+1)
	options[len(design)] = 1

	for i := len(design) - 1; i > -1; i-- {
		for j := i + 1; j < min(i+patternMaxLen+1, len(design)+1); j++ {
			if _, ok := patterns[design[i:j]]; ok {
				options[i] += options[j]
			}
		}
	}

	return options[0]
}

func clearCache(cache map[string]int) {
	for k := range cache {
		delete(cache, k)
	}
}
