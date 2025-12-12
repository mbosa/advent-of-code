package main

import (
	"testing"
)

const testInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestSolveTestInput(t *testing.T) {
	want1 := 4277556
	want2 := 3263827

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 6635273135233
	want2 := 12542543681221

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
