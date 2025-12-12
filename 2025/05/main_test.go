package main

import (
	"testing"
)

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestSolveTestInput(t *testing.T) {
	want1 := 3
	want2 := 14

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 761
	want2 := 345755049374932

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
