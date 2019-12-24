package main

import "testing"

func TestGetEnergy(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		steps  int
		energy int
	}{
		{
			"a",
			`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			10,
			179,
		},
		{
			"b",
			`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
			100,
			1940,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, want := getEnergy(tc.input, tc.steps), tc.energy; got != want {
				t.Errorf("energy got %d, want %d", got, want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 10944; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}
