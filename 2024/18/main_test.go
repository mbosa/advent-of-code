package main

import (
	"testing"
)

var testInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

const testRows = 7
const testCols = 7
const testPart1Bytes = 12

func TestSolveTestInput(t *testing.T) {
	want1 := 22
	want2 := "6,1"

	got1, got2 := solve(testInput, testRows, testCols, testPart1Bytes)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %s, got %s", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 292
	want2 := "58,44"

	got1, got2 := solve(input, rows, cols, part1Bytes)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %s, got %s", want2, got2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for range b.N {
		solve(input, rows, cols, part1Bytes)
	}
}
