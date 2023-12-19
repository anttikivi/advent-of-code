package main

import (
	"fmt"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func transpose(pattern []string) []string {
	result := make([]string, 0, len(pattern[0]))
	for i := 0; i < len(pattern[0]); i++ {
		var row string
		for j := 0; j < len(pattern); j++ {
			row += string(pattern[j][i])
		}
		result = append(result, row)
		row = ""
	}
	return result
}

func reflect(pattern []string) int {
	n := 0
	found := false
	for i := 1; i < len(pattern); i++ {
		reflectRange := min(i, len(pattern)-i)
		for j := 1; j <= reflectRange; j++ {
			if pattern[i-j] != pattern[i+j-1] {
				break
			}
			if j == reflectRange {
				found = true
				n += i
				break
			}
		}
		if found {
			break
		}
	}
	return n
}

func main() {
	fmt.Println("Advent of Code 2023, Day 13")

	input := "test.txt"

	lines, err := utils.ReadLines(input)
	if err != nil {
		panic("error reading input")
	}

	var patterns [][]string
	var current []string

	for _, line := range lines {
		fmt.Println(line)
		if line == "" {
			patterns = append(patterns, current)
			current = make([]string, 0)
		} else {
			current = append(current, line)
		}
	}
	patterns = append(patterns, current)

	sum := 0
	for _, pattern := range patterns {
		fmt.Println("Checking pattern")
		for _, line := range pattern {
			fmt.Println(line)
		}
		n := reflect(pattern) * 100
		if n == 0 {
			n = reflect(transpose(pattern))
		}
		sum += n
	}

	fmt.Println("Part 1: the number from summarising the notes is", sum)
}
