package main

import "testing"

func TestClosestIntersections(t *testing.T) {
	for _, tc := range []struct {
		name              string
		wire1, wire2      string
		manhattanDistance, stepsDistance int
	}{
		{
			"walk through",
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			6,
			30,
		},
		{
			"first example",
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			159,
			610,
		},
		{
			"second example",
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			135,
			410,
		},
	}{
		t.Run(tc.name, func(t *testing.T) {
			if got, want := closestIntersectionManhattan(tc.wire1, tc.wire2), tc.manhattanDistance; got != want {
				t.Errorf("manhattan distance got %d, want %d", got, want)
			}
			if got, want := closestIntersectionSteps(tc.wire1, tc.wire2), tc.stepsDistance; got != want {
				t.Errorf("steps distance got %d, want %d", got, want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 731; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 5672; got != want {
		t.Errorf("PartB got %d, want %d", got, want)
	}
}