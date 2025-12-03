package puzzles

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func findMaxWithRemaining(bank string, from, until int) int {
	curMaxPos := from
	curMax, err := strconv.Atoi(string(bank[from]))
	if err != nil { log.Fatalf("Failed to parse int") }

	for i := from; i < until; i++ {
		curVal, err := strconv.Atoi(string(bank[i]))
		if err != nil { log.Fatalf("Failed to parse int") }

		if curVal > curMax {
			curMax = curVal
			curMaxPos = i
		}
	}

	return curMaxPos
}

func maxJoltage(bank string) uint64 {
	joltage := ""
	from := 0
	until := len(bank) - 11

	for len(joltage) < 12 {
		from = findMaxWithRemaining(bank, from, until)
		joltage += string(bank[from])
		until++
		from++
	}

	res, err := strconv.ParseUint(joltage, 10, 64)
	if err != nil { log.Fatalf("Failed to parse int") }

	return res
}

func totalJoltage(banks string) uint64 {
	sum := uint64(0)

	for bank := range strings.SplitSeq(banks, "\n") {
		sum += maxJoltage(bank)
	}

	return sum
}

func Day3() {
	input := loadInput("day3.txt")
	banks := strings.TrimSpace(input)
	fmt.Println("Day3", totalJoltage(banks))
}
