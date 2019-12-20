package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	for _, x := range []int{
		1,
		42,
		-2,
		0,
	} {
		if got, want := run("3,0,4,0,99", x), x; got != want {
			t.Errorf("Identity program got %d, want %d", got, want)
		}
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 16225258; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestEqualTo8(t *testing.T) {
	for _, codeTC := range []struct {
		name string
		code string
	}{
		{"position", "3,9,8,9,10,9,4,9,99,-1,8"},
		{"immediate", "3,3,1108,-1,8,3,4,3,99"},
	} {
		for _, tc := range []struct {
			input, output int
		}{
			{
				8,
				1,
			},
			{
				7,
				0,
			},
			{
				9,
				0,
			},
			{
				0,
				0,
			},
			{
				-1,
				0,
			},
		} {
			if got, want := run(codeTC.code, tc.input), tc.output; got != want {
				t.Errorf("%s mode with input %d got %d, want %d", codeTC.name, tc.input, got, want)
			}
		}
	}
}

func TestLessThan8(t *testing.T) {
	for _, codeTC := range []struct {
		name string
		code string
	}{
		{"position", "3,9,7,9,10,9,4,9,99,-1,8"},
		{"immediate", "3,3,1107,-1,8,3,4,3,99"},
	} {
		for _, tc := range []struct {
			input, output int
		}{
			{
				8,
				0,
			},
			{
				7,
				1,
			},
			{
				9,
				0,
			},
			{
				0,
				1,
			},
			{
				-1,
				1,
			},
		} {
			if got, want := run(codeTC.code, tc.input), tc.output; got != want {
				t.Errorf("%s mode with input %d got %d, want %d", codeTC.name, tc.input, got, want)
			}
		}
	}
}

func TestJump(t *testing.T) {
	for _, codeTC := range []struct {
		name string
		code string
	}{
		{"position", "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"},
		{"immediate", "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"},
	} {
		for _, x := range []int{0, 1, 2, -1, -42, 99} {
			want := 0
			if x != 0 {
				want = 1
			}
			t.Run(fmt.Sprintf("%s mode input %d", codeTC.name, x), func(t *testing.T) {
				if got := run(codeTC.code, x); got != want {
					t.Errorf("got %d want %d", got, want)
				}
			})
		}
	}
}

func TestBig8(t *testing.T) {
	big8 := `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`
	for _, tc := range []struct {
		input, output int
	}{
		{-2, 999},
		{0, 999},
		{7, 999},
		{8, 1000},
		{9, 1001},
		{42, 1001},
	} {
		if got, want := run(big8, tc.input), tc.output; got != want {
			t.Errorf("input %d got %d, want %d", tc.input, got, want)
		}
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 2808771; got != want {
		t.Errorf("PartB got %d, want %d", got, want)
	}
}
