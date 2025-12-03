package puzzles

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func hasRepeatingPattern(num string) bool {
	outer:
	for end := 1; end < len(num); end++ {
		if len(num) % end != 0 {
			continue	
		}

		pattern := num[:end]
		for step := 1; step * end < len(num); step++ {
			from := step * end
			to := from + end
			if num[from:to] != pattern {
				continue outer
			}
		}
		return true
	}

	return false
}

func sumRepeats(idRange string, ch chan uint64) {
	parts := strings.Split(idRange, "-")
	lhs, err := strconv.ParseUint(parts[0], 10, 64)
	rhs, err := strconv.ParseUint(parts[1], 10, 64)
	sum := uint64(0)

	if err != nil {
		log.Fatalf("Failed to convert %s to ints", idRange)
	}

	for i := lhs; i <= rhs; i++ {
		asStr := strconv.FormatUint(i, 10)
		size := len(asStr)

		if size%2 == 0 && asStr[:size/2] == asStr[size/2:] {
			sum += i
		} else if hasRepeatingPattern(asStr) {
			sum += i
		}
	}

	ch <- sum
}

func sumAllRepeats(input string) uint64 {
	sum := uint64(0)
	idRanges := strings.Split(input, ",")
	ch := make(chan uint64)

	for _, idRange := range idRanges {
		go sumRepeats(strings.TrimSpace(idRange), ch)
	}

	for range idRanges {
		sum += <-ch
	}

	return sum
}

func Day2() {
	input := loadInput("day2.txt")
	fmt.Println("Day2", sumAllRepeats(input))
}
