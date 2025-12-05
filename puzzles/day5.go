package puzzles

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
	"log"
)

func prepareFreshRanges(fresh string) [][2]uint64 {
	var freshRanges [][2]uint64
	
	for idRange := range strings.SplitSeq(fresh, "\n") {
		rangeParts := strings.Split(idRange, "-")
		low, err := strconv.ParseUint(rangeParts[0], 10, 64)
		high, err := strconv.ParseUint(rangeParts[1], 10, 64)
		if err != nil {
			log.Fatal("Failed to convert to int")
		}

		freshRanges = append(freshRanges, [2]uint64{low, high})
	}

	// asc order
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})
	return freshRanges
}


func mergeFreshRanges(freshRanges [][2]uint64) [][2]uint64 {
	var newRange [][2]uint64
	currentRange := freshRanges[0]

	for i := 1; i < len(freshRanges); i++ {
		nextRange := freshRanges[i]
		if nextRange[0] <= currentRange[1] {
			currentRange[1] = max(nextRange[1], currentRange[1])
		} else {
			newRange = append(newRange, currentRange)
			currentRange = nextRange
		}
	}

	newRange = append(newRange, currentRange)
	return newRange
}

func countFresh(idList string) uint64 {
	parts := strings.Split(idList, "\n\n")
	freshRanges := prepareFreshRanges(strings.TrimSpace(parts[0]))

	freshRanges = mergeFreshRanges(freshRanges)

	count := uint64(0)
	for _, freshRange := range freshRanges {
		count += freshRange[1] - freshRange[0] + 1
	}

	return count
}

func Day5() {
	input := loadInput("day5.txt")
	fmt.Println("Day 5", countFresh(input))
}
