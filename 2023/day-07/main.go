package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

type htype int

const (
	highCard     htype = iota
	onePair      htype = iota
	twoPair      htype = iota
	threeOfAKind htype = iota
	fullHouse    htype = iota
	fourOfAKind  htype = iota
	fiveOfAKind  htype = iota
)

type hand struct {
	// A = 14, K = 13, Q = 12, J = 11, T = 10
	cards []int
	bid   int
	t     htype
}

func handType(cards []int) htype {
	a := append([]int{}, cards...)
	slices.Sort(a)
	if a[0] == a[4] {
		return fiveOfAKind
	}
	if a[0] == a[3] || a[1] == a[4] {
		return fourOfAKind
	}
	if (a[0] == a[2] && a[3] == a[4]) || (a[0] == a[1] && a[2] == a[4]) {
		return fullHouse
	}
	if a[0] == a[2] || a[1] == a[3] || a[2] == a[4] {
		return threeOfAKind
	}
	if (a[0] == a[1] && a[2] == a[3]) || (a[0] == a[1] && a[3] == a[4]) || (a[1] == a[2] && a[3] == a[4]) {
		return twoPair
	}
	if a[0] == a[1] || a[1] == a[2] || a[2] == a[3] || a[3] == a[4] {
		return onePair
	}
	return highCard
}

func handTypeWithJokers(cards []int) htype {
	if !slices.Contains(cards, 1) {
		return handType(cards)
	}
	a := append([]int{}, cards...)
	slices.Sort(a)
	j := 0
	var c []int
	var n []int
	for i := 0; i < 5; i++ {
		ci := a[i]
		if ci == 1 {
			j += 1
		} else {
			idx := slices.Index(c, ci)
			if idx == -1 {
				c = append(c, ci)
				n = append(n, 1)
			} else {
				n[idx] += 1
			}
		}
	}
	u := len(c)
	if j == 5 || u == 1 {
		return fiveOfAKind
	}
	if u == 2 {
		if n[0] == 2 && n[1] == 2 {
			return fullHouse
		}
		return fourOfAKind
	}
	if u == 3 {
		return threeOfAKind
	}
	return onePair
}

func main() {
	fmt.Println("Advent of Code 2023, Day 7")

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	hands := make([]hand, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		h := &hands[i]
		h.cards = make([]int, 0, 5)
		for _, c := range parts[0] {
			switch c {
			case 'A':
				{
					h.cards = append(h.cards, 14)
				}
			case 'K':
				{
					h.cards = append(h.cards, 13)
				}
			case 'Q':
				{
					h.cards = append(h.cards, 12)
				}
			case 'J':
				{
					h.cards = append(h.cards, 11)
				}
			case 'T':
				{
					h.cards = append(h.cards, 10)
				}
			default:
				{
					h.cards = append(h.cards, int(c-'0'))
				}
			}
		}
		h.bid, _ = strconv.Atoi(parts[1])
		h.t = handType(h.cards)
	}

	slices.SortFunc(hands, func(a, b hand) int {
		ac := a.cards
		bc := b.cards
		atype := a.t
		btype := b.t
		if atype != btype {
			return int(atype) - int(btype)
		}
		for i := 0; i < 5; i++ {
			if ac[i] != bc[i] {
				return ac[i] - bc[i]
			}
		}
		return 0
	})

	sum := 0

	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	fmt.Println("Part 1: the total winnings are", sum)

	hands = make([]hand, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		h := &hands[i]
		h.cards = make([]int, 0, 5)
		for _, c := range parts[0] {
			switch c {
			case 'A':
				{
					h.cards = append(h.cards, 14)
				}
			case 'K':
				{
					h.cards = append(h.cards, 13)
				}
			case 'Q':
				{
					h.cards = append(h.cards, 12)
				}
			case 'J':
				{
					h.cards = append(h.cards, 1)
				}
			case 'T':
				{
					h.cards = append(h.cards, 10)
				}
			default:
				{
					h.cards = append(h.cards, int(c-'0'))
				}
			}
		}
		h.bid, _ = strconv.Atoi(parts[1])
		h.t = handTypeWithJokers(h.cards)
	}

	slices.SortFunc(hands, func(a, b hand) int {
		ac := a.cards
		bc := b.cards
		if a.t != b.t {
			return int(a.t) - int(b.t)
		}
		for i := 0; i < 5; i++ {
			if ac[i] != bc[i] {
				return ac[i] - bc[i]
			}
		}
		return 0
	})

	sum = 0

	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	fmt.Println("Part 2: the total winnings with jokers are", sum)
}
