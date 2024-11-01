package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func isZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}

	return true
}

func extrapolate(nums []int) int {
	var rows [][]int
	i := 0
	rows = append(rows, nums)
	for !isZero(rows[i]) {
		row := rows[i]
		var next []int
		end := len(row) - 1
		for j := 0; j < end; j++ {
			if j < end {
				next = append(next, row[j+1]-row[j])
			}
		}
		rows = append(rows, next)
		i += 1
	}
	rows[i] = append(rows[i], 0)
	for i > 0 {
		len := len(rows[i])
		rows[i-1] = append(rows[i-1], rows[i-1][len-1]+rows[i][len-1])
		i -= 1
	}
	return rows[0][len(rows[0])-1]
}

func main() {
	fmt.Println("Advent of Code 2023, Day 9")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	c, c2 := 0, 0

	for _, line := range lines {
		fields := strings.Fields(line)
		nums := make([]int, len(fields))
		for i, field := range fields {
			nums[i], err = strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
		}
		c += extrapolate(nums)
		slices.Reverse(nums)
		c2 += extrapolate(nums)
	}

	fmt.Println("Part 1: the sum of the extrapolated values is", c)
	fmt.Println("Part 2: the sum of the extrapolated values is", c2)
}
