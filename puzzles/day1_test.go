package puzzles

import "testing"

func TestGetCode(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	if res := getCode(input, 50); res != 6 {
		t.Error("expected 6 got", res)
	}
}

func TestRotate(t *testing.T) {
	type Test struct {
		rot     int
		dial    int
		newDial int
		clicks  int
	}

	tests := []Test{
		{-68, 50, 82, 1},
		{-30, 82, 52, 0},
		{48, 52, 0, 1},
		{-5, 0, 95, 0},
		{60, 95, 55, 1},
		{60, 95, 55, 1},
		{-55, 55, 0, 1},
		{-1, 0, 99, 0},
		{-99, 99, 0, 1},
		{14, 0, 14, 0},
		{-82, 14, 32, 1},
	}

	for _, test := range tests {
		dial, clicks := rotate(test.rot, test.dial)
		if dial != test.newDial || clicks != test.clicks {
			t.Errorf(
				"expected dial: %d clicks: %d got dial: %d clicks: %d",
				test.newDial,
				test.clicks,
				dial,
				clicks,
			)
		}
	}
}
