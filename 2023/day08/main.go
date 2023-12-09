package main

import (
	"fmt"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type node struct {
	left  string
	right string
}

func gcd(a, b int) int {
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	fmt.Println("Advent of Code 2023, Day 8")

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

	i := 0
	c := 0
	current := "AAA"
	imax := len(instructions) - 1

	for current != "ZZZ" {
		instruction := instructions[i]
		node := network[current]
		if instruction == 'L' {
			current = node.left
		} else {
			current = node.right
		}
		c += 1
		if i == imax {
			i = 0
		} else {
			i += 1
		}
	}

	fmt.Println("Part 1: total of", c, "steps were taken to get to ZZZ")

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

	c = 1
	for _, count := range counts {
		c = lcm(c, count)
	}

	fmt.Println("Part 2: total of", c, "steps were taken")
}
