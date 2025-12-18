package main

import (
	"testing"
)

const testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestSolveTestInput(t *testing.T) {
	want1 := 50
	want2 := 24

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 4786902990
	want2 := 1571016172

	got1, got2 := solve(input)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		solve(input)
	}
}
