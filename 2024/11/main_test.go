package main

import (
	"testing"
)

const testInput = "125 17"

func TestSolveTestInput(t *testing.T) {
	want1 := 55312
	want2 := 65601038650482

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 229043
	want2 := 272673043446478

	got1, got2 := solve(input)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveMemoInput(t *testing.T) {
	want1 := 229043
	want2 := 272673043446478

	got1, got2 := solveMemo(input)

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

func BenchmarkSolveMemo(b *testing.B) {
	for range b.N {
		solveMemo(input)
	}
}
