package main

import (
	"container/list"
	"fmt"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func resolveConjunctionPulse(conj map[string]bool) bool {
	for _, v := range conj {
		if !v {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("***    Advent of Code 2023    ***")
	fmt.Println("--- Day 20: Pulse Propagation ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("unable to read the input file")
	}

	start := time.Now()

	graph := make(map[string][]string)
	flipflops := make(map[string]bool)
	conjunctions := make(map[string]map[string]bool)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		dests := strings.Split(parts[1], ", ")
		if parts[0] == "broadcaster" {
			graph[parts[0]] = dests
		} else {
			name := parts[0][1:]
			graph[name] = dests
			if parts[0][0] == '%' {
				flipflops[name] = false
			} else if parts[0][0] == '&' {
				conjunctions[name] = make(map[string]bool)
			}
		}
	}

	// Set the conjunctions to their initial values for each input.
	for src, dests := range graph {
		for _, dest := range dests {
			if _, ok := conjunctions[dest]; ok {
				conjunctions[dest][src] = false
			}
		}
	}

	lows, highs := 0, 0
	for i := 0; i < 1000; i++ {
		q := list.New()
		q.PushBack([]interface{}{"button", "broadcaster", false})
		for q.Len() > 0 {
			el := q.Front()
			sender := el.Value.([]interface{})[0].(string)
			mod := el.Value.([]interface{})[1].(string)
			pulse := el.Value.([]interface{})[2].(bool)
			q.Remove(el)

			// For all of the modules, true denotes high pulse and false
			// denotes low pulse.
			if pulse {
				highs += 1
			} else {
				lows += 1
			}

			if _, ok := flipflops[mod]; ok {
				if !pulse {
					flipflops[mod] = !flipflops[mod]
					newPulse := flipflops[mod]
					for _, dest := range graph[mod] {
						q.PushBack([]interface{}{mod, dest, newPulse})
					}
				}
			} else if _, ok := conjunctions[mod]; ok {
				conjunctions[mod][sender] = pulse
				newPulse := !resolveConjunctionPulse(conjunctions[mod])
				for _, dest := range graph[mod] {
					q.PushBack([]interface{}{mod, dest, newPulse})
				}
			} else if mod == "broadcaster" {
				for _, dest := range graph[mod] {
					q.PushBack([]interface{}{mod, dest, pulse})
				}
			}
		}
	}

	prod := lows * highs

	fmt.Println("Part 1: the product of the low and high pulses is", prod)
	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)
}
