package main

import (
	"testing"
)

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestSolveTestInput(t *testing.T) {
	want1 := 357
	want2 := 3121910778619

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 17095
	want2 := 168794698570517

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
