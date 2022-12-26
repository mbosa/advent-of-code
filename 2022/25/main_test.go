package main

import (
	"testing"
)

const testInputFile = "input_test.txt"

func TestSolve(t *testing.T) {
	want := "2=-1=0"

	got := solve(testInputFile)

	if got != want {
		t.Errorf("part1: want %v, got %v", want, got)
	}
}

func TestSolveInput(t *testing.T) {
	want := "2-=0-=-2=111=220=100"

	got := solve(inputFile)

	if got != want {
		t.Errorf("part1: want %v, got %v", want, got)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(inputFile)
	}
}
