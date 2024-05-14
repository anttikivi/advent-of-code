package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

const (
	lowerLimit = 200000000000000
	upperLimit = 400000000000000
)

type hailstone struct {
	Px, Py, Pz, Vx, Vy, Vz int
}

func main() {
	fmt.Println("*****    Advent of Code 2023     *****")
	fmt.Println("--- Day 24: Never Tell Me The Odds ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("Failed to read the input file")
	}

	start := time.Now()

	hailstones := make([]hailstone, 0)
	for _, line := range lines {
		parts := strings.Split(line, " @ ")
		p := strings.Split(parts[0], ", ")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		pz, _ := strconv.Atoi(p[2])
		v := strings.Split(parts[1], ", ")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		vz, _ := strconv.Atoi(v[2])
		stone := hailstone{px, py, pz, vx, vy, vz}
		hailstones = append(hailstones, stone)
	}

	count := 0
	for i, stoneA := range hailstones {
		for _, stoneB := range hailstones[i+1:] {
			dx := stoneA.Px - stoneB.Px
			dy := stoneA.Py - stoneB.Py
			denominator := stoneA.Vy*stoneB.Vx - stoneA.Vx*stoneB.Vy
			if denominator == 0 {
				continue
			}
			numerator := stoneB.Vy*dx - stoneB.Vx*dy
			tA := numerator / denominator
			if tA <= 0 {
				continue
			}
			numerator = stoneA.Vy*dx - stoneA.Vx*dy
			tB := numerator / denominator
			if tB <= 0 {
				continue
			}
			x := stoneA.Px + tA*stoneA.Vx
			y := stoneA.Py + tA*stoneA.Vy
			if x < lowerLimit || x > upperLimit || y < lowerLimit || y > upperLimit {
				continue
			}
			count++
		}
	}

	fmt.Println("Part 1: a total of", count, "intersections happen within the test area")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)
}
