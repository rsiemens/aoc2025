package puzzles

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func rotate(rot, dial int) (int, int) {
	newDial := (dial + rot) % 100
	if newDial < 0 {
		newDial += 100
	}
	clicks := abs(rot) / 100
	amount := abs(rot) - clicks*100
	passZero := false

	if rot < 0 && dial != 0 {
		passZero = dial-amount < 0
	} else if rot >= 0 {
		passZero = dial+amount > 100
	}

	if newDial == 0 {
		clicks += 1
	}

	if passZero {
		return newDial, clicks + 1
	}
	return newDial, clicks

}

func getCode(input string, dial int) int {
	var (
		clicks, count, val int
		err                error
	)

	for rot := range strings.SplitSeq(input, "\n") {
		val, err = strconv.Atoi(rot[1:])
		if err != nil {
			log.Fatalf("Unable to convert %q to an int", rot)
		}

		if strings.HasPrefix(rot, "L") {
			val *= -1
		}

		dial, clicks = rotate(val, dial)
		count += clicks
	}

	return count
}

func Day1() {
	input := strings.Trim(loadInput("day1.txt"), "\n ")
	fmt.Println("Day1", getCode(string(input), 50))
}
