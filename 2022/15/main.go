package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile, 2000000, 4000000)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string, rowPart1, maxCoordPart2 int) (resPart1, resPart2 int) {
	rawInput, _ := os.ReadFile(inputFile)
	input := parse(rawInput)

	return part1(input, rowPart1), part2(input, 0, maxCoordPart2)
}

func part1(pairs []Pair, row int) int {
	intersections := [][2]int{}

	beaconsOnRow := map[Position]bool{}

	for _, p := range pairs {
		distanceSensorRow := abs(p.Sensor.Y - row)
		if distanceSensorRow > p.Distance {
			continue
		}

		if p.Beacon.Y == row {
			beaconsOnRow[p.Beacon] = true
		}

		diff := p.Distance - distanceSensorRow
		intersection := [2]int{p.Sensor.X - diff, p.Sensor.X + diff}

		intersections = append(intersections, intersection)
	}

	intersectionsUnion := intervalsUnion(intersections)

	res := 0
	for _, intercection := range intersectionsUnion {
		res += intercection[1] - intercection[0] + 1
	}

	return res - len(beaconsOnRow)
}

func part2(pairs []Pair, minCoord, maxCoord int) int {
	min, max := Position{minCoord, minCoord}, Position{maxCoord, maxCoord}
	quadrant := Quadrant{BottomLeft: min, TopRight: max}
	foundPos := findUnseenPos(pairs, quadrant)

	return foundPos.X*4000000 + foundPos.Y
}

func findUnseenPos(pairs []Pair, quadrant Quadrant) Position {
	quadrantStack := Stack{quadrant}

	for len(quadrantStack) > 0 {
		quadrant := quadrantStack.Pop()

		if quadrant.BottomLeft == quadrant.TopRight {
			return quadrant.BottomLeft
		}

		midX, midY := (quadrant.BottomLeft.X+quadrant.TopRight.X)/2, (quadrant.BottomLeft.Y+quadrant.TopRight.Y)/2
		mid := Position{X: midX, Y: midY}
		subQuadrants := []Quadrant{
			{BottomLeft: quadrant.BottomLeft, TopRight: mid},
			{BottomLeft: Position{X: mid.X + 1, Y: quadrant.BottomLeft.Y}, TopRight: Position{X: quadrant.TopRight.X, Y: mid.Y}},
			{BottomLeft: Position{X: quadrant.BottomLeft.X, Y: mid.Y + 1}, TopRight: Position{X: mid.X, Y: quadrant.TopRight.Y}},
			{BottomLeft: Position{X: mid.X + 1, Y: mid.Y + 1}, TopRight: quadrant.TopRight},
		}

		for _, q := range subQuadrants {
			if q.BottomLeft.X > q.TopRight.X || q.BottomLeft.Y > q.TopRight.Y {
				continue
			}

			maybeTargetInQuadrant := true
			for _, p := range pairs {
				if p.IsAreaInSensorRange(q.BottomLeft, q.TopRight) {
					maybeTargetInQuadrant = false
					break
				}
			}
			if maybeTargetInQuadrant {
				quadrantStack.Push(q)
			}
		}
	}

	return Position{-1, -1}
}

type Position struct {
	X int
	Y int
}

type Pair struct {
	Sensor   Position
	Beacon   Position
	Distance int
}

func (p *Pair) IsAreaInSensorRange(bottomLeft, topRight Position) bool {
	corners := []Position{
		{X: bottomLeft.X, Y: bottomLeft.Y},
		{X: bottomLeft.X, Y: topRight.Y},
		{X: topRight.X, Y: topRight.Y},
		{X: topRight.X, Y: bottomLeft.Y},
	}
	for _, c := range corners {
		d := manhattanDistance(p.Sensor, c)

		if d > p.Distance {
			return false
		}
	}

	return true
}

type Quadrant struct {
	BottomLeft Position
	TopRight   Position
}

type Stack []Quadrant

func (s *Stack) Push(item Quadrant) {
	*s = append(*s, item)
}
func (s *Stack) Pop() Quadrant {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]

	return popped
}

func parse(raw []byte) []Pair {
	parsed := []Pair{}

	for _, line := range bytes.Split(raw, []byte{'\n'}) {
		split := bytes.Split(line, []byte{' '})

		sensorXBytes := split[2][2 : len(split[2])-1]
		sensorYBytes := split[3][2 : len(split[3])-1]
		beaconXBytes := split[8][2 : len(split[8])-1]
		beaconYBytes := split[9][2:]

		sensor := Position{X: bytesToInt(sensorXBytes), Y: bytesToInt(sensorYBytes)}
		beacon := Position{X: bytesToInt(beaconXBytes), Y: bytesToInt(beaconYBytes)}
		distance := manhattanDistance(sensor, beacon)

		parsed = append(parsed, Pair{Sensor: sensor, Beacon: beacon, Distance: distance})
	}

	return parsed
}

func bytesToInt(b []byte) int {
	c, _ := strconv.Atoi(string(b))
	return c
}

func manhattanDistance(p1, p2 Position) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func intervalsUnion(intervals [][2]int) [][2]int {
	if len(intervals) < 2 {
		return intervals
	}

	union := [][2]int{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	start, end := intervals[0][0], intervals[0][1]

	for _, interval := range intervals {
		if interval[0] > end+1 {
			union = append(union, [2]int{start, end})
			start, end = interval[0], interval[1]
		}
		if interval[0] >= start && interval[1] > end {
			end = interval[1]
		}
	}
	union = append(union, [2]int{start, end})

	return union
}
