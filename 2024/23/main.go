package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type graphT map[string][]string

func main() {
	res1, res2 := solve(input)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string) (int, string) {
	graph := parseInput(input)

	res1 := part1(graph)
	res2 := part2(graph)

	return res1, res2
}

func part1(graph graphT) int {
	res1 := 0

	for a, aLinks := range graph {
		for _, b := range aLinks {
			if a >= b {
				continue
			}

			for _, c := range graph[b] {
				if b >= c {
					continue
				}

				if !slices.Contains(graph[c], a) {
					continue
				}

				if a[0] == 't' || b[0] == 't' || c[0] == 't' {
					res1 += 1
				}
			}
		}
	}

	return res1
}

func part2(graph graphT) string {
	maxClique := bronKerbosch(graph)

	sort.Strings(maxClique)

	return strings.Join(maxClique, ",")
}

// Bron–Kerbosch algorithm
func bronKerbosch(graph graphT) []string {
	maxClique := []string{}

	var helper func(r []string, p []string, x []string)
	helper = func(r, p, x []string) {
		if len(p) == 0 && len(x) == 0 {
			if len(r) > len(maxClique) {
				maxClique = r
			}
			return
		}

		pivot := ""
		if len(p) > 0 {
			pivot = p[0]
		} else {
			pivot = x[0]
		}

		for _, v := range difference(p, graph[pivot]) {
			helper(union(r, []string{v}), intersection(p, graph[v]), intersection(x, graph[v]))

			p = slices.DeleteFunc(p, func(el string) bool { return el == v })
			x = slices.DeleteFunc(x, func(el string) bool { return el == v })
		}
	}

	vertices := make([]string, 0, len(graph))
	for k := range graph {
		vertices = append(vertices, k)
	}

	helper([]string{}, vertices, []string{})

	return maxClique
}

func union(a, b []string) []string {
	union := make([]string, 0, len(a)+len(b))
	union = append(union, a...)
	union = append(union, b...)

	return union
}

func intersection(a, b []string) []string {
	intersection := make([]string, 0, min(len(a), len(b)))

	for _, el := range b {
		if slices.Contains(a, el) {
			intersection = append(intersection, el)
		}
	}

	return intersection
}

func difference(a, b []string) []string {
	difference := make([]string, 0, len(a))

	for _, el := range a {
		if !slices.Contains(b, el) {
			difference = append(difference, el)
		}
	}

	return difference
}

func parseInput(input string) graphT {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	graph := graphT{}

	for _, line := range lines {
		computers := strings.Split(line, "-")
		c1, c2 := computers[0], computers[1]

		graph[c1] = append(graph[c1], c2)
		graph[c2] = append(graph[c2], c1)

	}

	return graph
}
