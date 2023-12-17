package main

import (
	"fmt"
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

func main() {
	fmt.Println("Advent of Code 2023, Day 12")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

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

	sum := 0
	for _, rec := range records {
		sum += rec.countArrangements()
	}

	fmt.Println("Part 1: the total number of possible arrangements is", sum)
}
