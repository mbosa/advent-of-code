package main

import (
	"testing"
)

const testInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func TestSolveTestInput(t *testing.T) {
	want1 := 7
	want2 := "co,de,ka,ta"

	got1, got2 := solve(testInput)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %s, got %s", want2, got2)
	}
}

func TestSolveInput(t *testing.T) {
	want1 := 1151
	want2 := "ar,cd,hl,iw,jm,ku,qo,rz,vo,xe,xm,xv,ys"

	got1, got2 := solve(input)

	if got1 != want1 {
		t.Errorf("part1: want %d, got %d", want1, got1)
	}
	if got2 != want2 {
		t.Errorf("part2: want %s, got %s", want2, got2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for range b.N {
		solve(input)
	}
}
