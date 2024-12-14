package main

import (
	"fmt"
	"math"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type position struct {
	x int
	y int
}
type velocity struct {
	dx int
	dy int
}
type robot struct {
	pos position
	vel velocity
}

var width = 101
var height = 103

func main() {
	res1, res2 := solve(input, width, height)

	fmt.Println("part1:", res1)
	fmt.Println("part2:", res2)
}

func solve(input string, width, height int) (int, int) {
	robots := parseInput(input)

	res1 := part1(robots, width, height)
	res2 := part2(robots, width, height)

	return res1, res2
}

func part1(robots []robot, width, height int) int {
	q1Count, q2Count, q3Count, q4Count := 0, 0, 0, 0

	// axis that split the grid in 4 quadrants
	xAxisIdx := width / 2
	yAxisIdx := height / 2

	for _, robot := range robots {
		newPosX := (((robot.pos.x + robot.vel.dx*100) % width) + width) % width
		newPosY := (((robot.pos.y + robot.vel.dy*100) % height) + height) % height

		if newPosX < xAxisIdx && newPosY < yAxisIdx {
			q1Count += 1
		}
		if newPosX > xAxisIdx && newPosY < yAxisIdx {
			q2Count += 1
		}
		if newPosX < xAxisIdx && newPosY > yAxisIdx {
			q3Count += 1
		}
		if newPosX > xAxisIdx && newPosY > yAxisIdx {
			q4Count += 1
		}
	}

	return q1Count * q2Count * q3Count * q4Count
}

/*
The movement of the robots is cyclical.
The x position repeats every `width` ticks, and the y position repeats every `height` ticks.
The position repeats every `width*height` ticks.

Assumption:
For the robots to display a christmas tree, many of them will be close together.
It will happen when both x and y have the lowest variance.
variance = SUM((value - mean)^2) / n

Strategy:
  - Find the ticks tick_min_variance_x, where variance_x is the lowest during the first cycle_x,
    and tick_min_variance_y, where variance_y is the lowest during the first cycle_y
  - Find a tick where both x and y have their lowest variance.
    Since both x and y are cyclical, the target_tick will happen when
    target_tick = tick_min_variance_x + m*width, and target_tick = tick_min_variance_y + n*height
    There is some math magic to find the result, but I didn't understand it,
    so I can just add cycles to tick_min_variance_x and tick_min_variance_y
    until they are equal (=> both min variances happen on the same tick)
*/
func part2(robotsInput []robot, width, height int) int {
	// copy the input since I will mutate it
	robots := make([]robot, len(robotsInput))
	copy(robots, robotsInput)

	lenRobots := float64(len(robots))
	loops := max(width, height)

	minVarianceX := math.MaxFloat64
	minVarianceY := math.MaxFloat64
	ticksToMinVarianceX := 0
	ticksToMinVarianceY := 0

	for tick := 1; tick < loops; tick++ {
		meanX, meanY := 0.0, 0.0
		varianceX, varianceY := 0.0, 0.0

		for i := range robots {
			robots[i].pos.x = (((robots[i].pos.x + robots[i].vel.dx) % width) + width) % width
			robots[i].pos.y = (((robots[i].pos.y + robots[i].vel.dy) % height) + height) % height

			meanX += float64(robots[i].pos.x)
			meanY += float64(robots[i].pos.y)
		}

		meanX /= lenRobots
		meanY /= lenRobots

		for _, robot := range robots {
			dx := float64(robot.pos.x) - meanX
			dy := float64(robot.pos.y) - meanY

			squareDiffX := dx * dx
			squareDiffY := dy * dy

			varianceX += squareDiffX
			varianceY += squareDiffY
		}
		varianceX /= lenRobots
		varianceY /= lenRobots

		if varianceX <= minVarianceX {
			minVarianceX = varianceX

			ticksToMinVarianceX = tick
		}

		if varianceY <= minVarianceY {
			minVarianceY = varianceY

			ticksToMinVarianceY = tick
		}
	}

	// add cycles until I find a tick where both min variances happen
	for range loops * 2 {
		if ticksToMinVarianceX < ticksToMinVarianceY {
			ticksToMinVarianceX += width
		} else if ticksToMinVarianceY < ticksToMinVarianceX {
			ticksToMinVarianceY += height
		}

		if ticksToMinVarianceX == ticksToMinVarianceY {
			return ticksToMinVarianceX
		}
	}

	return 0
}

func parseInput(input string) []robot {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	robots := make([]robot, len(lines))

	for i, line := range lines {
		pos := position{}
		vel := velocity{}

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pos.x, &pos.y, &vel.dx, &vel.dy)

		r := robot{pos, vel}
		robots[i] = r
	}
	return robots
}
