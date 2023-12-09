package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func calculateCardPoints(card string) int {
	points := 0
	f := func(c rune) bool {
		return c == ':' || c == '|'
	}
	parts := strings.FieldsFunc(card, f)
	var winning []int
	var nums []int

	for _, num := range strings.Fields(parts[1]) {
		n, _ := strconv.Atoi(num)
		winning = append(winning, n)
	}
	for _, num := range strings.Fields(parts[2]) {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}
	for _, num := range nums {
		if slices.Contains(winning, num) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
}

func main() {
	fmt.Println("Advent of Code 2023, Day 4")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range lines {
		sum += calculateCardPoints(line)
	}

	fmt.Println("Part 1: the cards are worth", sum, "points in total")

	sum = 0
	copies := make([]int, len(lines))

	// There is always at least one copy of each card.
	for i := range copies {
		copies[i] = 1
	}

	for i, line := range lines {
		f := func(c rune) bool {
			return c == ':' || c == '|'
		}
		parts := strings.FieldsFunc(line, f)
		var winning []int
		var nums []int

		for _, num := range strings.Fields(parts[1]) {
			n, _ := strconv.Atoi(num)
			winning = append(winning, n)
		}
		for _, num := range strings.Fields(parts[2]) {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		matching := 0
		for _, num := range nums {
			if slices.Contains(winning, num) {
				matching += 1
			}
		}

		if matching > 0 {
			for j := matching; j > 0; j-- {
				if i+j < len(copies) {
					copies[i+j] += copies[i]
				}
			}
		}
	}

	for _, c := range copies {
		sum += c
	}

	fmt.Println("Part 2: the total number of cards is", sum)
}
