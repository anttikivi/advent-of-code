package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// The input is so small it can reasonably be embedded.
//
//go:embed input.txt
var input string

func main() {
	fmt.Println("Advent of Code 2023, Day 6")

	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1.0

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])
		// Find the range of possibles time to push the button. The function is
		// a parabola, so we can find the roots and use them as the range. The
		// function for the distance the boat travels is:
		// f(x) = -x^2 + tx
		// where x is the time that the button is pushed and t is the time the
		// race lasts.
		det := math.Sqrt(float64(t*t - 4*d))
		tf := float64(t)
		l := math.Floor((tf - det) / 2)
		h := math.Ceil(((tf + det) / 2) - 1)
		p *= h - l
	}

	fmt.Println("Part 1: the product of the possible ways to win is", p)

	time := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	dist := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")
	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(dist)
	// Do the same calculation as in part 1.
	det := math.Sqrt(float64(t*t - 4*d))
	tf := float64(t)
	l := math.Floor((tf - det) / 2)
	h := math.Ceil(((tf + det) / 2) - 1)

	fmt.Println("Part 2: the number of different ways the race can be won is", int(h-l))
}
