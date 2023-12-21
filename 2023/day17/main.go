package main

import (
	"container/heap"
	"fmt"
	"image"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

const (
	VerticalPlane = iota
	HorizontalPlane
	UndecidedPlane
)

type vertex struct {
	pos                image.Point
	dir                int
	visited            bool
	heatloss           int
	calculatedHeatloss int
	total              int
	index              int
}

type graph struct {
	vertices []vertex
	width    int
	height   int
}

type pqueue []*vertex

func (pq pqueue) Len() int {
	return len(pq)
}

func (pq pqueue) Less(i, j int) bool {
	return pq[i].total < pq[j].total
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pqueue) Push(x any) {
	n := len(*pq)
	item := x.(*vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pqueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *pqueue) update(item *vertex, priority int) {
	heap.Fix(pq, item.index)
}

func (g *graph) getEdges(u *vertex, min int, max int) []*vertex {
	e := make([]*vertex, 0, 6)

	if u.dir == HorizontalPlane || u.dir == UndecidedPlane {
		for heatloss, dy := 0, 1; dy <= max; dy++ {
			v := g.getVertexByCoords(u.pos.X, u.pos.Y+dy, VerticalPlane)
			if v != nil {
				heatloss += v.heatloss
				if dy >= min {
					v.calculatedHeatloss = heatloss
					e = append(e, v)
				}
			}
		}
		for heatloss, dy := 0, 1; dy <= max; dy++ {
			v := g.getVertexByCoords(u.pos.X, u.pos.Y-dy, VerticalPlane)
			if v != nil {
				heatloss += v.heatloss
				if dy >= min {
					v.calculatedHeatloss = heatloss
					e = append(e, v)
				}
			}
		}
	}

	if u.dir == VerticalPlane || u.dir == UndecidedPlane {
		for heatloss, dx := 0, 1; dx <= max; dx++ {
			v := g.getVertexByCoords(u.pos.X+dx, u.pos.Y, HorizontalPlane)
			if v != nil {
				heatloss += v.heatloss
				if dx >= min {
					v.calculatedHeatloss = heatloss
					e = append(e, v)
				}
			}
		}
		for heatloss, dx := 0, 1; dx <= max; dx++ {
			v := g.getVertexByCoords(u.pos.X-dx, u.pos.Y, HorizontalPlane)
			if v != nil {
				heatloss += v.heatloss
				if dx >= min {
					v.calculatedHeatloss = heatloss
					e = append(e, v)
				}
			}
		}
	}

	return e
}

func (g *graph) getVertexByCoords(x int, y int, plane int) *vertex {
	if x < 0 || y < 0 || y >= g.height || x >= g.width {
		return nil
	}
	return &g.vertices[y*2*g.width+x*2+plane]
}

func createGraph(grid [][]int) graph {
	g := graph{}
	vertices := make([]vertex, 0, len(grid)*len(grid)*2)
	g.height = len(grid)
	for y := range grid {
		g.width = len(grid[y])
		for x := range grid[y] {
			vertices = append(vertices, vertex{
				pos:      image.Pt(x, y),
				dir:      VerticalPlane,
				total:    1 << 30,
				heatloss: grid[y][x],
			})
			vertices = append(vertices, vertex{
				pos:      image.Pt(x, y),
				dir:      HorizontalPlane,
				total:    1 << 30,
				heatloss: grid[y][x],
			})
		}
	}
	g.vertices = vertices
	return g
}

func findPath(grid [][]int, min int, max int) int {
	graph := createGraph(grid)
	vertices := graph.vertices

	vertices[0].total = 0
	vertices[0].dir = UndecidedPlane

	pq := make(pqueue, len(vertices))
	for i := 0; i < len(vertices); i++ {
		vertices[i].index = i
		pq[i] = &vertices[i]
	}
	heap.Init(&pq)

	var u *vertex
	var e = &vertices[len(vertices)-1]
	for {
		u = heap.Pop(&pq).(*vertex)

		if u.pos.X == e.pos.X && u.pos.Y == e.pos.Y {
			break
		}

		u.visited = true

		for _, e := range graph.getEdges(u, min, max) {
			if u.total+e.calculatedHeatloss < e.total {
				e.total = u.total + e.calculatedHeatloss
				pq.update(e, e.total)
			}
		}
	}
	return u.total
}

func main() {
	fmt.Println("Advent of Code 2023, Day 17")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("failed to read the input file")
	}

	start := time.Now()

	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, r := range line {
			grid[y][x] = int(r - '0')
		}
	}

	sum := findPath(grid, 1, 3)
	fmt.Println("Part 1: the least heat loss that can incur is", sum)
	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	sum = findPath(grid, 4, 10)
	fmt.Println("Part 2: the least heat loss that can incur is", sum)
	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
