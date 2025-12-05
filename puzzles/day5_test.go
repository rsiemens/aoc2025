package puzzles

import "testing"


func TestPrepareFreshRanges(t *testing.T) {
	test := `3-5
10-14
16-20
12-18`

	got := prepareFreshRanges(test)
	want := [][2]uint64{{3, 5}, {10, 14}, {12, 18}, {16, 20}}

	if len(got) != len(want) {
		t.Errorf("Expected %v got %v", want, got)
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("Expected %v got %v", v, got[i])
		}
	}
}

func TestCountFresh(t *testing.T) {
	test := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	if got, want := countFresh(test), uint64(14); got != want {
		t.Errorf("Expected %d got %d", want, got)
	}
}
