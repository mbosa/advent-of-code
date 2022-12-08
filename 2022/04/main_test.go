package main

import (
	"reflect"
	"testing"
)

const testInputFile = "input_test.txt"

func TestParseLine(t *testing.T) {
	wantPair1, wantPair2 := [2]int{2, 4}, [2]int{6, 8}

	gotPair1, gotPair2 := parseLine("2-4,6-8")

	if !reflect.DeepEqual(gotPair1, wantPair1) {
		t.Errorf("error parsing line: want %d, got %d", wantPair1, gotPair1)
	}
	if !reflect.DeepEqual(gotPair2, wantPair2) {
		t.Errorf("error parsing line: want %d, got %d", wantPair2, gotPair2)
	}
}

func TestFullyContained(t *testing.T) {
	testCases := []struct {
		pair1 [2]int
		pair2 [2]int
		want  bool
	}{
		{pair1: [2]int{2, 8}, pair2: [2]int{3, 7}, want: true},
		{pair1: [2]int{4, 6}, pair2: [2]int{6, 6}, want: true},
		{pair1: [2]int{2, 3}, pair2: [2]int{4, 5}, want: false},
		{pair1: [2]int{2, 4}, pair2: [2]int{3, 5}, want: false},
		{pair1: [2]int{3, 5}, pair2: [2]int{2, 4}, want: false},
	}

	for _, test := range testCases {
		got := fullyContained(test.pair1, test.pair2)
		if got != test.want {
			t.Errorf("fullyContained %v in %v: want %v, got %v", test.pair2, test.pair1, test.want, got)
		}
	}
}

func TestOverlap(t *testing.T) {
	testCases := []struct {
		pair1 [2]int
		pair2 [2]int
		want  bool
	}{
		{pair1: [2]int{2, 8}, pair2: [2]int{3, 7}, want: true},
		{pair1: [2]int{4, 6}, pair2: [2]int{6, 6}, want: true},
		{pair1: [2]int{2, 4}, pair2: [2]int{3, 5}, want: true},
		{pair1: [2]int{3, 5}, pair2: [2]int{2, 4}, want: true},
		{pair1: [2]int{5, 7}, pair2: [2]int{7, 9}, want: true},
		{pair1: [2]int{2, 3}, pair2: [2]int{4, 5}, want: false},
	}

	for _, test := range testCases {
		got := overlap(test.pair1, test.pair2)
		if got != test.want {
			t.Errorf("overlap %v with %v: want %v, got %v", test.pair2, test.pair1, test.want, got)
		}
	}
}

func TestSolve(t *testing.T) {
	wantPart1, wantPart2 := 2, 4

	gotPart1, gotPart2 := solve(testInputFile)

	if gotPart1 != wantPart1 {
		t.Errorf("part1: want %d, got %d", wantPart1, gotPart1)
	}
	if gotPart2 != wantPart2 {
		t.Errorf("part2: want %d, got %d", wantPart2, gotPart2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(inputFile)
	}
}
