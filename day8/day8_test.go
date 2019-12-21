package main

import "testing"

func TestPartA(t *testing.T) {
	if got, want := PartA(), 1905; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}
