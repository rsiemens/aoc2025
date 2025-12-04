package puzzles

import (
	"fmt"
	"strings"
)

func hasEnoughSpace(x, y int, grid [][]rune) bool {
	checks := [8][2]int{
		{x - 1, y},  // left
		{x + 1, y},  // right
		{x, y - 1},  // above
		{x, y + 1},  // below
		{x - 1, y - 1},  // top left
		{x + 1, y - 1},  // top right
		{x - 1, y + 1},  // bottom left
		{x + 1, y + 1},  // bottom right
	}
	nCols := len(grid)
	nRows := len(grid[0])

	rolls := 0
	for _, check := range checks {
		x, y := check[0], check[1]

		if x >= 0 && x < nRows && y >= 0 && y < nCols && grid[y][x] == '@' {
			rolls += 1
			if rolls > 3 {
				return false
			}
		}
	}

	return true
}

func removalPass(grid [][]rune) int {
	removed := 0
	for y, row := range grid {
		for x, ch := range row {
			if ch == '@' && hasEnoughSpace(x, y, grid) {
				removed += 1
				grid[y][x] = '.'
			}
		}
	}

	return removed
}

func removeRolls(input string) int {
	var grid [][]rune
	for row := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		grid = append(grid, []rune(row))
	}

	removed := 0
	for {
		pass := removalPass(grid)
		if pass == 0 {
			return removed
		}
		removed += pass
	}
}

func Day4() {
	input := loadInput("day4.txt")
	fmt.Println("Day4", removeRolls(strings.TrimSpace(input)))
}

