package main

import "testing"

func TestMeetCriteria(t *testing.T) {
	for _, tc := range []struct {
		x    int
		pass bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	} {
		if got, want := meetCriteria(tc.x), tc.pass; got != want {
			t.Errorf("%d got %v, want %v", tc.x, got, want)
		}
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 454; got != want {
		t.Errorf("PartA  got %d, want %d", got, want)
	}
}
