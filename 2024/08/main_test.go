package main

import (
	"testing"
)

var testInput = []byte(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

func TestSolveTestInput(t *testing.T) {
	want1 := 14
	want2 := 34

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 341
	want2 := 1134

	got1, got2 := solve(input)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for range b.N {
		solve(input)
	}
}
