package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRunAmps(t *testing.T) {
	for _, tc := range []struct {
		name     string
		code     string
		settings []int
		output   int
	}{
		{
			"a",
			"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			[]int{4, 3, 2, 1, 0},
			43210,
		},
		{
			"b",
			"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			[]int{0, 1, 2, 3, 4},
			54321,
		},
		{
			"c",
			"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			[]int{1, 0, 4, 3, 2},
			65210,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// test that the correct setting  produces the expected output
			if got, want := runAmplifiers(tc.code, tc.settings), tc.output; got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			// test that the right setting is found
			bestCombo, bestOutput := findMaxAmp(tc.code)
			if got, want := bestCombo, tc.settings; !reflect.DeepEqual(got, want) {
				t.Errorf("best combo got %v, want %v", got, want)
			}
			if got, want := bestOutput, tc.output; got != want {
				t.Errorf("best  output got %d, want %d", got, want)
			}
		})
	}
}

func TestCombos(t *testing.T) {
	combos([]int{1, 2, 4}, func(ints []int) {
		fmt.Println(ints)
	}, 0)
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 13848; got != want {
		t.Errorf("PartA got %d, want  %d", got, want)
	}
}
