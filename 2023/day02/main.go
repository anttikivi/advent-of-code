package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func isValid(n int, color string) bool {
	switch color {
	case "red":
		{
			return n <= MAX_RED
		}
	case "green":
		{
			return n <= MAX_GREEN
		}
	case "blue":
		{
			return n <= MAX_BLUE
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Advent of Code 2023, Day 2")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range lines {
		var (
			valid   = true
			id      = ""
			isIdSet = false
			current = ""
			color   = ""
		)

		for i, r := range line {
			if !isIdSet && i > 4 {
				if unicode.IsDigit(r) {
					id += string(r)
				} else {
					isIdSet = true
				}
			}

			if isIdSet {
				if unicode.IsDigit(r) {
					current += string(r)
				} else if unicode.IsLetter(r) {
					color += string(r)
				} else if r == ',' || r == ';' {
					n, _ := strconv.Atoi(current)
					valid = isValid(n, color)
					current = ""
					color = ""
				}
			}

			if !valid {
				break
			}
		}

		n, _ := strconv.Atoi(current)
		valid = isValid(n, color)

		if valid {
			intid, _ := strconv.Atoi(id)
			sum += intid
		}
	}

	fmt.Println("Part 1: the sum of the IDs is", sum)

	sum = 0

	for _, line := range lines {
		var (
			colonFound = false
			current    = ""
			color      = ""
			minred     = 0
			mingreen   = 0
			minblue    = 0
		)

		for _, r := range line {
			if r == ':' {
				colonFound = true
			}

			if colonFound {
				if unicode.IsDigit(r) {
					current += string(r)
				} else if unicode.IsLetter(r) {
					color += string(r)
				} else if r == ',' || r == ';' {
					n, _ := strconv.Atoi(current)
					switch color {
					case "red":
						{
							minred = max(minred, n)
						}
					case "green":
						{
							mingreen = max(mingreen, n)
						}
					case "blue":
						{
							minblue = max(minblue, n)
						}
					}
					current = ""
					color = ""
				}
			}
		}

		n, _ := strconv.Atoi(current)
		switch color {
		case "red":
			{
				minred = max(minred, n)
			}
		case "green":
			{
				mingreen = max(mingreen, n)
			}
		case "blue":
			{
				minblue = max(minblue, n)
			}
		}

		sum += minred * mingreen * minblue
	}

	fmt.Println("Part 2: the sum of the powers is", sum)
}
