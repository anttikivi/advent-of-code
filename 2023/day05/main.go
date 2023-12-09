package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func min(slice []int) int {
	m := slice[0]
	for _, num := range slice {
		if num < m {
			m = num
		}
	}
	return m
}

func covert(seeds []int, mapping []string) []int {
	var dests []int
	var srcs []int
	var ranges []int

	for _, m := range mapping {
		parts := strings.Fields(m)
		destStart, _ := strconv.Atoi(parts[0])
		srcStart, _ := strconv.Atoi(parts[1])
		r, _ := strconv.Atoi(parts[2])
		dests = append(dests, destStart)
		srcs = append(srcs, srcStart)
		ranges = append(ranges, r)
	}

	converted := make([]int, len(seeds))

	for i, seed := range seeds {
		done := false
		for j, src := range srcs {
			dest := dests[j]
			r := ranges[j]
			if seed >= src && seed < src+r {
				converted[i] = dest + (seed - src)
				done = true
			}
		}
		if !done {
			converted[i] = seed
		}
	}
	return converted
}

func main() {
	fmt.Println("Advent of Code 2023, Day 5")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	var seeds []int
	for _, num := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		seed, _ := strconv.Atoi(num)
		seeds = append(seeds, seed)
	}

	// I can't set sensibly set any initial memory to the slice as the length
	// of each of the maps should not be assumed.
	var mapping []string
	for i, line := range lines {
		if line == "" || i+1 == len(lines) {
			seeds = covert(seeds, mapping)
			mapping = nil
		} else if line[0] >= '0' && line[9] <= '9' {
			mapping = append(mapping, line)
		}
	}
	fmt.Println("Part 1: the lowest location number that corresponds to a seed is", min(seeds))

	seeds = nil

	var seedStarts []int
	var seedRanges []int
	s := strings.Fields(strings.Split(lines[0], ":")[1])
	for i := 0; i < len(s); i += 2 {
		seed, _ := strconv.Atoi(s[i])
		seedStarts = append(seedStarts, seed)
		sRange, _ := strconv.Atoi(s[i+1])
		seedRanges = append(seedRanges, sRange)
	}

	for i, start := range seedStarts {
		for j := 0; j < seedRanges[i]; j++ {
			seeds = append(seeds, start+j)
		}
	}
	for i, line := range lines {
		if line == "" || i+1 == len(lines) {
			seeds = covert(seeds, mapping)
			mapping = nil
		} else if line[0] >= '0' && line[9] <= '9' {
			mapping = append(mapping, line)
		}
	}
	fmt.Println("Part 2: the lowest location number that corresponds to a seed is", min(seeds))
}
