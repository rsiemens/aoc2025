package puzzles

import (
	"fmt"
	"strings"
)


type beam struct {
	x int
	y int
}

type splitter struct {
	x int
	y int
}

func (s *splitter) split() (beam, beam) {
	lhs := beam{s.x - 1, s.y + 1}
	rhs := beam{s.x + 1, s.y + 1}
	return lhs, rhs
}

type splitterMap map[[2]int]splitter

var cache = make(map[[2]int]int)

func countTimelines(beam_ beam, splitters splitterMap, height int) int {
	if beam_.y >= height {
		return 1
	}

	coords := [2]int{beam_.x, beam_.y}
	if splitter, ok := splitters[coords]; ok {
		if hit, ok := cache[coords]; ok {
			//fmt.Printf("Cache hit of %d on coords %v\n", hit, coords)
			return hit
		}

		lhs, rhs := splitter.split()
		result := countTimelines(lhs, splitters, height) + countTimelines(rhs, splitters, height)
		cache[coords] = result
		return result
	}

	// propagate the beam
	beam_.y++
	return countTimelines(beam_, splitters, height)
}

func parseDiagram(diagram string) (beam, splitterMap, int) {
	var beamStart beam
	splitters := make(splitterMap)
	rows := strings.Split(strings.TrimSpace(diagram), "\n")
	for y, row := range rows  {

		for x, col := range row {
			switch col {
			case 'S':
				beamStart = beam{x, y}
			case '^':
				splitters[[2]int{x, y}] = splitter{x, y} 
			}
		}
	}

	return beamStart, splitters, len(rows)
}

func Day7() {
	input := loadInput("day7.txt")
	beamStart, splitters, height := parseDiagram(input)
	fmt.Println("Day 7", countTimelines(beamStart, splitters, height))
}
