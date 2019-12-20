package main

import (
	"testing"
)

func TestGetFuel(t *testing.T) {
	for _, tc := range []struct{
		mass, fuel int
	}{
		{12, 2},
		{14, 2},
		{ 1969, 654},
		{100756, 33583},
	}{
		if got, want := getFuel(tc.mass), tc.fuel; got != want {
			t.Errorf("For mass %d got %d, want %d", tc.mass, got, want)
		}
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 3325156; got != want {
		t.Fail()
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 4984866; got != want {
		t.Fatalf("PartB got %d, want %d", got, want)
	}
}