package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

const defaultDistance = math.MaxInt

type GraphNode struct {
	name        string
	connections map[int]bool
	index       int
}

type Edge struct {
	a, b int
}

func main() {
	fmt.Println("*** Advent of Code 2023 ***")
	fmt.Println("--- Day 25: Snowverload ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	nodes := make([]GraphNode, 0)
	nodeMap := make(map[string]int)

	addNode := func(nodeName string) int {
		if idx, ok := nodeMap[nodeName]; ok {
			return idx
		}

		nodeIndex := len(nodes)
		var n GraphNode
		n.name = nodeName
		n.connections = make(map[int]bool)
		n.index = nodeIndex
		nodes = append(nodes, n)
		nodeMap[nodeName] = nodeIndex
		return nodeIndex
	}

	for _, line := range lines {
		splits := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':' || r == ' '
		})

		name := splits[0]
		nodeIndex := addNode(name)

		for i := 1; i < len(splits); i++ {
			c := splits[i]
			cIndex := addNode(c)
			nodes[cIndex].connections[nodeIndex] = true
			nodes[nodeIndex].connections[cIndex] = true
		}
	}

	// Make a set of all of the candidate edges in this graph.
	edgeFrequencies := make(map[Edge]int)

	for _, n := range nodes {
		for c := range n.connections {
			edge := Edge{min(c, n.index), max(c, n.index)}
			edgeFrequencies[edge] = 0
		}
	}

	// Pick a bunch of random pairs of nodes and track how often we see any
	// given edge on the shortest path between them.
	distances := make([]int, len(nodes))
	for iter := 0; iter < 200; iter++ {
		for i := range distances {
			distances[i] = defaultDistance
		}

		// Pick a random pair of graph nodes
		nodeA := rand.Intn(len(nodes))
		nodeB := nodeA
		for nodeB == nodeA {
			nodeB = rand.Intn(len(nodes))
		}

		// Do a BFS to find the shortest path from A to B
		distances[nodeA] = 0
		indexQueue := []int{nodeA}
		for len(indexQueue) > 0 {
			i := indexQueue[0]
			indexQueue = indexQueue[1:]

			for c := range nodes[i].connections {
				if c == nodeB {
					indexQueue = nil
					break
				}

				if distances[c] != defaultDistance {
					continue
				}

				distances[c] = distances[i] + 1
				indexQueue = append(indexQueue, c)
			}
		}

		// Now trace back from nodeB to nodeA along the distances,
		// incrementing all the edge frequencies.
		cur := nodeB
		for cur != nodeA {
			for c := range nodes[cur].connections {
				if distances[c] < distances[cur] {
					edge := Edge{a: min(c, cur), b: max(c, cur)}
					edgeFrequencies[edge]++
					cur = c
					break
				}
			}
		}
	}

	// Flatten the edges to make indexing them easier, then sort them in
	// decreasing frequency.
	var edges []Edge
	for edge := range edgeFrequencies {
		edges = append(edges, edge)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edgeFrequencies[edges[i]] > edgeFrequencies[edges[j]]
	})

	filled := make([]bool, len(nodes))
	groupSize := 0

	// Now iterate through the edges from highest frequency to lowest frequency
	// checking every triple along the way to see if that is the triple that
	// splits the graph. We'll iterate in such a way that we push the worst
	// (lowest-frequency) node at the lowest rate, so we will exhaust all
	// higher-frequency choices before pushing our worst-case node farther up.
	for i := 2; i < len(edges); i++ {
		// This is the highest-frequency node, always starts at 0, iterates
		// next-slowest
		for j := 0; j < i-1; j++ {
			// Then this is the middle node, exhausting all middle choices
			// before moving the highest-frequency one up.
			for k := j + 1; k < i; k++ {
				for i := range filled {
					filled[i] = false
				}
				q := make([]int, 0)
				q = append(q, nodes[0].index)
				fillCount := 0

				// Spider out from node 0 to see how many we fill.
				for len(q) > 0 {
					n := q[0]
					q = q[1:]

					if filled[nodes[n].index] {
						continue
					}

					filled[nodes[n].index] = true
					fillCount++

					for c := range nodes[n].connections {
						// Don't follow edges if they're one of our three.
						e := Edge{a: min(n, c), b: max(n, c)}
						// edge := Edge{a: n, b: c}
						if edges[i] == e || edges[j] == e || edges[k] == e {
							continue
						}

						q = append(q, c)
					}
				}

				if fillCount < len(nodes) {
					// We got two groups!
					groupSize = fillCount
					goto endLoop // if Go had breaking out of outer loops I'd use that instead.
				}
			}
		}
	}

endLoop:

	result := groupSize * (len(nodes) - groupSize)

	fmt.Println("The product of the sizes of the groups is", result)

	elapsed := time.Since(start)
	fmt.Println("The program ran for", elapsed)
}
