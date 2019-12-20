package main

import "testing"

func TestNumOrbits(t *testing.T) {
	if got, want := numOrbits([]string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}), 42; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 621125; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestShortestPath(t *testing.T) {
	if got, want := shortestPathYouToSanta([]string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}), 4; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 550; got != want {
		t.Errorf("PartB got %d, want %d", got, want)
	}
}
