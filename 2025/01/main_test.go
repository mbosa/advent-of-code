package main

import (
	"testing"
)

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestSolveTestInput(t *testing.T) {
	want1 := 3
	want2 := 6

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 1177
	want2 := 6768

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
