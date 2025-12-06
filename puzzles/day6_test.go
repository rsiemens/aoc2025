package puzzles

import "testing"

func TestFindSplitPoints(t *testing.T) {
	test := []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}
	want := []int{3, 7, 11}
	got := findSplitPoints(test)

	if len(got) != len(want) {
		t.Errorf("Expected %v got %v", want, got)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected %v got %v", want, got)
		}
	}
}

func TestPrepareHomework(t *testing.T) {
	test := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	want := [][]string{
		{"123", "328", " 51", "64 "},
		{" 45", "64 ", "387", "23 "},
		{"  6", "98 ", "215", "314"},
		{"*  ", "+  ", "*  ", "+  "},
	}
	got := prepareHomework(test)

	if len(got) != len(want) {
		t.Errorf("Expected %v got %v", want, got)
	}

	for y, row := range want {
		for x := range row {
			if want[y][x] != got[y][x] {
				t.Errorf("Expected %v got %v", want[y][x], got[y][x])
			}
		}
	}
}

func TestEvaluateHomework(t *testing.T) {
	test := [][]string{
		{"123", "328", " 51", "64 "},
		{" 45", "64 ", "387", "23 "},
		{"  6", "98 ", "215", "314"},
		{"*  ", "+  ", "*  ", "+  "},
	}
	if got, want := evaluateHomework(test), 3263827; got != want {
		t.Errorf("Expected %d got %d", want, got)
	}
}
