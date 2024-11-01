package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

func suffixToNumber(str string) (int, error) {
	if strings.HasSuffix(str, "one") {
		return 1, nil
	}

	if strings.HasSuffix(str, "two") {
		return 2, nil
	}

	if strings.HasSuffix(str, "three") {
		return 3, nil
	}

	if strings.HasSuffix(str, "four") {
		return 4, nil
	}

	if strings.HasSuffix(str, "five") {
		return 5, nil
	}

	if strings.HasSuffix(str, "six") {
		return 6, nil
	}

	if strings.HasSuffix(str, "seven") {
		return 7, nil
	}

	if strings.HasSuffix(str, "eight") {
		return 8, nil
	}

	if strings.HasSuffix(str, "nine") {
		return 9, nil
	}

	return -1, fmt.Errorf("invalid number string: %s", str)
}

func main() {
	fmt.Println("Advent of Code 2023, Day 1")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range lines {
		first, last := -1, -1

		for _, r := range line {
			if unicode.IsDigit(r) {
				if first == -1 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}

		lineValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		sum += lineValue
	}

	fmt.Println("Part 1: the sum of the calibration values is", sum)

	sum = 0

	for _, line := range lines {
		str, first, last := "", -1, -1

		for _, r := range line {
			if unicode.IsLetter(r) {
				str += string(r)

				num, err := suffixToNumber(str)
				if err == nil {
					if first == -1 {
						first = num
					}
					last = num
				}
			} else if unicode.IsDigit(r) {
				if first == -1 {
					first = int(r - '0')
				}
				last = int(r - '0')
				str = ""
			}
		}

		lineValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		sum += lineValue
	}

	fmt.Println("Part 2: the sum of the calibration values is", sum)
}
