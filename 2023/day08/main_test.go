package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func part2(instructions string, network map[string]*node, starts []string) int {
	counts := make([]int, 0, len(starts))
	for _, start := range starts {
		j := 0
		count := 0
		current := start
		imax := len(instructions) - 1

		for !strings.HasSuffix(current, "Z") {
			instruction := instructions[j]
			node := network[current]
			if instruction == 'L' {
				current = node.left
			} else {
				current = node.right
			}
			count += 1
			if j == imax {
				j = 0
			} else {
				j += 1
			}
		}
		counts = append(counts, count)
	}

	c := 1
	for _, count := range counts {
		c = lcm(c, count)
	}

	return c
}

func part2Goroutines(instructions string, network map[string]*node, starts []string) int {
	counts := make([]int, 0, len(starts))
	c := make(chan int, len(starts))
	for _, start := range starts {
		go func(start string) {
			j := 0
			count := 0
			current := start
			imax := len(instructions) - 1

			for !strings.HasSuffix(current, "Z") {
				instruction := instructions[j]
				node := network[current]
				if instruction == 'L' {
					current = node.left
				} else {
					current = node.right
				}
				count += 1
				if j == imax {
					j = 0
				} else {
					j += 1
				}
			}
			c <- count
		}(start)
		counts = append(counts, <-c)
	}

	ct := 1
	for _, count := range counts {
		ct = lcm(ct, count)
	}

	return ct
}

func TestPart2(t *testing.T) {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	instructions := lines[0]
	network := make(map[string]*node)
	var starts []string

	for _, line := range lines[2:] {
		var current string
		var left string
		var right string
		_, err := fmt.Sscanf(line, "%s = %s %s", &current, &left, &right)
		if err != nil {
			panic(err)
		}
		network[current] = &node{left: left[1:4], right: right[:3]}
		if strings.HasSuffix(current, "A") {
			starts = append(starts, current)
		}
	}

	want := 21366921060721
	got1 := part2(instructions, network, starts)
	if got1 != want {
		t.Errorf("part2() = %d, want %d", got1, want)
	}
	got2 := part2Goroutines(instructions, network, starts)
	if got2 != want {
		t.Errorf("part2Goroutines() = %d, want %d", got2, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	instructions := lines[0]
	network := make(map[string]*node)
	var starts []string
	for _, line := range lines[2:] {
		var current string
		var left string
		var right string
		_, err := fmt.Sscanf(line, "%s = %s %s", &current, &left, &right)
		if err != nil {
			panic(err)
		}
		network[current] = &node{left: left[1:4], right: right[:3]}
		if strings.HasSuffix(current, "A") {
			starts = append(starts, current)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(instructions, network, starts)
	}
}

func BenchmarkPart2Goroutines(b *testing.B) {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	instructions := lines[0]
	network := make(map[string]*node)
	var starts []string
	for _, line := range lines[2:] {
		var current string
		var left string
		var right string
		_, err := fmt.Sscanf(line, "%s = %s %s", &current, &left, &right)
		if err != nil {
			panic(err)
		}
		network[current] = &node{left: left[1:4], right: right[:3]}
		if strings.HasSuffix(current, "A") {
			starts = append(starts, current)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2Goroutines(instructions, network, starts)
	}
}
