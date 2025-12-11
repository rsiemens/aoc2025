package puzzles

import (
	"strings"
	"fmt"
	"log"
)

type coord struct {
	x int
	y int
}

func (c *coord) area(other *coord) int {
	return (abs(c.x - other.x) +1) * (abs(c.y - other.y)+1)
}

func pointInBounds(point [2]int, coords []*coord) bool {
	crossings := 0

	for i := range coords {
		p1 := coords[i]
		p2 := coords[(i + 1) % len(coords)]

		// on boundary horizontal
		if p1.y == p2.y && point[1] == p1.y && min(p1.x, p2.x) <= point[0] && max(p1.x, p2.x) >= point[0] {
			return true
		// on boundary vertical
		} else if p1.x == p2.x && point[0] == p1.x && min(p1.y, p2.y) <= point[1] && max(p1.y, p2.y) >= point[1] {
			return true
		}

		if p1.x == p2.x && min(p1.y, p2.y) <= point[1] && max(p1.y, p2.y) >= point[1] {
			if point[0] < p1.x {
				crossings += 1
			}
		}
	}

	if crossings != 1 {
		return false
	}
	return true
}

func inBounds(coordA, coordB *coord, coords []*coord) bool {
	corners := [4][2]int{
		{coordA.x, coordA.y},
		{coordB.x, coordB.y},
		{coordA.x, coordB.y},
		{coordB.x, coordA.y},
	}

	/*
 012345678
0.........
1.#XXXXX#.
2.X.....X.
3.X.#X#.X.
4.X.X.X.X.
5.#X#.#X#.
 .........
	*/

	for _, corner := range corners {
		if !pointInBounds(corner, coords) {
			return false
		}
	}

	x := min(coordA.x, coordB.x) + 1
	for y := min(coordA.y, coordB.y); y < max(coordA.y, coordB.y); y++ {
		if !pointInBounds([2]int{x, y}, coords) {
			return false
		}
	}

	return true
}

func findMaxArea(coords []*coord) int {
	maxArea := 0

	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			coordA := coords[i]
			coordB := coords[j]
			area := coordA.area(coordB)

			if area > maxArea && inBounds(coordA, coordB, coords) {
				maxArea = area
			}
		}
	}
	return maxArea
}

func parseCoords(input string) []*coord {
	var coords []*coord

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		x, y := 0, 0
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			log.Fatalf("Failed to parse %q\n", line)
		}
		coords = append(coords, &coord{x, y})
	}

	return coords
}

func Day9() {
	input := loadInput("day9.txt")
	coords := parseCoords(input)
	fmt.Println("Day 9", findMaxArea(coords))
}
