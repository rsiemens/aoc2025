package puzzles

import "testing"

func TestMaxJoltage(t *testing.T) {
	tests := []struct{
		input string
		expect uint64
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, test := range tests {
		if got := maxJoltage(test.input); got != test.expect {
			t.Errorf("Expected %d got %d", test.expect, got)
		}
	}
}

func TestTotalJoltage(t *testing.T) {
	test := `987654321111111
811111111111119
234234234234278
818181911112111`
	if got, expect := totalJoltage(test), uint64(3121910778619); got != expect {
		t.Errorf("Expected %d got %d", expect, got)
	}
}
