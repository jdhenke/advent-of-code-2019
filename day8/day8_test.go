package main

import "testing"

func TestPartA(t *testing.T) {
	if got, want := PartA(), 1905; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestDecode(t *testing.T) {
	if got, want := decode("0222112222120000", 2, 2), "01\n10\n"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), `0110001100100101110011110
1001010010101001001000010
1001010000110001001000100
1111010000101001110001000
1001010010101001000010000
1001001100100101000011110
`; got != want {
		t.Errorf("PartB got %s, wanted %s", got, want)
	}
}
