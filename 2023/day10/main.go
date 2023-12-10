package main

import (
	"fmt"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func findFirst(sketch []string, x, y int) (int, int) {
	var nextX, nextY int
	dy := []int{0, 1, 0, -1}
	dx := []int{-1, 0, 1, 0}
	for i := 0; i < 4; i++ {
		if y+dy[i] >= 0 && y+dy[i] < len(sketch) && x+dx[i] >= 0 && x+dx[i] < len(sketch[0]) {
			nx, ny := dx[i], dy[i]
			nextX, nextY = x+nx, y+ny
			c := rune(sketch[nextY][nextX])
			if (nx == -1 && ny == 0) && (c == '-' || c == 'L' || c == 'F') {
				return nextX, nextY
			}
			if (nx == 0 && ny == 1) && (c == '|' || c == '7' || c == 'F') {
				return nextX, nextY
			}
			if (nx == 1 && ny == 0) && (c == '-' || c == 'J' || c == '7') {
				return nextX, nextY
			}
			if (nx == 0 && ny == -1) && (c == '|' || c == 'L' || c == 'J') {
				return nextX, nextY
			}
		}
	}
	return nextX, nextY
}

func findNext(sketch []string, x, y, lastX, lastY int) (int, int) {
	current := rune(sketch[y][x])
	if current == '|' {
		if lastY < y {
			return x, y + 1
		}
		return x, y - 1
	}
	if current == '-' {
		if lastX < x {
			return x + 1, y
		}
		return x - 1, y
	}
	if current == 'L' {
		if lastY < y {
			return x + 1, y
		}
		return x, y - 1
	}
	if current == 'J' {
		if lastY < y {
			return x - 1, y
		}
		return x, y - 1
	}
	if current == '7' {
		if lastY > y {
			return x - 1, y
		}
		return x, y + 1
	}
	if current == 'F' {
		if lastY > y {
			return x + 1, y
		}
		return x, y + 1
	}
	return x, y
}

func main() {
	fmt.Println("Advent of Code 2023, Day 10")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	var sketch []string
	x, y := 0, 0

	for i, line := range lines {
		sketch = append(sketch, line)
		if strings.Contains(line, "S") {
			y = i
			x = strings.Index(line, "S")
		}
	}

	lastX, lastY := x, y
	x, y = findFirst(sketch, x, y)

	current := rune(sketch[y][x])
	steps := 1

	for current != 'S' {
		nx, ny := findNext(sketch, x, y, lastX, lastY)
		lastX, lastY = x, y
		x, y = nx, ny
		current = rune(sketch[y][x])
		steps += 1
	}

	fmt.Println("Part 1: the farthest point is", int(steps>>1), "steps away")
}
