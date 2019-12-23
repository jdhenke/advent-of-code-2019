package main

import (
	"testing"
)

func TestBestAsteroid(t *testing.T) {
	for _, tc := range []struct {
		name               string
		input              string
		x, y, numAsteroids int
	}{
		{
			"a",
			`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`,
			5, 8, 33,
		},
		{
			"b",
			`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
`,
			1, 2, 35,
		},
		{
			"c",
			`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
`,
			6, 3, 41,
		},
		{
			"d",
			`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`,
			11, 13, 210,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			x, y, num := findBestAsteroid(tc.input)
			if x != tc.x || y != tc.y {
				t.Errorf("Got (%d, %d), want (%d, %d)", x, y, tc.x, tc.y)
			}
			if got, want := num, tc.numAsteroids; got != want {
				t.Errorf("Got %d asteroids, want %d", got, want)
			}
		})
	}
}

func TestSimplify(t *testing.T) {
	for _, tc := range []struct {
		a, b, sa, sb int
	}{
		{1, 2, 1, 2},
		{2, 4, 1, 2},
		{4, 6, 2, 3},
	} {
		if gotA, gotB := simplify(tc.a, tc.b); gotA != tc.sa || gotB != tc.sb {
			t.Errorf("(%d,%d) got (%d,%d), want (%d,%d)", tc.a, tc.b, gotA, gotB, tc.sa, tc.sb)
		}
	}
}

func TestGetNumAsteroidsInSight(t *testing.T) {
	for _, tc := range []struct {
		name       string
		input      string
		x, y, want int
	}{
		{
			"a",
			`.#..#
.....
#####
....#
...##`, 3, 4, 8,
		},
		{
			"b",
			`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`, 5, 8, 33,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, want := getNumAsteroidsInSight(tc.input, tc.x, tc.y), tc.want; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 221; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestGetVaporizationOrder(t *testing.T) {
	//	getVaporizationOrder(`.#....#####...#..
	//##...##.#####..##
	//##...#...#.#####.
	//..#.....#...###..
	//..#.#.....#....##`, 8, 3 )
	order := getVaporizationOrder(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`, 11, 13)
	for _, entry := range []struct {
		i   int
		loc int
	}{
		{1, 1112},
		{2, 1201},
		{3, 1202},
		{10, 1208},
		{20, 1600},
		{50, 1609},
		{100, 1016},
		{199, 906},
		{200, 802},
		{201, 1009},
		{299, 1101},
	} {
		if got, want := order[entry.i-1], entry.loc; got != want {
			t.Errorf("%d got %d, want %d", entry.i, got, want)
		}
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 806; got != want {
		t.Errorf("PartB got %d, want %d", got, want)
	}
}
