package main

import (
	"testing"
)

const testInput1 = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

const testInput2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestSolveTestInput1(t *testing.T) {
	want1 := "4,6,3,5,6,3,5,2,1,0"
	want2 := 29328

	got1, got2 := solve(testInput1)

	if got1 != want1 {
		t.Errorf("part1: want %s, got %s", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveTestInput2(t *testing.T) {
	want1 := "5,7,3,0"
	want2 := 117440

	got1, got2 := solve(testInput2)

	if got1 != want1 {
		t.Errorf("part1: want %s, got %s", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %d, got %d", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := "1,6,3,6,5,6,5,1,7"
	want2 := 247839653009594

	got1, got2 := solve(input)

	if got1 != want1 {
		t.Errorf("part1: want %s, got %s", want1, got1)
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
