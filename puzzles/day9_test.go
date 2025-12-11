package puzzles

import "testing"




func TestInBounds(t *testing.T) {
	var test = `1,1
	1,5
	3,5
	3,3
	5,3
	5,5
	7,5
	7,1`
	var coords = parseCoords(test)
	a := &coord{x:7, y:1}
	b := &coord{x:3, y:3}
	if got := inBounds(a, b, coords); got != true {
		t.Errorf("Expected %v got %v", true, got)
	}

	a = &coord{x:7, y:1}
	b = &coord{x:1, y:5}
	if got := inBounds(a, b, coords); got != false {
		t.Errorf("Expected %v got %v", false, got)
	}
}

func TestFindMaxArea(t *testing.T) {
	var test = `1,1
	1,5
	3,5
	3,3
	5,3
	5,5
	7,5
	7,1`
	var coords = parseCoords(test)
	got := findMaxArea(coords)
	if got != 15 {
		t.Errorf("Expected %d got %d", 15, got)
	}
}
