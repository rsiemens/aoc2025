package puzzles

import (
	"log"
	"fmt"
	"strings"
	"strconv"
	"math"
	"math/rand"
)


type machine struct {
	desiredState []int
	currentState []int
	buttons [][]int
	presses int
}

func (m *machine) String() string {
	return fmt.Sprintf("current=%v desired=%v presses=%d", m.currentState, m.desiredState, m.presses)
}

func (m *machine) needsReset() bool {
	for i := range m.currentState {
		if m.currentState[i] > m.desiredState[i] {
			return true
		}
	}
	return false
}

func (m *machine) reset() {
	for i := range m.currentState {
		m.currentState[i] = 0
	}
	m.presses = 0
}

func (m *machine) isCorrect() bool {
	for i := range m.currentState {
		if m.currentState[i] != m.desiredState[i] {
			return false
		}
	}

	return true
}

func (m *machine) push(n int) {
	btn := m.buttons[n]
	for _, val := range btn {
		m.currentState[val]++
	}
	m.presses++
}

func (m *machine) randomPush() {
	choice := rand.Intn(len(m.buttons))
	m.push(choice)
}

func (m *machine) distancePush() {
	distances := make([]int, len(m.desiredState))
	maxDistance := 0
	position := 0

	for i, desired := range m.desiredState {
		dist := desired - m.currentState[i]
		distances[i] = dist
		if dist > maxDistance {
			maxDistance = dist
			position = i
		}
	}

	var choices []int
	for i, vals := range m.buttons {
		hasPosition := false
		maintainsDistance := true
		for _, val := range vals {
			if distances[val] - 1 < 0 {
				maintainsDistance = false
				break
			}

			if val == position {
				hasPosition = true
			}
		}
		if hasPosition && maintainsDistance {
			choices = append(choices, i)
		}
	}

	if len(choices) >= 1 {
		choice := rand.Intn(len(choices))
		m.push(choices[choice])
	} else {
		// were screwed
		m.randomPush()
	}
}

func randomDistanceStrategy(machines []*machine, iterations int) int {
	total := 0

	for _, mach := range machines {
		best := math.MaxInt
		for range iterations {
			for !mach.isCorrect() {
				if mach.needsReset() {
					mach.reset()
				}
				mach.distancePush()
				//fmt.Println(mach)
			}
			fmt.Println("Solved", mach)
			if mach.presses < best {
				best = mach.presses
			}
			mach.reset()
		}
		total += best
	}

	return total
}

func parseMachines(input string) []*machine {
	var machines []*machine

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		var mach *machine = &machine{}

		cur := 0
		for cur < len(line) {
			char := string(line[cur])
			switch char {
			//case "[":
			//	cur++
			//	mach.desiredState, mach.currentState = parseLightState(&cur, line)
			case "(":
				cur++
				mach.buttons = append(mach.buttons, parseButton(&cur, line))
			case "{":
				cur++
				mach.desiredState, mach.currentState = parseJoltageState(&cur, line)
			default:
				cur++
			}
		}

		machines = append(machines, mach)
	}

	return machines 
}

func parseLightState(cur *int, line string) ([]string, []string) {
	var desiredState []string
	var currentState []string

	char := string(line[*cur])
	for char != "]" {
		*cur++
		desiredState = append(desiredState, char)
		currentState = append(currentState, ".")
		char = string(line[*cur])
	}
	*cur++

	return desiredState, currentState
}

func parseButton(cur *int, line string) []int {
	var button []int

	char := string(line[*cur])
	for char != ")" {
		*cur++
		if char != "," {
			n, err := strconv.Atoi(char)
			if err != nil { log.Fatalf("Failed to parse %q", char) }
			button = append(button, n)
		}
		char = string(line[*cur])
	}
	*cur++

	return button
}

func parseJoltageState(cur *int, line string) ([]int, []int) {
	var desiredState []int
	var currentState []int

	char := string(line[*cur])
	for char != "}" {
		if char != "," {
			n := parseInt(cur, line)
			desiredState = append(desiredState, n)
			currentState = append(currentState, 0)
		}
		if string(line[*cur]) == "," {
			*cur++
		}
		char = string(line[*cur])
	}
	*cur++

	return desiredState, currentState
}

func parseInt(cur *int, line string) int {
	n := ""
	char := string(line[*cur])
	for char != "}" && char != "," {
		n += char
		*cur++
		char = string(line[*cur])
	}
	rv, err := strconv.Atoi(n)
	if err != nil { log.Fatalf("Failed to parse %q", n) }
	return rv
}


func Day10() {
	input := loadInput("day10.txt")
	machines := parseMachines(input)
	fmt.Println("Day 10", randomDistanceStrategy(machines, 1))
}
