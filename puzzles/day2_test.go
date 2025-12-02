package puzzles

import (
	"testing"
	"strconv"
)

func TestHasRepeatingPattern(t *testing.T) {
	for i := 95; i <= 115; i++ {
		asStr := strconv.Itoa(i)
		if i == 99 || i == 111 {
			if !hasRepeatingPattern(asStr) {
				t.Errorf("Expected %d to be true", i)
			}
		} else if hasRepeatingPattern(asStr) {
			t.Errorf("Expected %d to be false", i)
		}
	}
}

func TestSumRepeats(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"11-22", 33},
		{"95-115", 210},
		{"998-1012", 2009},
		{"1188511880-1188511890", 1188511885},
		{"222220-222224", 222222},
		{"1698522-1698528", 0},
		{"446443-446449", 446446},
		{"38593856-38593862", 38593859},
		{"565653-565659", 565656},
		{"824824821-824824827", 824824824},
		{"2121212118-2121212124", 2121212121},
	}

	ch := make(chan uint64)
	for _, test := range tests {
		go sumRepeats(test.input, ch)
		if result := <-ch; result != test.expected {
			t.Errorf("Expected %d, got %d (%s)", test.expected, result, test.input)
		}
	}
}

func TestSumAllRepeats(t *testing.T) {
	test := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`
	if result := sumAllRepeats(test); result != 4174379265 {
		t.Errorf("Expected %d, got %d", 4174379265, result)
	}
}
