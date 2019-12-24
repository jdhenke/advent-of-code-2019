package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const realInput = `<x=-3, y=10, z=-1>
<x=-12, y=-10, z=-5>
<x=-9, y=0, z=10>
<x=7, y=-5, z=-3>`

var re = regexp.MustCompile(`<x=(-?[0-9]+), y=(-?[0-9]+), z=(-?[0-9]+)>`)

func main() {
	fmt.Println(PartA())
}

type planet struct {
	x, y, z    int
	vx, vy, vz int
}

func (p planet) Energy() int {
	return (abs(p.x) + abs(p.y) + abs(p.z)) * (abs(p.vx) + abs(p.vy) + abs(p.vz))
}

func PartA() int {
	return getEnergy(realInput, 1000)
}

func getEnergy(input string, steps int) int {
	planets := simulate(input, steps)
	energy := 0
	for _, planet := range planets {
		energy += planet.Energy()
	}
	return energy
}

func simulate(input string, steps int) []planet {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var planets []planet
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			panic(line)
		}
		x, y, z := mustInt(matches[1]), mustInt(matches[2]), mustInt(matches[3])
		planets = append(planets, planet{x: x, y: y, z: z})
	}
	for i := 0; i < steps; i++ {
		// apply gravity
		var newPlanets []planet
		for _, p := range planets {
			newP := p
			for _, p2 := range planets {
				newP.vx += getVelocityChange(newP.x, p2.x)
				newP.vy += getVelocityChange(newP.y, p2.y)
				newP.vz += getVelocityChange(newP.z, p2.z)
			}
			newPlanets = append(newPlanets, newP)
		}
		// apply velocity
		for i := range newPlanets {
			newPlanets[i].x += newPlanets[i].vx
			newPlanets[i].y += newPlanets[i].vy
			newPlanets[i].z += newPlanets[i].vz
		}

		planets = newPlanets
	}
	return planets
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getVelocityChange(xMe, xOther int) int {
	if xMe == xOther {
		return 0
	}
	if xOther > xMe {
		return 1
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
