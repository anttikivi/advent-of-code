package main

import (
	"fmt"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type space int

const (
	round space = iota
	cube
	empty
)

func main() {
	fmt.Println("Advent of Code 2023, Day 14")

	input := "input.txt"

	lines, err := utils.ReadLines(input)
	if err != nil {
		panic("failed to read the input")
	}

	platform := make([][]space, 0, len(lines))
	for _, line := range lines {
		row := make([]space, 0, len(line))
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == 'O' {
				row = append(row, round)
			} else if c == '#' {
				row = append(row, cube)
			} else if c == '.' {
				row = append(row, empty)
			} else {
				panic("invalid character")
			}
		}
		platform = append(platform, row)
	}

	load := 0
	rows := len(platform)
	for x := range platform[0] {
		empties := 0
		for i := 0; i < len(platform); i++ {
			s := platform[i][x]
			if s == empty {
				empties += 1
			} else if s == round {
				load += rows - i + empties
			} else if s == cube {
				empties = 0
			} else {
				panic("invalid space")
			}
		}
	}
	fmt.Println("Part 1: the total load on the north support beams is", load)
}
