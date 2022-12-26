package main

import (
	"bytes"
	"fmt"
	"os"
)

const inputFile = "input.txt"

var snafuIntMap = map[byte]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}
var intSnafuMap = map[int]byte{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func main() {
	part1 := solve(inputFile)

	fmt.Println("part 1:", part1)
}

func solve(inputFile string) (resPart1 string) {
	rawInput, _ := os.ReadFile(inputFile)
	for _, line := range bytes.Split(rawInput, []byte{'\n'}) {
		resPart1 = addSnafu([]byte(resPart1), line)
	}

	return resPart1
}

func addSnafu(a, b []byte) string {
	n, m := len(a), len(b)

	res := make([]byte, max(n, m))

	remainder := 0
	for i := 0; i < n || i < m; i++ {
		valInt := remainder
		if i < n {
			valInt += snafuIntMap[a[n-1-i]]
		}
		if i < m {
			valInt += snafuIntMap[b[m-1-i]]
		}

		remainder = valInt / 3

		if remainder == 0 {
			res[max(n, m)-1-i] = intSnafuMap[valInt]
		} else {
			if valInt > 0 {
				res[max(n, m)-1-i] = intSnafuMap[valInt-5]
			} else {
				res[max(n, m)-1-i] = intSnafuMap[valInt+5]
			}
		}
	}

	if remainder != 0 {
		res = append([]byte{intSnafuMap[remainder]}, res...)
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
