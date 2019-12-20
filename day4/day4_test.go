package main

import "testing"

func TestMeetCriteriaPartA(t *testing.T) {
	for _, tc := range []struct {
		x    int
		pass bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	} {
		if got, want := meetCriteriaPartA(tc.x), tc.pass; got != want {
			t.Errorf("%d got %v, want %v", tc.x, got, want)
		}
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 454; got != want {
		t.Errorf("PartA  got %d, want %d", got, want)
	}
}

func TestMeetCriteriaPartB(t *testing.T) {
	for _, tc := range []struct {
		x    int
		pass bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	} {
		if got, want := meetCriteriaPartB(tc.x), tc.pass; got != want {
			t.Errorf("%d got %v, want %v", tc.x, got, want)
		}
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 288; got != want {
		t.Errorf("PartB  got %d, want %d", got, want)
	}
}
