package main

import "testing"

func TestPartA(t *testing.T) {
	if got, want := PartA(), 2720; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), `   XX XXXX XXX    XX XXX   XX   XX    XX   
    X    X X  X    X X  X X  X X  X    X   
    X   X  X  X    X X  X X  X X       X   
    X  X   XXX     X XXX  XXXX X XX    X   
 X  X X    X    X  X X X  X  X X  X X  X   
  XX  XXXX X     XX  X  X X  X  XXX  XX    `; got != want {
		t.Errorf("PartB got:\n%s, want:\n%s", got, want)
	}
}
