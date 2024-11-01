package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type lens struct {
	label       string
	focalLength int
}

func hash(s string) int {
	n := 0
	for _, c := range s {
		n += int(c)
		n *= 17
		n %= 256
	}
	return n
}

func indexLens(s []lens, k string) int {
	for i, l := range s {
		if l.label == k {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Advent of Code 2023, Day 15")

	input := "input.txt"
	init, err := os.ReadFile(input)
	if err != nil {
		panic("failed to read the input")
	}

	start := time.Now()

	sum := 0
	for _, s := range strings.Split(strings.ReplaceAll(string(init), "\n", ""), ",") {
		sum += hash(s)
	}
	fmt.Println("Part 1: the sum of the result is", sum)

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran in", elapsed)

	start = time.Now()

	boxes := make([][]lens, 256)
	for _, s := range strings.Split(strings.ReplaceAll(string(init), "\n", ""), ",") {
		k := ""
		op := rune(0)
		focallen := ""
		for _, c := range s {
			if c == '-' || c == '=' {
				op = c
			} else if unicode.IsDigit(c) {
				focallen += string(c)
			} else {
				k += string(c)
			}
		}
		b := hash(k)

		if op == '-' {
			boxes[b] = slices.DeleteFunc(boxes[b], func(l lens) bool { return l.label == k })
		} else if op == '=' {
			fl, _ := strconv.Atoi(focallen)
			if slices.ContainsFunc(boxes[b], func(l lens) bool { return l.label == k }) {
				i := indexLens(boxes[b], k)
				boxes[b][i].focalLength = fl
			} else {
				boxes[b] = append(boxes[b], lens{k, fl})
			}
		}
	}

	fp := 0
	for i, box := range boxes {
		for j, lens := range box {
			fp += (i + 1) * (j + 1) * lens.focalLength
		}
	}

	fmt.Println("Part 2: the sum of the result is", fp)

	elapsed = time.Since(start)
	fmt.Println("Part 1 ran in", elapsed)
}
