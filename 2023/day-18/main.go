package main

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func main() {
	fmt.Println("****  Advent of Code 2023  ****")
	fmt.Println("--- Day 18: Lavaduct Lagoon ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("Failed to read the input file")
	}

	start := time.Now()

	edge := make([]image.Point, 0, len(lines))
	current := image.Pt(0, 0)
	edge = append(edge, current)
	edgelen := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		d := parts[0]
		n, _ := strconv.Atoi(parts[1])
		switch d {
		case "U":
			current.Y -= n
		case "R":
			current.X += n
		case "D":
			current.Y += n
		case "L":
			current.X -= n
		}
		edge = append(edge, current)
		edgelen += n
	}

	// Use the shoelace formula for the inner area.
	sum, l := 0, len(edge)
	for i := 0; i < l; i++ {
		p1 := edge[i]
		var p2 image.Point
		if i == l-1 {
			p2 = edge[0]
		} else {
			p2 = edge[i+1]
		}
		sum += (p1.X * p2.Y) - (p1.Y * p2.X)
	}
	sum = int(math.Abs(float64(sum >> 1)))
	// Add the length of the edge of the lagoon as the shoelace formula doesn't
	// include it in the area.
	sum += edgelen>>1 + 1

	fmt.Println("Part 1: the lagoon can hold a total of", sum, "cubic metres of lava")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	edge = make([]image.Point, 0, len(lines))
	current = image.Pt(0, 0)
	edge = append(edge, current)
	edgelen = 0
	for _, line := range lines {
		parts := strings.Fields(line)
		hex := strings.TrimPrefix(strings.TrimSuffix(parts[2], ")"), "(#")
		d, _ := strconv.Atoi(hex[len(hex)-1:])
		n64, _ := strconv.ParseInt(hex[:len(hex)-1], 16, 64)
		n := int(n64)
		switch d {
		case 0:
			current.X += n
		case 1:
			current.Y += n
		case 2:
			current.X -= n
		case 3:
			current.Y -= n
		}
		edge = append(edge, current)
		edgelen += n
	}

	sum, l = 0, len(edge)
	for i := 0; i < l; i++ {
		p1 := edge[i]
		var p2 image.Point
		if i == l-1 {
			p2 = edge[0]
		} else {
			p2 = edge[i+1]
		}
		sum += (p1.X * p2.Y) - (p1.Y * p2.X)
	}
	sum = int(math.Abs(float64(sum >> 1)))
	sum += edgelen>>1 + 1

	fmt.Println("Part 2: the lagoon can hold a total of", sum, "cubic metres of lava")

	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
