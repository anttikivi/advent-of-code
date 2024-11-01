package main

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func Part1Quadratic() float64 {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1.0

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])
		det := math.Sqrt(float64(t*t - 4*d))
		tf := float64(t)
		l := math.Floor((tf - det) / 2)
		h := math.Ceil(((tf + det) / 2) - 1)
		p *= h - l
	}
	return p
}

func Part1QuadraticGoroutines() float64 {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1.0
	c := make(chan float64)

	for i, time := range times {
		go func(i int, time string) {
			t, _ := strconv.Atoi(time)
			d, _ := strconv.Atoi(distances[i])
			det := math.Sqrt(float64(t*t - 4*d))
			tf := float64(t)
			l := math.Floor((tf - det) / 2)
			h := math.Ceil(((tf + det) / 2) - 1)
			c <- h - l
		}(i, time)
	}

	for range times {
		p *= <-c
	}

	close(c)

	return p
}

func Part1Loop() int {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])
		w := 0

		for x := 0; x <= t; x++ {
			dist := t*x - x*x
			if dist > d {
				w += 1
			}
		}
		p *= w
	}
	return p
}

func Part1Loop2() int {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])
		w := 0

		for x := 0; x <= t; x++ {
			tleft := t - x
			dist := tleft * x
			if dist > d {
				w += 1
			}
		}
		p *= w
	}
	return p
}

func Part1QuadraticGoroutines2() float64 {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1.0
	c := make(chan float64, len(times))

	for i, time := range times {
		go func(c chan<- float64) {
			t, _ := strconv.Atoi(time)
			d, _ := strconv.Atoi(distances[i])
			det := math.Sqrt(float64(t*t - 4*d))
			tf := float64(t)
			l := math.Floor((tf - det) / 2)
			h := math.Ceil(((tf + det) / 2) - 1)
			c <- h - l
		}(c)
		p *= <-c
	}
	return p
}

func Part1LoopGoroutines() int {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1

	for i, time := range times {
		c := make(chan int)
		go func(c chan<- int) {
			t, _ := strconv.Atoi(time)
			d, _ := strconv.Atoi(distances[i])
			w := 0

			for x := 0; x <= t; x++ {
				dist := t*x - x*x
				if dist > d {
					w += 1
				}
			}
			c <- w
		}(c)
		p *= <-c
	}
	return p
}

func Part1LoopGoroutines2() int {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1
	c := make(chan int, len(times))

	for i, time := range times {
		go func(c chan<- int) {
			t, _ := strconv.Atoi(time)
			d, _ := strconv.Atoi(distances[i])
			w := 0

			for x := 0; x <= t; x++ {
				dist := t*x - x*x
				if dist > d {
					w += 1
				}
			}
			c <- w
		}(c)
		p *= <-c
	}
	return p
}

func Part1LoopTwoGoroutines() int {
	lines := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])
	p := 1
	c := make(chan int, len(times))

	for i, time := range times {
		go func(c chan<- int) {
			var wg sync.WaitGroup
			var mu sync.Mutex
			t, _ := strconv.Atoi(time)
			d, _ := strconv.Atoi(distances[i])
			w := 0
			for x := 0; x <= t; x++ {
				wg.Add(1)
				go func(x int) {
					defer wg.Done()
					dist := t*x - x*x
					if dist > d {
						mu.Lock()
						defer mu.Unlock()
						w += 1
					}
				}(x)
			}
			wg.Wait()
			c <- w
		}(c)
		p *= <-c
	}
	return p
}

func TestPart1(t *testing.T) {
	want := 138915
	got := Part1Quadratic()
	if got != float64(want) {
		t.Errorf("Part1Quadratic() = %f, want %d", got, want)
	}
	got2 := Part1QuadraticGoroutines()
	if got2 != float64(want) {
		t.Errorf("Part1QuadraticGoroutines() = %f, want %d", got2, want)
	}
	got3 := Part1Loop()
	if got3 != want {
		t.Errorf("Part1Loop() = %d, want %d", got3, want)
	}
	got4 := Part1Loop2()
	if got4 != want {
		t.Errorf("Part1Loop2() = %d, want %d", got4, want)
	}
	got5 := Part1QuadraticGoroutines2()
	if got5 != float64(want) {
		t.Errorf("Part1QuadraticGoroutines2() = %f, want %d", got5, want)
	}
	got6 := Part1LoopGoroutines()
	if got6 != want {
		t.Errorf("Part1LoopGoroutines() = %d, want %d", got6, want)
	}
	got7 := Part1LoopGoroutines2()
	if got7 != want {
		t.Errorf("Part1LoopGoroutines2() = %d, want %d", got7, want)
	}
	got8 := Part1LoopTwoGoroutines()
	if got8 != want {
		t.Errorf("Part1LoopTwoGoroutines() = %d, want %d", got8, want)
	}
}

func BenchmarkPart1Quadratic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1Quadratic()
	}
}

func BenchmarkPart1QuadraticGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1QuadraticGoroutines()
	}
}

func BenchmarkPart1Loop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1Loop()
	}
}

func BenchmarkPart1Loop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1Loop2()
	}
}

func BenchmarkPart1QuadraticGoroutines2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1QuadraticGoroutines2()
	}
}

func BenchmarkPart1LoopGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1LoopGoroutines()
	}
}

func BenchmarkPart1LoopGoroutines2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1LoopGoroutines2()
	}
}

func BenchmarkPart1LoopTwoGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1LoopTwoGoroutines()
	}
}

func Part2Quadratic(t int, d int) int {
	det := math.Sqrt(float64(t*t - 4*d))
	tf := float64(t)
	l := math.Floor((tf - det) / 2)
	h := math.Ceil(((tf + det) / 2) - 1)
	return int(h - l)
}

func Part2Loop(t int, d int) int {
	w := 0

	for x := 0; x <= t; x++ {
		dist := t*x - x*x
		if dist > d {
			w += 1
		}
	}

	return w
}

func TestPart2(t *testing.T) {
	lines := strings.Split(input, "\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	dist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))

	want := 27340847
	got := Part2Quadratic(time, dist)
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
	got2 := Part2Loop(time, dist)
	if got2 != want {
		t.Errorf("Part2Loop() = %d, want %d", got2, want)
	}
	// Don't do goroutines for loops in part 2, it seems to be very slow.
}

func BenchmarkPart2Quadratic(b *testing.B) {
	lines := strings.Split(input, "\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	dist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	for i := 0; i < b.N; i++ {
		Part2Quadratic(time, dist)
	}
}

func BenchmarkPart2Loop(b *testing.B) {
	lines := strings.Split(input, "\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	dist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	for i := 0; i < b.N; i++ {
		Part2Loop(time, dist)
	}
}
