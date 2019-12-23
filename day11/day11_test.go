package main

import "testing"

func TestPartA(t *testing.T) {
	if got, want := PartA(), 2720; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}
