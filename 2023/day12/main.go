package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type spring int

const (
	unknown     spring = iota
	damaged     spring = iota
	operational spring = iota
)

type record struct {
	springs []spring
	counts  []int
}

func sum(a []int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}

func (r *record) isValid() bool {
	var current int
	var cIdx int

	for i, s := range r.springs {
		if s == damaged {
			current += 1
		}

		if s != damaged || i == len(r.springs)-1 {
			if current > 0 {
				if cIdx >= len(r.counts) || r.counts[cIdx] != current {
					return false
				}
				cIdx += 1
				current = 0
			}
		}
	}

	// Ensure all counts were checked
	return cIdx == len(r.counts)
}

func (r *record) countArrangements() int {
	for i, s := range r.springs {
		if s == unknown {
			damagedSpring := make([]spring, len(r.springs))
			copy(damagedSpring, r.springs)
			damagedSpring[i] = damaged
			d := record{damagedSpring, r.counts}

			operationalSpring := make([]spring, len(r.springs))
			copy(operationalSpring, r.springs)
			operationalSpring[i] = operational
			o := record{operationalSpring, r.counts}

			return d.countArrangements() + o.countArrangements()
		}
	}

	if r.isValid() {
		return 1
	}

	return 0
}

func parse(lines []string) []record {
	records := make([]record, 0, len(lines))

	for _, line := range lines {
		parts := strings.Fields(line)
		var rec record
		for _, c := range parts[0] {
			if c == '.' {
				rec.springs = append(rec.springs, operational)
			} else if c == '#' {
				rec.springs = append(rec.springs, damaged)
			} else if c == '?' {
				rec.springs = append(rec.springs, unknown)
			} else {
				panic("invalid input")
			}
		}
		for _, c := range strings.Split(parts[1], ",") {
			n, err := strconv.Atoi(c)
			if err != nil {
				panic("parsing count failed!")
			}
			rec.counts = append(rec.counts, n)
		}
		records = append(records, rec)
	}

	return records
}

func countInnerArrangements(springs []spring, counts []int, cache [][]*int) int {
	if len(counts) == 0 {
		if slices.Contains(springs, damaged) {
			return 0
		}
		return 1
	}

	if len(springs) < sum(counts)+len(counts) {
		return 0
	}

	if cache[len(counts)-1][len(springs)-1] != nil {
		return *cache[len(counts)-1][len(springs)-1]
	}

	arrangements := 0

	if springs[0] != damaged {
		arrangements += countInnerArrangements(springs[1:], counts, cache)
	}

	nextGroupSize := counts[0]

	if !slices.Contains(springs[:nextGroupSize], operational) && springs[nextGroupSize] != damaged {
		arrangements += countInnerArrangements(springs[nextGroupSize+1:], counts[1:], cache)
	}

	cache[len(counts)-1][len(springs)-1] = &arrangements

	return arrangements
}

func countUnfoldedArrangements(springs []spring, counts []int) int {
	// Add an operational spring at the end to simplify the recursion.
	springs = append(springs, operational)

	cache := make([][]*int, len(counts))
	for i := range cache {
		cache[i] = make([]*int, len(springs))
	}

	return countInnerArrangements(springs, counts, cache)
}

func part2(input string) int {
	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	records := parse(lines)
	sum := 0

	for _, r := range records {
		e := record{make([]spring, 0, len(r.springs)*5+4), make([]int, 0, len(r.counts)*5)}
		for i := 0; i < 5; i++ {
			if i > 0 {
				e.springs = append(e.springs, unknown)
			}
			e.springs = append(e.springs, r.springs...)
		}
		for i := 0; i < 5; i++ {
			e.counts = append(e.counts, r.counts...)
		}
		sum += countUnfoldedArrangements(e.springs, e.counts)
	}

	return sum
}

func main() {
	fmt.Println("Advent of Code 2023, Day 12")

	input := "input.txt"

	lines, err := utils.ReadLines(input)
	if err != nil {
		panic(err)
	}

	records := parse(lines)

	sum := 0
	for _, rec := range records {
		sum += rec.countArrangements()
	}

	fmt.Println("Part 1: the total number of possible arrangements is", sum)
	fmt.Println("Part 2: the total number of possible arrangements is", part2(input))
}
