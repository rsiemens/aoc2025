package main

import (
	"aoc2025/puzzles"
	"fmt"
	"time"
)

func timed(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("    in", elapsed)
}

func main() {
	timed(puzzles.Day1)
	timed(puzzles.Day2)
	timed(puzzles.Day3)
	timed(puzzles.Day4)
	timed(puzzles.Day5)
}
