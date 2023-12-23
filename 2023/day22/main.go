package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type point struct {
	X, Y, Z int
}

type brick struct {
	Start, End point
}

func resolveZs(xy [][]int, p1, p2 point) (int, int) {
	z1, z2 := 0, 0

	if p1.X == p2.X && p1.Z == p2.Z {
		for i := p1.Y; i <= p2.Y; i++ {
			if xy[p1.X][i] > z1 {
				z1 = xy[p1.X][i]
			}
		}

		z1, z2 = z1+1, z1+1

		for i := p1.Y; i <= p2.Y; i++ {
			xy[p1.X][i] = z1
		}
	} else if p1.Y == p2.Y && p1.Z == p2.Z {
		for i := p1.X; i <= p2.X; i++ {
			if xy[i][p1.Y] > z1 {
				z1 = xy[i][p1.Y]
			}
		}

		z1, z2 = z1+1, z1+1

		for i := p1.X; i <= p2.X; i++ {
			xy[i][p1.Y] = z1
		}
	} else if p1.X == p2.X && p1.Y == p2.Y {
		z1 = xy[p1.X][p1.Y] + 1
		z2 = z1 + p2.Z - p1.Z
		xy[p1.X][p1.Y] = z2
	}
	return z1, z2
}

func shiftBricks(snapshot []brick) int {
	shifted := 0
	xy := make([][]int, 10)
	for i := range xy {
		xy[i] = make([]int, 10)
	}

	for i, b := range snapshot {
		b.Start.Z, b.End.Z = resolveZs(xy, b.Start, b.End)

		if b.Start != snapshot[i].Start || b.End != snapshot[i].End {
			snapshot[i] = brick{b.Start, b.End}
			shifted += 1
		}
	}
	return shifted
}

func main() {
	fmt.Println("*** Advent of Code 2023 ***")
	fmt.Println("--- Day 22: Sand Slabs  ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("Failed to read the input file")
	}

	start := time.Now()

	snapshot := make([]brick, 0)
	for _, line := range lines {
		ends := strings.Split(line, "~")
		start := strings.Split(ends[0], ",")
		end := strings.Split(ends[1], ",")
		sx, _ := strconv.Atoi(start[0])
		sy, _ := strconv.Atoi(start[1])
		sz, _ := strconv.Atoi(start[2])
		ex, _ := strconv.Atoi(end[0])
		ey, _ := strconv.Atoi(end[1])
		ez, _ := strconv.Atoi(end[2])
		snapshot = append(snapshot, brick{point{sx, sy, sz}, point{ex, ey, ez}})
	}
	slices.SortFunc(snapshot, func(a, b brick) int { return a.Start.Z - b.Start.Z })

	shiftBricks(snapshot)

	canDisintegrate := 0

	for i := range snapshot {
		newSnapshot := make([]brick, len(snapshot))
		copy(newSnapshot, snapshot)
		if i == 0 {
			newSnapshot = newSnapshot[1:]
		} else if i == len(newSnapshot)-1 {
			newSnapshot = snapshot[:len(newSnapshot)-1]
		} else {
			newSnapshot = append(newSnapshot[:i], newSnapshot[i+1:]...)
		}

		shifted := shiftBricks(newSnapshot)
		if shifted == 0 {
			canDisintegrate++
		}
	}

	fmt.Println("Part 1: a total of", canDisintegrate, "bricks can be safely disintegrated")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	snapshot = make([]brick, 0)
	for _, line := range lines {
		ends := strings.Split(line, "~")
		start := strings.Split(ends[0], ",")
		end := strings.Split(ends[1], ",")
		sx, _ := strconv.Atoi(start[0])
		sy, _ := strconv.Atoi(start[1])
		sz, _ := strconv.Atoi(start[2])
		ex, _ := strconv.Atoi(end[0])
		ey, _ := strconv.Atoi(end[1])
		ez, _ := strconv.Atoi(end[2])
		snapshot = append(snapshot, brick{point{sx, sy, sz}, point{ex, ey, ez}})
	}
	slices.SortFunc(snapshot, func(a, b brick) int { return a.Start.Z - b.Start.Z })

	shiftBricks(snapshot)

	sum := 0

	for i := range snapshot {
		newSnapshot := make([]brick, len(snapshot))
		copy(newSnapshot, snapshot)
		if i == 0 {
			newSnapshot = newSnapshot[1:]
		} else if i == len(newSnapshot)-1 {
			newSnapshot = snapshot[:len(newSnapshot)-1]
		} else {
			newSnapshot = append(newSnapshot[:i], newSnapshot[i+1:]...)
		}

		sum += shiftBricks(newSnapshot)
	}

	fmt.Println("Part 2: the sum of the bricks that would fall is", sum)
	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
