package main

import "testing"

func TestPartA(t *testing.T) {
	if got, want := PartA(), 4090701; got != want {
		t.Errorf("PartA got %d want %d", got, want)
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 6421; got != want {
		t.Errorf("PartB got %d want %d", got, want)
	}
}