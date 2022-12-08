package main

import (
	_ "embed"

	"testing"
)

//go:embed input_test.txt
var testInput []byte
var testInputStart = []byte("qmgbljsphdztnvjfqwrcgsmlb")
var testInputEnd = []byte("mjqjpqmgbljsphdztnv")

func TestSolve(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		wantPart1, wantPart2 := 7, 19

		gotPart1, gotPart2 := solve(testInput)

		if gotPart1 != wantPart1 {
			t.Errorf("part1: want %v, got %v", wantPart1, gotPart1)
		}
		if gotPart2 != wantPart2 {
			t.Errorf("part2: want %v, got %v", wantPart2, gotPart2)
		}
	})
	t.Run("the first sequence of unique chars is at the start of the input", func(t *testing.T) {
		wantPart1, wantPart2 := 4, 14

		gotPart1, gotPart2 := solve(testInputStart)

		if gotPart1 != wantPart1 {
			t.Errorf("part1: want %v, got %v", wantPart1, gotPart1)
		}
		if gotPart2 != wantPart2 {
			t.Errorf("part2: want %v, got %v", wantPart2, gotPart2)
		}
	})
	t.Run("the first sequence of unique chars is at the end of the input", func(t *testing.T) {
		wantPart1, wantPart2 := 7, 19

		gotPart1, gotPart2 := solve(testInputEnd)

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
		solve(input)
	}
}
