package main

import (
	"fmt"
	"math"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type cacheKeyT struct {
	seq   string
	depth int
}

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, int) {
	cache := make(map[cacheKeyT]int)

	codes := parseInput(input)

	res1 := part1(codes, cache)
	res2 := part2(codes, cache)

	return res1, res2
}

func part1(codes []string, cache map[cacheKeyT]int) int {
	res := 0

	for _, code := range codes {
		l := expandSeqNum(code, 3, cache)

		score := l * codeToNum(code)
		res += score
	}

	return res
}

func part2(codes []string, cache map[cacheKeyT]int) int {
	res := 0

	for _, code := range codes {
		l := expandSeqNum(code, 26, cache)

		score := l * codeToNum(code)
		res += score
	}

	return res
}

// return the lenght of expanding a numeric sequence for `depth` times
func expandSeqNum(seq string, depth int, cache map[cacheKeyT]int) int {
	res := 0

	var from byte = 'A'

	for i := 0; i < len(seq); i += 1 {
		to := seq[i]

		res += expandStepNum(from, to, depth, cache)

		from = to
	}

	return res
}

// return the length of expanding a step on the numeric pad from `from` to `to` for `depth` times
// expanding a step on the numeric pad results in one of more directional sequences
func expandStepNum(from, to byte, depth int, cache map[cacheKeyT]int) int {
	res := math.MaxInt

	paths := pathToBtnNum(from, to)

	for _, path := range paths {
		l := expandSeqDir(path, depth-1, cache)

		if l < res {
			res = l
		}
	}

	return res
}

// return the lenght of expanding a directional sequence for `depth` times
func expandSeqDir(seq string, depth int, cache map[cacheKeyT]int) int {
	if depth == 0 {
		return len(seq)
	}

	cacheKey := cacheKeyT{seq, depth}

	if r, ok := cache[cacheKey]; ok {
		return r
	}

	res := 0

	var from byte = 'A'

	for i := 0; i < len(seq); i += 1 {
		to := seq[i]

		res += expandStepDir(from, to, depth, cache)

		from = to
	}

	cache[cacheKey] = res

	return res
}

// return the length of expanding a step on the directional pad from `from` to `to` for `depth` times
func expandStepDir(from, to byte, depth int, cache map[cacheKeyT]int) int {
	res := math.MaxInt

	paths := pathToBtnDir(from, to)

	for _, path := range paths {
		l := expandSeqDir(path, depth-1, cache)

		if l < res {
			res = l
		}
	}

	return res
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")

}

func codeToNum(code string) int {
	return int(code[0]-'0')*100 + int(code[1]-'0')*10 + int(code[2]-'0')
}

func pathToBtnDir(from, to byte) []string {
	switch from {
	case 'A':
		switch to {
		case 'A':
			return []string{"A"}
		case '^':
			return []string{"<A"}
		case '>':
			return []string{"vA"}
		case 'v':
			return []string{"<vA", "v<A"}
		case '<':
			return []string{"v<<A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '^':
		switch to {
		case 'A':
			return []string{">A"}
		case '^':
			return []string{"A"}
		case '>':
			return []string{">vA", "v>A"}
		case 'v':
			return []string{"vA"}
		case '<':
			return []string{"v<A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '>':
		switch to {
		case 'A':
			return []string{"^A"}
		case '^':
			return []string{"<^A", "^<A"}
		case '>':
			return []string{"A"}
		case 'v':
			return []string{"<A"}
		case '<':
			return []string{"<<A"}
		default:
			panic("invalid to: " + string(to))
		}
	case 'v':
		switch to {
		case 'A':
			return []string{">^A", "^>A"}
		case '^':
			return []string{"^A"}
		case '>':
			return []string{">A"}
		case 'v':
			return []string{"A"}
		case '<':
			return []string{"<A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '<':
		switch to {
		case 'A':
			return []string{">>^A"}
		case '^':
			return []string{">^A"}
		case '>':
			return []string{">>A"}
		case 'v':
			return []string{">A"}
		case '<':
			return []string{"A"}
		default:
			panic("invalid to: " + string(to))
		}
	default:
		panic("invalid from: " + string(from))
	}
}

func pathToBtnNum(from, to byte) []string {
	switch from {
	case 'A':
		switch to {
		case 'A':
			return []string{"A"}
		case '0':
			return []string{"<A"}
		case '1':
			return []string{"^<<A"}
		case '2':
			return []string{"<^A", "^<A"}
		case '3':
			return []string{"^A"}
		case '4':
			return []string{"^^<<A"}
		case '5':
			return []string{"<^^A", "^^<A"}
		case '6':
			return []string{"^^A"}
		case '7':
			return []string{"^^^<<A"}
		case '8':
			return []string{"<^^^A", "^^^<A"}
		case '9':
			return []string{"^^^A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '0':
		switch to {
		case 'A':
			return []string{">A"}
		case '0':
			return []string{"A"}
		case '1':
			return []string{"^<A"}
		case '2':
			return []string{"^A"}
		case '3':
			return []string{">^A", "^>A"}
		case '4':
			return []string{"^^<A"}
		case '5':
			return []string{"^^A"}
		case '6':
			return []string{">^^A", "^^>A"}
		case '7':
			return []string{"^^^<A"}
		case '8':
			return []string{"^^^A"}
		case '9':
			return []string{">^^^A", "^^^>A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '1':
		switch to {
		case 'A':
			return []string{">>vA"}
		case '0':
			return []string{">vA"}
		case '1':
			return []string{"A"}
		case '2':
			return []string{">A"}
		case '3':
			return []string{">>A"}
		case '4':
			return []string{"^A"}
		case '5':
			return []string{">^A", "^>A"}
		case '6':
			return []string{">>^A", "^>>A"}
		case '7':
			return []string{"^^A"}
		case '8':
			return []string{">^^A", "^^>A"}
		case '9':
			return []string{">>^^A", "^^>>A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '2':
		switch to {
		case 'A':
			return []string{">vA", "v>A"}
		case '0':
			return []string{"vA"}
		case '1':
			return []string{"<A"}
		case '2':
			return []string{"A"}
		case '3':
			return []string{">A"}
		case '4':
			return []string{"<^A", "^<A"}
		case '5':
			return []string{"^A"}
		case '6':
			return []string{">^A", "^>A"}
		case '7':
			return []string{"<^^A", "^^<A"}
		case '8':
			return []string{"^^A"}
		case '9':
			return []string{">^^A", "^^>A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '3':
		switch to {
		case 'A':
			return []string{"vA"}
		case '0':
			return []string{"<vA", "v<A"}
		case '1':
			return []string{"<<A"}
		case '2':
			return []string{"<A"}
		case '3':
			return []string{"A"}
		case '4':
			return []string{"<<^A", "^<<A"}
		case '5':
			return []string{"<^A", "^<A"}
		case '6':
			return []string{"^A"}
		case '7':
			return []string{"<<^^A", "^^<<A"}
		case '8':
			return []string{"<^^A", "^^<A"}
		case '9':
			return []string{"^^A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '4':
		switch to {
		case 'A':
			return []string{">>vvA"}
		case '0':
			return []string{">vvA"}
		case '1':
			return []string{"vA"}
		case '2':
			return []string{">vA", "v>A"}
		case '3':
			return []string{">>vA", "v>>A"}
		case '4':
			return []string{"A"}
		case '5':
			return []string{">A"}
		case '6':
			return []string{">>A"}
		case '7':
			return []string{"^A"}
		case '8':
			return []string{">^A", "^>A"}
		case '9':
			return []string{">>^A", "^>>A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '5':
		switch to {
		case 'A':
			return []string{">vvA", "vv>A"}
		case '0':
			return []string{"vvA"}
		case '1':
			return []string{"<vA", "v<A"}
		case '2':
			return []string{"vA"}
		case '3':
			return []string{">vA", "v>A"}
		case '4':
			return []string{"<A"}
		case '5':
			return []string{"A"}
		case '6':
			return []string{">A"}
		case '7':
			return []string{"<^A", "^<A"}
		case '8':
			return []string{"^A"}
		case '9':
			return []string{">^A", "^>A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '6':
		switch to {
		case 'A':
			return []string{"vvA"}
		case '0':
			return []string{"<vvA", "vv<A"}
		case '1':
			return []string{"<<vA", "v<<A"}
		case '2':
			return []string{"<vA", "v<A"}
		case '3':
			return []string{"vA"}
		case '4':
			return []string{"<<A"}
		case '5':
			return []string{"<A"}
		case '6':
			return []string{"A"}
		case '7':
			return []string{"<<^A", "^<<A"}
		case '8':
			return []string{"<^A", "^<A"}
		case '9':
			return []string{"^A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '7':
		switch to {
		case 'A':
			return []string{">>vvvA"}
		case '0':
			return []string{">vvvA"}
		case '1':
			return []string{"vvA"}
		case '2':
			return []string{">vvA", "vv>A"}
		case '3':
			return []string{">>vvA", "vv>>A"}
		case '4':
			return []string{"v"}
		case '5':
			return []string{">vA", "v>A"}
		case '6':
			return []string{">>vA", "v>>A"}
		case '7':
			return []string{"A"}
		case '8':
			return []string{">A"}
		case '9':
			return []string{">>A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '8':
		switch to {
		case 'A':
			return []string{">vvvA", "vvv>A"}
		case '0':
			return []string{"vvvA"}
		case '1':
			return []string{"<vvA", "vv<A"}
		case '2':
			return []string{"vvA"}
		case '3':
			return []string{">vvA", "vv>A"}
		case '4':
			return []string{"<vA", "v<A"}
		case '5':
			return []string{"vA"}
		case '6':
			return []string{">vA", "v>A"}
		case '7':
			return []string{"<"}
		case '8':
			return []string{"A"}
		case '9':
			return []string{">A"}
		default:
			panic("invalid to: " + string(to))
		}
	case '9':
		switch to {
		case 'A':
			return []string{"vvvA"}
		case '0':
			return []string{"<vvvA", "vvv<A"}
		case '1':
			return []string{"<<vvA", "vv<<A"}
		case '2':
			return []string{"<vvA", "vv<A"}
		case '3':
			return []string{"vvA"}
		case '4':
			return []string{"<<vA", "v<<A"}
		case '5':
			return []string{"<vA", "v<A"}
		case '6':
			return []string{"vA"}
		case '7':
			return []string{"<<A"}
		case '8':
			return []string{"<A"}
		case '9':
			return []string{"A"}
		default:
			panic("invalid to: " + string(to))
		}
	default:
		panic("invalid from: " + string(from))
	}
}
