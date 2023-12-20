package main

import (
	"fmt"
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
}
