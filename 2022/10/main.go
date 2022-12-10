package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", "\n"+part2)
}

func solve(inputFile string) (resPart1 int, resPart2 string) {
	d := newDevice(40, 6)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		d.ApplyInstruction(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	return d.signalStrength, d.drawImage()
}

func newDevice(screenSizeX, screenSizeY int) *device {
	screen := make([][]string, screenSizeY)
	for i := range screen {
		screen[i] = make([]string, screenSizeX)
	}

	return &device{register: 1, screen: screen}
}

type position struct {
	x int
	y int
}

type device struct {
	cycles         int
	register       int
	signalStrength int
	screen         [][]string
	crtPointer     position
}

func (d *device) ApplyInstruction(instruction []byte) {
	d.performCycle()

	if string(instruction) == "noop" {
		return
	}

	d.performCycle()
	value, _ := strconv.Atoi(string(instruction[5:]))
	d.register += value
}
func (d *device) performCycle() {
	d.cycles++

	d.updateSignalStrength()
	d.updateCrtPointer()
	d.drawPixel()
}
func (d *device) updateSignalStrength() {
	if (d.cycles-20)%40 == 0 {
		d.signalStrength += d.cycles * d.register
	}
}
func (d *device) updateCrtPointer() {
	d.crtPointer.y = (d.cycles - 1) / 40
	d.crtPointer.x = (d.cycles - 1) % 40
}
func (d *device) drawPixel() {
	if d.crtPointer.x < d.register-1 || d.crtPointer.x > d.register+1 {
		d.screen[d.crtPointer.y][d.crtPointer.x] = "."
	} else {
		d.screen[d.crtPointer.y][d.crtPointer.x] = "#"
	}
}
func (d *device) drawImage() string {
	image := ""
	for _, line := range d.screen {
		image += strings.Join(line, "") + "\n"
	}

	return image[:len(image)-1]
}
