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

func calculateDeterminant(m [][]int) int {
	if len(m) == 0 {
		return 1
	}

	var (
		l    = m[0]
		r    = m[1:]
		prod = make([]int, 0)
	)

	for i, n := range l {
		spliced := make([][]int, 0)
		for _, row := range r {
			spliced = append(spliced, append([]int{}, row...))
		}
		for j := range spliced {
			spliced[j] = append(spliced[j][:i], spliced[j][i+1:]...)
		}
		prod = append(prod, n*calculateDeterminant(spliced))
	}

	var det int

	for i, n := range prod {
		if i%2 == 0 {
			det += n
		} else {
			det -= n
		}
	}

	return det
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

	start = time.Now()

	a, b := make([][]int, 0), make([]int, 0)

	var (
		px0 = hailstones[0].Px
		py0 = hailstones[0].Py
		pz0 = hailstones[0].Pz
		vx0 = hailstones[0].Vx
		vy0 = hailstones[0].Vy
		vz0 = hailstones[0].Vz
	)

	for i := 1; i <= 3; i++ {
		var (
			pxN = hailstones[i].Px
			pyN = hailstones[i].Py
			pzN = hailstones[i].Pz
			vxN = hailstones[i].Vx
			vyN = hailstones[i].Vy
			vzN = hailstones[i].Vz
		)
		a = append(a, []int{vy0 - vyN, vxN - vx0, 0, pyN - py0, px0 - pxN, 0})
		b = append(b, px0*vy0-py0*vx0-pxN*vyN+pyN*vxN)
		a = append(a, []int{vz0 - vzN, 0, vxN - vx0, pzN - pz0, 0, px0 - pxN})
		b = append(b, px0*vz0-pz0*vx0-pxN*vzN+pzN*vxN)
	}

	det := calculateDeterminant(a)
	blen := len(b)
	var c []int
	for i := range a {
		bi := make([][]int, blen)
		for j, row := range a {
			bi[j] = append([]int{}, row...)
		}
		for j := range b {
			bi[j][i] = b[j]
		}
		c = append(c, calculateDeterminant(bi)/det)
	}

	sum := c[0] + c[1] + c[2]

	fmt.Println("Part 2: the sum of the initial coordinates is", sum)

	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
