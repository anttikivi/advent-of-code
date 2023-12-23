package main

import (
	"fmt"
	"image"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

// I suspect there is going to be many visited spots so the constant-time
// lookup in maps is probably better than using a slice.
func findPaths(m map[image.Point]rune, w, h int, visited map[image.Point]bool, p image.Point) int {
	next := make([]image.Point, 0)
	next = append(next, p)
	ds := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for len(next) == 1 {
		p = next[0]
		if p == image.Pt(w-2, h-1) {
			return len(visited)
		}
		visited[p] = true
		newNext := make([]image.Point, 0)
		for _, d := range ds {
			np := p.Add(d)
			if np.X < 0 || np.X >= w || np.Y < 0 || np.Y >= h {
				continue
			}
			c := m[np]
			if c == '#' {
				continue
			}
			if np.Y < p.Y && c == 'v' {
				continue
			}
			if np.Y > p.Y && c == '^' {
				continue
			}
			if np.X < p.X && c == '>' {
				continue
			}
			if np.X > p.X && c == '<' {
				continue
			}
			if _, ok := visited[np]; !ok {
				newNext = append(newNext, np)
			}
		}
		next = newNext
	}

	if len(next) == 0 {
		return 0
	}
	longest := 0
	for _, np := range next {
		newVisited := make(map[image.Point]bool)
		for k, v := range visited {
			newVisited[k] = v
		}
		result := findPaths(m, w, h, newVisited, np)
		if result > longest {
			longest = result
		}
	}
	return longest
}

func main() {
	fmt.Println("*** Advent of Code 2023 ***")
	fmt.Println("--- Day 23: A Long Walk ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	m := make(map[image.Point]rune)
	for y, line := range lines {
		for x, r := range line {
			m[image.Pt(x, y)] = r
		}
	}

	p := image.Pt(1, 0)
	hike := findPaths(m, len(lines[0]), len(lines), make(map[image.Point]bool), p)

	fmt.Println("Part 1: the longest hike is", hike, "steps")

	elapse := time.Since(start)
	fmt.Println("Part 1 ran for", elapse)
}
