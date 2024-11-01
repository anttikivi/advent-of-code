package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type space int

const (
	round space = iota
	cube
	empty
)

func tiltNorth(p *[][]string) {
	for row := 1; row < len(*p); row++ {
		for col := 0; col < len((*p)[0]); col++ {
			if (*p)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := row; i > 0; i-- {
					if (*p)[i-1][col] == "O" || (*p)[i-1][col] == "#" {
						// Hit rock
						(*p)[i][col] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*p)[0][col] = "O"
				}

				if rockPlaced != row {
					(*p)[row][col] = "."
				}
			}
		}
	}
}

func tiltWest(p *[][]string) {
	for col := 1; col < len((*p)[0]); col++ {
		for row := 0; row < len(*p); row++ {
			if (*p)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := col; i > 0; i-- {
					if (*p)[row][i-1] == "O" || (*p)[row][i-1] == "#" {
						// Hit rock
						(*p)[row][i] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*p)[row][0] = "O"
				}

				if rockPlaced != col {
					(*p)[row][col] = "."
				}
			}
		}
	}
}

func tiltSouth(p *[][]string) {
	for row := len(*p) - 2; row >= 0; row-- {
		for col := 0; col < len((*p)[0]); col++ {
			if (*p)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := row; i < len(*p)-1; i++ {
					if (*p)[i+1][col] == "O" || (*p)[i+1][col] == "#" {
						// Hit rock
						(*p)[i][col] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*p)[len(*p)-1][col] = "O"
				}

				if rockPlaced != row {
					(*p)[row][col] = "."
				}
			}
		}
	}
}

func tiltEast(p *[][]string) {
	for col := len((*p)[0]) - 2; col >= 0; col-- {
		for row := 0; row < len(*p); row++ {
			if (*p)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := col; i < len((*p)[0])-1; i++ {
					if (*p)[row][i+1] == "O" || (*p)[row][i+1] == "#" {
						// Hit rock
						(*p)[row][i] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*p)[row][len((*p)[0])-1] = "O"
				}

				if rockPlaced != col {
					(*p)[row][col] = "."
				}
			}
		}
	}
}

func main() {
	fmt.Println("Advent of Code 2023, Day 14")

	input := "input.txt"

	lines, err := utils.ReadLines(input)
	if err != nil {
		panic("failed to read the input")
	}

	start := time.Now()

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

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran in", elapsed)

	start = time.Now()

	// Use a new platform variable as in the second part strings are more
	// straightforward for creating cache keys.
	p := make([][]string, 0, len(lines))
	for _, line := range lines {
		row := make([]string, 0, len(line))
		for _, c := range line {
			row = append(row, string(c))
		}
		p = append(p, row)
	}

	const cycles = 1_000_000_000
	cache := make(map[string]int)
	loads := make([]int, 1)
	key := ""
	i := 0
	for i < cycles {
		tiltNorth(&p)
		tiltWest(&p)
		tiltSouth(&p)
		tiltEast(&p)

		n := 0
		for j, row := range p {
			for _, c := range row {
				if c == "O" {
					n += len(p) - j
				}
			}
		}
		loads = append(loads, n)

		for _, row := range p {
			key += strings.Join(row, "")
		}
		if _, ok := cache[key]; ok {
			break
		}
		cache[key] = i
		key = ""
		i += 1
	}

	looplen := i - cache[key]
	i = cache[key]

	t := i + (cycles-i)%looplen
	load = loads[t]

	fmt.Println("Part 2: the total load on the north support beams is", load)

	elapsed = time.Since(start)
	fmt.Println("Part 2 ran in", elapsed)
}
