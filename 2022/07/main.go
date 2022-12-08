package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input.txt"
const totalDiskSpace = 70000000
const updateSpace = 30000000
const smallDirSizeLimit = 100000

func main() {
	part1, part2 := solve(inputFile)

	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func solve(inputFile string) (resPart1, resPart2 int) {
	fsys := newFilesystem()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		fsys.ingestTerminalOutput(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %v: %v", inputFile, err)
	}

	rootDirSize := fsys.dirSize["//"]
	freeSpace := totalDiskSpace - rootDirSize
	sumOfSmallDir, smallestDeletableDirSize := 0, rootDirSize
	for _, size := range fsys.dirSize {
		if size <= smallDirSizeLimit {
			sumOfSmallDir += size
		}

		if freeSpace+size >= updateSpace {
			if size < smallestDeletableDirSize {
				smallestDeletableDirSize = size
			}
		}
	}

	return sumOfSmallDir, smallestDeletableDirSize
}

func indexOf(sl []byte, char byte) int {
	for i, v := range sl {
		if v == char {
			return i
		}
	}
	return -1
}

type stack []string

func (s *stack) Push(item string) {
	*s = append(*s, item)
}
func (s *stack) Pop() string {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]

	return popped
}
func (s *stack) Peek() string {
	n := len(*s)
	if n == 0 {
		return ""
	}

	return (*s)[n-1]
}
func (s *stack) Items() []string {
	return *s
}

func newFilesystem() *filesystem {
	return &filesystem{
		dirStack: stack{},
		dirSize:  make(map[string]int),
	}
}

type filesystem struct {
	dirStack stack
	dirSize  map[string]int
}

func (fs *filesystem) ingestTerminalOutput(line terminalOutput) {
	if dir, isCd := line.cdCommand(); isCd {
		fs.applyCd(dir)
	} else if size, isFile := line.lsFileOutput(); isFile {
		fs.updateDirSize(size)
	}
}
func (fs *filesystem) applyCd(dir string) {
	if dir == ".." {
		fs.dirStack.Pop()
	} else {
		currentDir := fs.dirStack.Peek()
		nextDir := currentDir + dir + "/"
		fs.dirStack.Push(nextDir)
	}
}
func (fs *filesystem) updateDirSize(size int) {
	openDirs := fs.dirStack.Items()
	for _, dir := range openDirs {
		fs.dirSize[dir] += size
	}
}

type terminalOutput []byte

func (t terminalOutput) cdCommand() (dir string, isCd bool) {
	if t[0] != '$' || t[2] != 'c' {
		return "", false
	}
	return string(t[5:]), true
}
func (t terminalOutput) lsFileOutput() (size int, isFile bool) {
	if t[0] < '0' || t[0] > '9' {
		return 0, false
	}

	sizeB := t[:indexOf(t, ' ')]

	size, _ = strconv.Atoi(string(sizeB))
	return size, true
}
