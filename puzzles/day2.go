package puzzles

import (
	"os"
	"io"
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

func sumRepeats(idRange string) uint64 {
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

	return sum
}

func sumAllRepeats(input string) uint64 {
	sum := uint64(0)

	for idRange := range strings.SplitSeq(input, ",") {
		sum += sumRepeats(strings.TrimSpace(idRange))
	}

	return sum
}

func Day2() {
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatalf("Failed to open input file")
	}

	input, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to open input file")
	}

	fmt.Println("Day2", sumAllRepeats(string(input)))
}
