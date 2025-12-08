package puzzles

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"sort"
)

type circuitMap map[int][]*junctionBox

func (c *circuitMap) merge(into, from int) {
	if into == from {
		return
	}

	intoBoxes, ok := (*c)[into]
	if !ok {
		return
	}

	fromBoxes, ok := (*c)[from]
	if !ok {
		return
	}

	for _, box := range fromBoxes {
		box.circuit = into
		intoBoxes = append(intoBoxes, box)
	}

	(*c)[into] = intoBoxes
	delete(*c, from)
}

type junctionBox struct {
	id int
	x float64
	y float64
	z float64
	circuit int
	connections map[int]*junctionBox
}

func (p *junctionBox) String() string {
	return fmt.Sprintf("<%g,%g,%g, circuit:%d>", p.x, p.y, p.z, p.circuit)
}

func (p *junctionBox) distance(other *junctionBox) float64 {
	return math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2) + math.Pow(p.z-other.z, 2)
}

func (p *junctionBox) isConnected(other *junctionBox) bool {
	_, ok := p.connections[other.id]
	return ok
}

func (p *junctionBox) connect(other *junctionBox) {
	p.connections[other.id] = other
	other.connections[p.id] = p
}

func connectClosestPair(pairs []pair, circuits circuitMap) pair {
	for _, pair := range pairs {
		boxA := pair.a
		boxB := pair.b

		if !boxA.isConnected(boxB) {
			boxA.connect(boxB)
			circuits.merge(boxA.circuit, boxB.circuit)
			if len(circuits) == 1 {
				return pair
			}
		}
	}

	return pairs[len(pairs)-1]
}

type pair struct {
	a *junctionBox
	b *junctionBox
	dist float64
}

func makeCartesianPairs(boxes []*junctionBox) []pair {
	var pairs []pair

	for i := range boxes {
		for j := range boxes {
			if i == j { continue }
			boxA := boxes[i]
			boxB := boxes[j]
			dist := boxA.distance(boxB)
			pairs = append(pairs, pair{boxA, boxB, dist})
		}
	}

	return pairs
}

func connectJunctionBoxes(boxes []*junctionBox, circuits circuitMap, times int) int {
	pairs := makeCartesianPairs(boxes)
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].dist < pairs[b].dist
	})

	pair := connectClosestPair(pairs, circuits)
	return int(pair.a.x * pair.b.x)
}

func parseBoxes(input string) ([]*junctionBox, circuitMap) {
	var boxes []*junctionBox
	circuits := make(circuitMap)

	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		z, _ := strconv.ParseFloat(parts[2], 64)
		box := &junctionBox{id: i, x:x,y:y,z:z, circuit: i, connections: make(map[int]*junctionBox)}
		boxes = append(boxes, box)
		circuits[i] = []*junctionBox{box}
	}

	return boxes, circuits
}

func Day8() {
	input := loadInput("day8.txt")
	boxes, circuits := parseBoxes(input)
	fmt.Println("Day 8", connectJunctionBoxes(boxes, circuits, 1000))
}
