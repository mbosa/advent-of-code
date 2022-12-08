package main

import (
	"testing"
)

const testInputFile = "input_test.txt"

func TestSolve(t *testing.T) {
	wantPart1, wantPart2 := 21, 8

	gotPart1, gotPart2 := solve(testInputFile)

	if gotPart1 != wantPart1 {
		t.Errorf("part1: want %v, got %v", wantPart1, gotPart1)
	}
	if gotPart2 != wantPart2 {
		t.Errorf("part2: want %v, got %v", wantPart2, gotPart2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(inputFile)
	}
}
