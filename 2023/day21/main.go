package main

import (
	"container/list"
	"fmt"
	"image"
	"slices"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

const InputFile = "input.txt"
const StepsRemaining = 64

func main() {
	fmt.Println("*** Advent of Code 2023  ***")
	fmt.Println("--- Day 21: Step Counter ---")

	lines, err := utils.ReadLines(InputFile)
	if err != nil {
		panic("Failed to read the input file")
	}

	start := time.Now()

	m := make(map[image.Point]rune)
	var first image.Point
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				pt := image.Pt(x, y)
				m[pt] = '.'
				first = pt
			} else {
				m[image.Pt(x, y)] = r
			}
		}
	}

	plots := 0
	ds := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited := make([]image.Point, 0)
	q := list.New()
	q.PushBack(first)
	for steps := 0; steps <= StepsRemaining; steps++ {
		nq := list.New()
		for q.Len() > 0 {
			el := q.Front()
			v := el.Value.(image.Point)
			q.Remove(el)
			// It seems that the possible points where the steps can end are
			// the destinations of the even steps.
			if steps%2 == 0 && steps > 0 {
				plots += 1
			}
			for _, d := range ds {
				np := v.Add(d)
				// Get the key for the map by modulo arithmetic to ensure that
				// the values is within bounds.
				p := image.Pt(((np.X%len(lines[0]))+len(lines[0]))%len(lines[0]), ((np.Y%len(lines))+len(lines))%len(lines))
				if m[p] != '#' && !slices.Contains(visited, p) {
					visited = append(visited, p)
					nq.PushBack(p)
				}
			}
		}
		q = nq
	}
	fmt.Println("Part 1: the elf could reach", plots, "garden plots in exactly", StepsRemaining, "steps")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)
}
