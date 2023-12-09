package main

import (
	"fmt"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func main() {
	fmt.Println("Advent of Code 2023, Day 9")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
}
