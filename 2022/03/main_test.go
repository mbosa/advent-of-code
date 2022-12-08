package main

import (
	"reflect"
	"testing"
)

const testInputFile = "input_test.txt"

func TestCommonChars(t *testing.T) {
	testCases := []struct {
		chars1 []byte
		chars2 []byte
		want   []byte
	}{
		{chars1: []byte("abcde"), chars2: []byte("char"), want: []byte("ca")},
		{chars1: []byte("abCde"), chars2: []byte("CHAR"), want: []byte("C")},
		{chars1: []byte("aBcaBc"), chars2: []byte("aBcdaBcd"), want: []byte("aBc")},
		{chars1: []byte("abc"), chars2: []byte("def"), want: []byte("")},
	}

	for _, test := range testCases {
		got := commonChars(test.chars1, test.chars2)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("common chars between %v and %v: want %v, got %v", string(test.chars2), string(test.chars1), string(test.want), string(got))
		}
	}
}

func TestSolve(t *testing.T) {
	wantPart1, wantPart2 := 157, 70

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
