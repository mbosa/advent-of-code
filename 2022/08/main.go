package main

import (
	"fmt"
	"os"
	"strings"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	input, _ := os.ReadFile(inputFile)
	grid := parseInput(string(input))

	nRows, nCols := len(grid), len(grid[0])

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			currTree := grid[i][j]
			visible, scenicScore := 0, 1
			for k := 1; ; k++ {
				nextJ := j + k
				if nextJ >= nCols {
					visible = 1
					scenicScore *= (k - 1)
					break
				} else if grid[i][nextJ] >= currTree {
					scenicScore *= k
					break
				}
			}
			for k := 1; ; k++ {
				nextJ := j - k
				if nextJ < 0 {
					visible = 1
					scenicScore *= (k - 1)
					break
				} else if grid[i][nextJ] >= currTree {
					scenicScore *= k
					break
				}
			}
			for k := 1; ; k++ {
				nextI := i + k
				if nextI >= nRows {
					visible = 1
					scenicScore *= (k - 1)
					break
				} else if grid[nextI][j] >= currTree {
					scenicScore *= k
					break
				}
			}
			for k := 1; ; k++ {
				nextI := i - k
				if nextI < 0 {
					visible = 1
					scenicScore *= (k - 1)
					break
				} else if grid[nextI][j] >= currTree {
					scenicScore *= k
					break
				}
			}
			resPart1 += visible
			if scenicScore > resPart2 {
				resPart2 = scenicScore
			}
		}
	}
	return resPart1, resPart2
}

func parseInput(input string) [][]int {
	inputLines := strings.Split(input, "\n")
	nRows, nCols := len(inputLines), len(inputLines[0])

	res := make([][]int, nRows)
	for i := range res {
		res[i] = make([]int, nCols)
	}

	for i, row := range inputLines {
		for j, tree := range row {
			res[i][j] = int(tree - '0')
		}
	}

	return res
}
