package main

import (
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type direction int

const (
	left direction = 1 << iota
	right
	up
	down
)

type tile struct {
	beams    direction
	contents rune
}

func isEnergized(t tile) bool {
	return t.beams > 0
}

func getEnergy(contraption [][]tile, x, y int, d direction) int {
	if x < 0 || x >= len(contraption[0]) || y < 0 || y >= len(contraption) || (contraption[y][x].beams&d) > 0 {
		return 0
	}

	e := 0
	if !isEnergized(contraption[y][x]) {
		e += 1
	}
	contraption[y][x].beams |= d

	switch contraption[y][x].contents {
	case '|':
		if d == left || d == right {
			e += getEnergy(contraption, x, y-1, up)
			e += getEnergy(contraption, x, y+1, down)
			return e
		}
	case '-':
		if d == up || d == down {
			e += getEnergy(contraption, x-1, y, left)
			e += getEnergy(contraption, x+1, y, right)
			return e
		}
	case '/':
		switch d {
		case left:
			d = down
		case right:
			d = up
		case up:
			d = right
		case down:
			d = left
		}
	case '\\':
		switch d {
		case left:
			d = up
		case right:
			d = down
		case up:
			d = left
		case down:
			d = right
		}
	}

	switch d {
	case left:
		x -= 1
	case right:
		x += 1
	case up:
		y -= 1
	case down:
		y += 1
	}

	return e + getEnergy(contraption, x, y, d)
}

func cloneContraption(c [][]tile) [][]tile {
	n := make([][]tile, len(c))
	for i := range c {
		n[i] = slices.Clone(c[i])
	}
	return n
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Advent of Code 2023, Day 16")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("failed to read the input file")
	}

	start := time.Now()

	var contraption [][]tile
	for _, line := range lines {
		var row []tile
		for _, c := range line {
			row = append(row, tile{beams: 0, contents: c})
		}
		contraption = append(contraption, row)
	}

	sum := getEnergy(contraption, 0, 0, right)

	fmt.Println("Part 1:", sum, "tiles got energised")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran in", elapsed)

	start = time.Now()

	contraption = make([][]tile, 0)
	for _, line := range lines {
		var row []tile
		for _, c := range line {
			row = append(row, tile{beams: 0, contents: c})
		}
		contraption = append(contraption, row)
	}

	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption {
			c <- getEnergy(cloneContraption(contraption), 0, i, right)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption {
			c <- getEnergy(cloneContraption(contraption), len(contraption)-1, i, left)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption[0] {
			c <- getEnergy(cloneContraption(contraption), i, 0, down)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range contraption[0] {
			c <- getEnergy(cloneContraption(contraption), i, len(contraption[0])-1, up)
		}
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	sum = 0
	for n := range c {
		sum = max(sum, n)
	}

	fmt.Println("Part 2:", sum, "tiles got energised")
	elapsed = time.Since(start)
	fmt.Println("Part 2 ran in", elapsed)
}
