package puzzles

import "testing"

var test = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

var machines = parseMachines(test)


func TestRandomDistanceStrategy(t *testing.T) {
	if got := randomDistanceStrategy(machines, 20); got != 33 {
		t.Errorf("Expected %d got %d", 33, got)
	}
}
