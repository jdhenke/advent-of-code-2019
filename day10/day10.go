package main

import (
	"fmt"
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
}

func PartA() int {
	_, _, num := findBestAsteroid(realInput)
	return num
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
			if a := getAsteroidsInSight(input, x, y); a > numAsteroids {
				bestX, bestY, numAsteroids = x, y, a
			}
		}
	}
	return
}

func getAsteroidsInSight(input string, baseX, baseY int) int {
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
	//fmt.Println(grid)
	numInSight := 0
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
				numInSight++
			}
		}
	}
	return numInSight
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
