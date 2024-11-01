package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func isAdjacentToSymbol(schematic [][]rune, xStart, xEnd, y int) bool {
	// The differences in the coordinates for checking the adjacent symbols.
	// The same index should be used for both x and y differences.
	dx := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dy := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	difflen := 8

	// This loop could probably optimised to not check all of the coordinates
	// around as some of the adjacent coordinates are checked twice.
	for x := xStart; x <= xEnd; x++ {
		for i := 0; i < difflen; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if ny >= 0 && nx >= 0 && ny < len(schematic) && nx < len(schematic[ny]) {
				r := schematic[ny][nx]
				if r != '.' && !unicode.IsDigit(r) {
					return true
				}
			}
		}
	}
	return false
}

func sumPartNumbers(schematic [][]rune) int {
	sum := 0
	for y, line := range schematic {
		num, startX := "", 0
		for x, r := range line {
			if unicode.IsDigit(r) {
				if num == "" {
					startX = x
				}
				num += string(r)

				if x == len(line)-1 || !unicode.IsDigit(line[x+1]) {
					if isAdjacentToSymbol(schematic, startX, x, y) {
						i, _ := strconv.Atoi(num)
						sum += i
					}
					num = ""
				}
			}
		}
	}
	return sum
}

func isCoordinateDigit(schematic [][]rune, x, y int) bool {
	if y < 0 || x < 0 || y >= len(schematic) || x >= len(schematic[y]) {
		return false
	}
	return unicode.IsDigit(schematic[y][x])
}

func scanLine(schematic [][]rune, x, y int) []int {
	var nums []int
	num := ""

	// Do this the stupid, simple way. Start by checking the rune at the same
	// x coordinate.
	if isCoordinateDigit(schematic, x, y) {
		num += string(schematic[y][x])

		if isCoordinateDigit(schematic, x-1, y) {
			for nx := x - 1; isCoordinateDigit(schematic, nx, y); nx-- {
				num = string(schematic[y][nx]) + num
			}
		}

		if isCoordinateDigit(schematic, x+1, y) {
			for nx := x + 1; isCoordinateDigit(schematic, nx, y); nx++ {
				num += string(schematic[y][nx])
			}
		}

		n, _ := strconv.Atoi(num)
		num = ""
		nums = append(nums, n)
	} else {
		if isCoordinateDigit(schematic, x-1, y) {
			for nx := x - 1; isCoordinateDigit(schematic, nx, y); nx-- {
				num = string(schematic[y][nx]) + num
			}
			n, _ := strconv.Atoi(num)
			num = ""
			nums = append(nums, n)
		}

		if isCoordinateDigit(schematic, x+1, y) {
			for nx := x + 1; isCoordinateDigit(schematic, nx, y); nx++ {
				num += string(schematic[y][nx])
			}
			n, _ := strconv.Atoi(num)
			num = ""
			nums = append(nums, n)
		}
	}

	return nums
}

func findAdjacentPartNumbers(schematic [][]rune, x, y int) []int {
	var nums []int

	// Let's make this simple.
	nums = append(nums, scanLine(schematic, x, y-1)...)
	nums = append(nums, scanLine(schematic, x, y+1)...)
	nums = append(nums, scanLine(schematic, x, y)...)

	return nums
}

func sumGearRatios(schematic [][]rune) int {
	sum := 0
	for y, line := range schematic {
		for x, r := range line {
			if r == '*' {
				nums := findAdjacentPartNumbers(schematic, x, y)
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Advent of Code 2023, Day 3")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	schematic := make([][]rune, len(lines))

	for i, line := range lines {
		schematic[i] = []rune(line)
	}

	sum := sumPartNumbers(schematic)
	fmt.Println("Part 1: the sum of the part numbers is", sum)

	sum = sumGearRatios(schematic)
	fmt.Println("Part 2: the sum of the gear ratios is", sum)
}
