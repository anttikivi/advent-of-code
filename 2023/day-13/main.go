package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

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

func mirror(s []string, equal func([]string, []string) bool) int {
	for i := 1; i < len(s); i++ {
		l := slices.Min([]int{i, len(s) - i})
		a, b := slices.Clone(s[i-l:i]), s[i:i+l]
		slices.Reverse(a)
		if equal(a, b) {
			return i
		}
	}
	return 0
}

func smudge(a, b []string) bool {
	diffs := 0
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				diffs++
			}
		}
	}
	return diffs == 1
}

func part2(file string) int {
	input, err := os.ReadFile(file)
	if err != nil {
		panic("failed to read the file")
	}

	sum := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		rows, cols := []string{}, make([]string, len(strings.Fields(s)[0]))
		for _, s := range strings.Fields(s) {
			rows = append(rows, s)
			for i, r := range s {
				cols[i] += string(r)
			}
		}

		sum += mirror(cols, smudge) + 100*mirror(rows, smudge)
	}
	return sum
}

func main() {
	fmt.Println("Advent of Code 2023, Day 13")

	input := "input.txt"

	lines, err := utils.ReadLines(input)
	if err != nil {
		panic("error reading input")
	}

	var patterns [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, current)
			current = make([]string, 0)
		} else {
			current = append(current, line)
		}
	}
	patterns = append(patterns, current)

	sum := 0
	for _, p := range patterns {
		n := reflect(p) * 100
		if n == 0 {
			n = reflect(transpose(p))
		}

		sum += n
	}

	fmt.Println("Part 1: the number from summarising the notes is", sum)
	fmt.Println("Part 2: the number from summarising the notes is", part2(input))
}
