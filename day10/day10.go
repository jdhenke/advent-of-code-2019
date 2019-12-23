package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const realInput = `###..#########.#####.
.####.#####..####.#.#
.###.#.#.#####.##..##
##.####.#.###########
###...#.####.#.#.####
#.##..###.########...
#.#######.##.#######.
.#..#.#..###...####.#
#######.##.##.###..##
#.#......#....#.#.#..
######.###.#.#.##...#
####.#...#.#######.#.
.######.#####.#######
##.##.##.#####.##.#.#
###.#######..##.#....
###.##.##..##.#####.#
##.########.#.#.#####
.##....##..###.#...#.
#..#.####.######..###
..#.####.############
..##...###..#########`

func main() {
	fmt.Println(PartA())
	fmt.Println(PartB())
}

func PartA() int {
	_, _, num := findBestAsteroid(realInput)
	return num
}

func PartB() int {
	x, y, _ := findBestAsteroid(realInput)
	return getVaporizationOrder(realInput, x, y)[200-1]
}

func findBestAsteroid(input string) (bestX, bestY, numAsteroids int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	cols := len(lines[0])
	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			if lines[y][x] != '#' {
				continue
			}
			if a := getNumAsteroidsInSight(input, x, y); a > numAsteroids {
				bestX, bestY, numAsteroids = x, y, a
			}
		}
	}
	return
}

func getNumAsteroidsInSight(input string, baseX, baseY int) int {
	grid := toGrid(input)
	return len(getAsteroidsInSight(grid, baseX, baseY))
}

func getAsteroidsInSight(grid [][]bool, baseX, baseY int) []int {
	//fmt.Println(grid)
	var asteroidsInSight []int
	for y := range grid {
		for x := range grid[y] {
			if !grid[y][x] {
				continue
			}
			if x == baseX && y == baseY {
				continue
			}

			dx, dy := simplify(abs(x-baseX), abs(y-baseY))
			//fmt.Println(x, y,  abs(x - baseX), abs(y-baseY), dx, dy)
			if x < baseX {
				dx = -dx
			}
			if y < baseY {
				dy = -dy
			}
			inSight := true
			for lx, ly := baseX+dx, baseY+dy; !(lx == x && ly == y); lx, ly = lx+dx, ly+dy {
				//fmt.Println(lx, ly)
				if grid[ly][lx] {
					//fmt.Printf("(%d,%d) blocks (%d, %d)\n", lx, ly, x, y)
					inSight = false
					break
				}
			}
			if inSight {
				//fmt.Printf("In sight (%d, %d)\n", x, y)
				asteroidsInSight = append(asteroidsInSight, x*100+y)
			}
		}
	}
	return asteroidsInSight
}

func toGrid(input string) [][]bool {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var grid [][]bool
	for _, line := range lines {
		var gridRow []bool
		for _, c := range line {
			val := false
			switch c {
			case '#':
				val = true
			case '.':
			default:
				panic(c)
			}
			gridRow = append(gridRow, val)
		}
		grid = append(grid, gridRow)
	}
	return grid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func simplify(a, b int) (f1, f2 int) {
	if a == 0 {
		return 0, 1
	}
	if b == 0 {
		return 1, 0
	}
	for f := 2; f <= a; f++ {
		if a%f == 0 && b%f == 0 {
			return simplify(a/f, b/f)
		}
	}
	return a, b
}

func getVaporizationOrder(input string, baseX, baseY int) []int {
	grid := toGrid(input)
	var output []int
	for inSight := getAsteroidsInSight(grid, baseX, baseY); len(inSight) > 0; inSight = getAsteroidsInSight(grid, baseX, baseY) {
		sort.Slice(inSight, func(i, j int) bool {
			xi, yi := inSight[i]/100-baseX, inSight[i]%100-baseY
			xj, yj := inSight[j]/100-baseX, inSight[j]%100-baseY
			return getTheta(xi, yi) < getTheta(xj, yj)
		})
		output = append(output, inSight...)
		for _, loc := range inSight {
			//xi, yi := loc/100-baseX, loc%100-baseY
			//fmt.Println(loc, xi, yi, getTheta(xi, yi))
			grid[loc%100][loc/100] = false
		}
	}
	return output
}

func getTheta(xi, yi int) float64 {
	fxi := float64(xi)
	fyi := float64(yi)
	return math.Atan2(-fxi, fyi)
}
