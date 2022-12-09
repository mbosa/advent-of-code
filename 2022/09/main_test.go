package main

import (
	"testing"
)

const testInputFile = "input_test.txt"
const testInputFile2 = "input_test_2.txt"

func TestSolve(t *testing.T) {
	t.Run("test_input", func(t *testing.T) {
		wantPart1, wantPart2 := 13, 1

		gotPart1, gotPart2 := solve(testInputFile)

		if gotPart1 != wantPart1 {
			t.Errorf("part1: want %v, got %v", wantPart1, gotPart1)
		}
		if gotPart2 != wantPart2 {
			t.Errorf("part2: want %v, got %v", wantPart2, gotPart2)
		}
	})
	t.Run("test_input_2", func(t *testing.T) {
		wantPart1, wantPart2 := 88, 36

		gotPart1, gotPart2 := solve(testInputFile2)

		if gotPart1 != wantPart1 {
			t.Errorf("part1: want %v, got %v", wantPart1, gotPart1)
		}
		if gotPart2 != wantPart2 {
			t.Errorf("part2: want %v, got %v", wantPart2, gotPart2)
		}
	})
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(inputFile)
	}
}
