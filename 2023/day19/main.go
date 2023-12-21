package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type comp byte
type category rune

const (
	LessThan comp = iota
	GreaterThan
	// This denotes that the rule only unconditionally either accepts, rejects
	// or sends the part to another workflow.
	Send
)

type rule struct {
	Category    category
	Comparator  comp
	Rating      int
	Destination string
}

type part struct {
	Ratings map[category]int
}

func (p *part) applyRules(wfs map[string][]rule) bool {
	wf := wfs["in"]
	for {
		for _, r := range wf {
			if r.Comparator == Send {
				d := r.Destination
				if d == "A" {
					return true
				} else if d == "R" {
					return false
				} else {
					wf = wfs[d]
				}
				break
			} else {
				c := r.Category
				if r.Comparator == LessThan && p.Ratings[c] < r.Rating {
					if r.Destination == "A" {
						return true
					} else if r.Destination == "R" {
						return false
					}
					wf = wfs[r.Destination]
				} else if r.Comparator == GreaterThan && p.Ratings[c] > r.Rating {
					if r.Destination == "A" {
						return true
					} else if r.Destination == "R" {
						return false
					}
					wf = wfs[r.Destination]
				} else {
					continue
				}
				break
			}
		}
	}
}

func resolveParts(wfs map[string][]rule, parts []part) []part {
	accepted := make([]part, 0)
	for _, part := range parts {
		if part.applyRules(wfs) {
			accepted = append(accepted, part)
		}
	}
	return accepted
}

func main() {
	fmt.Println("*** Advent of Code 2023 ***")
	fmt.Println("---   Day 19: Aplenty   ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("failed to read the input file")
	}

	start := time.Now()

	// This parsing is a mess but I guess it works
	workflows := make(map[string][]rule)
	parts := make([]part, 0)
	workflowsParsed := false
	for _, line := range lines {
		if line == "" {
			workflowsParsed = true
			continue
		}

		if !workflowsParsed {
			firstParts := strings.Split(line, "{")
			name := firstParts[0]
			rules := make([]rule, 0)
			rulesString := strings.TrimSuffix(firstParts[1], "}")
			for _, rs := range strings.Split(rulesString, ",") {
				var r rule
				var p []string
				if strings.Contains(rs, "<") {
					r.Comparator = LessThan
					p = strings.FieldsFunc(rs, func(r rune) bool {
						return r == '<' || r == ':'
					})
				} else if strings.Contains(rs, ">") {
					r.Comparator = GreaterThan
					p = strings.FieldsFunc(rs, func(r rune) bool {
						return r == '>' || r == ':'
					})
				} else {
					r.Comparator = Send
					r.Destination = rs
					rules = append(rules, r)
					continue
				}
				r.Category = category(p[0][0])
				r.Rating, _ = strconv.Atoi(p[1])
				r.Destination = p[2]
				rules = append(rules, r)
			}
			workflows[name] = rules
		} else {
			s := line[1 : len(line)-1]
			p := part{make(map[category]int)}
			for _, a := range strings.Split(s, ",") {
				ps := strings.Split(a, "=")
				c := category(ps[0][0])
				r, _ := strconv.Atoi(ps[1])
				p.Ratings[c] = r
			}
			parts = append(parts, p)
		}
	}

	accepted := resolveParts(workflows, parts)
	sum := 0
	for _, part := range accepted {
		for _, v := range part.Ratings {
			sum += v
		}
	}

	fmt.Println("Part 1: the sum of the accepted parts' rating numbers is", sum)
	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	sum2 := part2(inputFile)
	fmt.Println("Part 2: the sum of the accepted parts' rating numbers is", sum2)
	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
