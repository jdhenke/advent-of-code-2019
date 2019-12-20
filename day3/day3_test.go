package main

import "testing"

func TestClosestIntersection(t *testing.T) {
	for _, tc := range []struct {
		name string
		wire1, wire2 string
		distance int
	}{
		{
			"walk through",
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			6,
		},
		{
			"first example",
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			159,
		},
		{
			"second example",
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			135,
		},
	}{
		t.Run(tc.name, func(t *testing.T) {
			if got, want := closestIntersection(tc.wire1, tc.wire2), tc.distance; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 731; got != want {
		t.Errorf("PatA got %d, want %d", got, want)
	}
}