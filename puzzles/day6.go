package puzzles

import (
	"fmt"
	"log"
	"strings"
	"strconv"
)

func mult(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func columnOrdered(strNums []string) []int {
	var result []int
	maxLen := len(strNums[0])

	for i := range maxLen {
		var cur string
		for _, strNum := range strNums {
			// 0, 1, 2
			// "123" " 45" "  6"
			// cur = ""

			// 3 - 0 - 1 = 2
			// cur = "3"

			// 3 - 0 - 1 = 2
			// cur = "5"

			if len(strNum) > i {
				cur += string(strNum[len(strNum) - i - 1])
			}
		}
		num, err := strconv.Atoi(strings.TrimSpace(cur))
		if err != nil { log.Fatal("failed to convert to int") }
		result = append(result, num)
	}

	return result
}

func evaluateHomework(hw [][]string) int {
	rows := len(hw)
	cols := len(hw[0])

	sum := 0
	for col := range cols {
		var op func(a, b int) int

		switch strings.TrimSpace(hw[rows - 1][col]) {
		case "*":
			op = mult
		case "+":
			op = add
		}

		var strNums[]string
		for row := 0; row < rows - 1; row++ {
			strNums = append(strNums, hw[row][col])
		}

		nums := columnOrdered(strNums)
		result := nums[0]
		nums = nums[1:]
		for len(nums) > 0 {
			next := nums[0]
			nums = nums[1:]
			result = op(result, next)
		}
		sum += result
	}

	return sum
}

func prepareHomework(input string) [][]string {
	var result [][]string

	rows := strings.Split(strings.Trim(input, "\n"), "\n")
	splitPoints := findSplitPoints(rows)

	for _, row := range rows {
		lastPoint := 0
		var stack []string

		for _, point := range splitPoints {
			stack = append(stack, row[lastPoint:point])
			lastPoint = point + 1
		}

		stack = append(stack, row[lastPoint:])
		result = append(result, stack)
	}

	return result
}

func findSplitPoints(rows []string) []int {
	var points []int

	nCols := len(rows[0])
	for col := range nCols {
		allSpace := true
		for row := range len(rows) {
			if string(rows[row][col]) != " " {
				allSpace = false
			}
		}

		if allSpace {
			points = append(points, col)
		}
	}

	return points
}

func Day6() {
	input := loadInput("day6.txt")
	preped := prepareHomework(input)
	fmt.Println("Day 6", evaluateHomework(preped))
}
