package puzzles

import "testing"


func TestRemoveRolls(t *testing.T) {
	test := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	if got, want := removeRolls(test), 43; got != want {
		t.Errorf("Expected %d got %d", want, got)
	}
}
