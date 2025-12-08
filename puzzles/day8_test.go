package puzzles

import "testing"


func TestConnectJunctionBoxes(t *testing.T) {
	input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	boxes, circuits := parseBoxes(input)
	got := connectJunctionBoxes(boxes, circuits, 10)

	if got != 25272 {
		t.Errorf("Expected %d got %d", 25272, got)
	}
}

func TestParseBoxes(t *testing.T) {
	test := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
	boxes, circuits := parseBoxes(test)

	if len(boxes) != 20 {
		t.Errorf("Expected %d got %d", 20, len(boxes))
	}

	if len(circuits) != 20 {
		t.Errorf("Expected %d got %d", 20, len(circuits))
	}
}
