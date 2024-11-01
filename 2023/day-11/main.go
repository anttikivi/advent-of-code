package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type point struct {
	x int
	y int
}

func originalPart1(input string) {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	var nonEmptyCols []int

	for _, line := range lines {
		if strings.Contains(line, "#") {
			for j, char := range line {
				if char == '#' {
					nonEmptyCols = append(nonEmptyCols, j)
				}
			}
		}
	}

	var emptyCols []int
	for i := 0; i < len(lines[0]); i++ {
		if !slices.Contains(nonEmptyCols, i) {
			emptyCols = append(emptyCols, i)
		}
	}

	var galaxies []point
	extraRows := 0

	for i, line := range lines {
		extraCols := 0
		var row []rune
		for j, char := range line {
			row = append(row, char)
			if slices.Contains(emptyCols, j) {
				extraCols += 1
				row = append(row, '.')
			}
			if char == '#' {
				galaxies = append(galaxies, point{x: j + extraCols, y: i + extraRows})
			}
		}
		if !slices.Contains(row, '#') {
			extraRows += 1
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1, g2 := galaxies[i], galaxies[j]
			sum += int(math.Abs(float64(g1.x-g2.x)) + math.Abs(float64(g1.y-g2.y)))
		}
	}

	fmt.Println("Part 1: the sum of the lengths between the galaxies is", sum, "in the first implementation of the part 1 solution")
}

func solve(input string, multiplier int) int {
	fmt.Println("Advent of Code 2023, Day 11")

	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	var nonEmptyCols []int

	for _, line := range lines {
		if strings.Contains(line, "#") {
			for j, char := range line {
				if char == '#' {
					nonEmptyCols = append(nonEmptyCols, j)
				}
			}
		}
	}

	var emptyCols []int
	for i := 0; i < len(lines[0]); i++ {
		if !slices.Contains(nonEmptyCols, i) {
			emptyCols = append(emptyCols, i)
		}
	}

	var galaxies []point
	extraRows := 0
	for i, line := range lines {
		extraCols := 0
		if strings.Contains(line, "#") {
			for j, char := range line {
				if char == '#' {
					galaxies = append(galaxies, point{x: j + extraCols, y: i + extraRows})
				}
				if slices.Contains(emptyCols, j) {
					extraCols += multiplier - 1
				}
			}
		} else {
			extraRows += multiplier - 1
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1, g2 := galaxies[i], galaxies[j]
			sum += int(math.Abs(float64(g1.x-g2.x)) + math.Abs(float64(g1.y-g2.y)))
		}
	}

	return sum
}

func main() {
	fmt.Println("Advent of Code 2023, Day 11")

	input := "input.txt"

	originalPart1(input)

	part1 := solve(input, 2)
	part2 := solve(input, 1000000)

	fmt.Println("Part 1: the sum of the lengths between the galaxies is", part1)
	fmt.Println("Part 2: the sum of the lengths between the galaxies is", part2)
}
