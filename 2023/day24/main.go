package main

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/anttikivi/advent-of-code/2023/utils"
)

const (
	lowerLimit = 200000000000000
	upperLimit = 400000000000000
)

type hailstone struct {
	Px, Py, Pz, Vx, Vy, Vz int
}

func calculateDeterminant(m [][]*big.Int) *big.Int {
	if len(m) == 0 {
		return big.NewInt(1)
	}

	l := m[0]
	r := m[1:]
	rProduct := make([]*big.Int, 0)

	for i, n := range l {
		var spliced [][]*big.Int
		for _, row := range r {
			var newRow []*big.Int
			newRow = append(make([]*big.Int, 0), row[:i]...)
			newRow = append(newRow, row[i+1:]...)
			spliced = append(spliced, newRow)
		}
		rProduct = append(
			rProduct,
			new(big.Int).Mul(n, calculateDeterminant(spliced)),
		)
	}

	result := big.NewInt(0)

	for i, b := range rProduct {
		if i%2 == 0 {
			result.Add(result, b)
		} else {
			result.Sub(result, b)
		}
	}

	return result
}

func main() {
	fmt.Println("*****    Advent of Code 2023     *****")
	fmt.Println("--- Day 24: Never Tell Me The Odds ---")

	inputFile := "input.txt"
	lines, err := utils.ReadLines(inputFile)
	if err != nil {
		panic("Failed to read the input file")
	}

	start := time.Now()

	hailstones := make([]hailstone, 0)
	for _, line := range lines {
		parts := strings.Split(line, " @ ")
		p := strings.Split(parts[0], ", ")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		pz, _ := strconv.Atoi(p[2])
		v := strings.Split(parts[1], ", ")
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		vz, _ := strconv.Atoi(v[2])
		stone := hailstone{px, py, pz, vx, vy, vz}
		hailstones = append(hailstones, stone)
	}

	count := 0
	for i, stoneA := range hailstones {
		for _, stoneB := range hailstones[i+1:] {
			dx := stoneA.Px - stoneB.Px
			dy := stoneA.Py - stoneB.Py
			denominator := stoneA.Vy*stoneB.Vx - stoneA.Vx*stoneB.Vy
			if denominator == 0 {
				continue
			}
			numerator := stoneB.Vy*dx - stoneB.Vx*dy
			tA := numerator / denominator
			if tA <= 0 {
				continue
			}
			numerator = stoneA.Vy*dx - stoneA.Vx*dy
			tB := numerator / denominator
			if tB <= 0 {
				continue
			}
			x := stoneA.Px + tA*stoneA.Vx
			y := stoneA.Py + tA*stoneA.Vy
			if x < lowerLimit || x > upperLimit || y < lowerLimit || y > upperLimit {
				continue
			}
			count++
		}
	}

	fmt.Println("Part 1: a total of", count, "intersections happen within the test area")

	elapsed := time.Since(start)
	fmt.Println("Part 1 ran for", elapsed)

	start = time.Now()

	var stones [][]*big.Int
	for _, line := range lines {
		numbers := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
		var vals []*big.Int
		for _, num := range numbers {
			n := new(big.Int)
			n.SetString(num, 10)
			vals = append(vals, n)
		}
		stones = append(stones, vals)
	}

	a := make([][]*big.Int, 0)
	b := make([]*big.Int, 0)

	var (
		px0 = stones[0][0]
		py0 = stones[0][1]
		pz0 = stones[0][2]
		vx0 = stones[0][3]
		vy0 = stones[0][4]
		vz0 = stones[0][5]
	)

	for i := 1; i <= 3; i++ {
		var (
			pxN = stones[i][0]
			pyN = stones[i][1]
			pzN = stones[i][2]
			vxN = stones[i][3]
			vyN = stones[i][4]
			vzN = stones[i][5]
		)

		a = append(a, []*big.Int{
			new(big.Int).Sub(vy0, vyN),
			new(big.Int).Sub(vxN, vx0),
			big.NewInt(0),
			new(big.Int).Sub(pyN, py0),
			new(big.Int).Sub(px0, pxN),
			big.NewInt(0),
		})

		t1 := new(big.Int).Mul(px0, vy0)
		t2 := new(big.Int).Mul(py0, vx0)
		t3 := new(big.Int).Mul(pxN, vyN)
		t4 := new(big.Int).Mul(pyN, vxN)

		// t1 - t2 - t3 + t4
		bN := new(big.Int).Sub(t1, t2)
		bN.Sub(bN, t3)
		bN.Add(bN, t4)
		b = append(b, bN)

		a = append(a, []*big.Int{
			new(big.Int).Sub(vz0, vzN),
			big.NewInt(0),
			new(big.Int).Sub(vxN, vx0),
			new(big.Int).Sub(pzN, pz0),
			big.NewInt(0),
			new(big.Int).Sub(px0, pxN),
		})

		t1 = new(big.Int).Mul(px0, vz0)
		t2 = new(big.Int).Mul(pz0, vx0)
		t3 = new(big.Int).Mul(pxN, vzN)
		t4 = new(big.Int).Mul(pzN, vxN)

		// t1 - t2 - t3 + t4
		bN = new(big.Int).Sub(t1, t2)
		bN.Sub(bN, t3)
		bN.Add(bN, t4)
		b = append(b, bN)
	}

	detA := calculateDeterminant(a)

	var result []*big.Int
	for i := range a {
		m := make([][]*big.Int, 0)
		for j, row := range a {
			newRow := make([]*big.Int, 0)
			newRow = append(newRow, row...)
			newRow[i] = b[j]
			m = append(m, newRow)
		}

		d := calculateDeterminant(m)
		result = append(result, new(big.Int).Div(d, detA))
	}

	sum := new(big.Int).Add(result[0], result[1])
	sum.Add(sum, result[2])

	fmt.Println("Part 2: the sum of coordinates of the initial position is", sum)

	elapsed = time.Since(start)
	fmt.Println("Part 2 ran for", elapsed)
}
