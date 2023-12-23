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

	stepsRemaining := 64

	plots := 0
	ds := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited := make([]image.Point, 0)
	q := list.New()
	q.PushBack(first)
	for steps := 0; steps <= stepsRemaining; steps++ {
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
	fmt.Println("Part 1: the elf could reach", plots, "garden plots in exactly", stepsRemaining, "steps")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	stepsRemaining = 26501365

	plots = 0
	q = list.New()
	q.PushBack(first)
	polynomial := make([]int, 0)
	for steps := 0; steps < stepsRemaining; {
		nq := list.New()
		// The slice used earlier is too slow for this.
		visited := make(map[image.Point]bool)

		for q.Len() > 0 {
			el := q.Front()
			v := el.Value.(image.Point)
			q.Remove(el)
			for _, d := range ds {
				np := v.Add(d)
				mp := image.Pt(((np.X%len(lines[0]))+len(lines[0]))%len(lines[0]), ((np.Y%len(lines))+len(lines))%len(lines))
				if _, ok := visited[np]; !ok && m[mp] != '#' {
					visited[np] = true
					nq.PushBack(np)
				}

			}

		}
		steps += 1
		q = nq
		if steps%(len(lines)) == stepsRemaining%len(lines) {
			polynomial = append(polynomial, len(visited))

			if len(polynomial) == 3 {
				p0 := polynomial[0]
				p1 := polynomial[1] - polynomial[0]
				p2 := polynomial[2] - polynomial[1]

				plots = p0 + (p1 * (stepsRemaining / len(lines))) + ((stepsRemaining/len(lines))*((stepsRemaining/len(lines))-1)/2)*(p2-p1)
				break
			}
		}

	}

	fmt.Println("Part 2: the elf could reach", plots, "garden plots in exactly", stepsRemaining, "steps")

	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
