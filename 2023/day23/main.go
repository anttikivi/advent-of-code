package main

import (
	"fmt"
	"image"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/emirpasic/gods/queues/arrayqueue"
)

type pointAndLength struct {
	P   image.Point
	Len int
}

type Graph [142][142]int

var graph Graph
var maxID int

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

func isBitSet(bitSet, bitPos int) bool {
	return (bitSet & (1 << bitPos)) != 0
}

func setBit(bitSet, bitPos int) int {
	return bitSet | (1 << bitPos)
}

func searchDepthFirst(current, prev, end int, visited int) int {
	if current == end {
		return 0
	}
	maxi := -1
	for j := 1; j < maxID; j++ {
		if graph[current][j] != 0 && j != prev && !isBitSet(visited, j) {
			d := searchDepthFirst(j, current, end, setBit(visited, j))
			if d != -1 {
				maxi = max(maxi, d+graph[current][j])
			}
		}
	}
	return maxi
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

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	w := len(lines[0])
	h := len(lines)
	sp := image.Pt(1, 0)
	ep := image.Pt(w-2, h-1)
	adjacents := make(map[image.Point][]image.Point)
	nodes := make(map[image.Point]int)
	nodes[sp] = 1
	nodes[ep] = 2
	id := 3
	ds := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := image.Pt(x, y)
			if p.X < 0 || p.X >= w || p.Y < 0 || p.Y >= h || m[p] == '#' {
				continue
			}

			adjacents[p] = make([]image.Point, 0)
			for _, d := range ds {
				np := p.Add(d)
				if np.X < 0 || np.X >= w || np.Y < 0 || np.Y >= h || m[np] == '#' {
					continue
				}
				adjacents[p] = append(adjacents[p], np)
			}
			if len(adjacents[p]) > 2 {
				nodes[p] = id
				id += 1
			}
		}
	}
	maxID = id

	graph = Graph{}
	for p := range nodes {
		id, queue, visited := nodes[p], arrayqueue.New(), mapset.NewSet[image.Point]()
		queue.Enqueue(pointAndLength{p, 1})
		visited.Add(p)

		for !queue.Empty() {
			pwl, _ := queue.Dequeue()
			p, distance := pwl.(pointAndLength).P, pwl.(pointAndLength).Len
			for _, neighbor := range adjacents[p] {
				if !visited.Contains(neighbor) {
					if nid, has := nodes[neighbor]; has {
						graph[id][nid] = distance
						graph[nid][id] = distance
					} else {
						queue.Enqueue(pointAndLength{neighbor, distance + 1})
					}
					visited.Add(neighbor)
				}
			}
		}
	}

	hike = searchDepthFirst(1, 1, 2, 0)

	fmt.Println("Part 2: the longest hike is", hike, "steps")
	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
