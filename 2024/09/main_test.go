package main

import (
	"testing"
)

var testInput = []byte("2333133121414131402")

func TestSolveTestInput(t *testing.T) {
	want1 := 1928
	want2 := 2858

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 6320029754031
	want2 := 6347435485773

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
