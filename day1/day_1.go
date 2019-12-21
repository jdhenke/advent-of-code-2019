package main

import "github.com/jdhenke/advent-of-code-2019/input"

import (
	"fmt"
)

func main() {
	fmt.Println(PartA())
	fmt.Println(PartB())
}

func PartA() int {
	totalFuel := 0
	if err := input.ForEachNumInFile("day1.txt", func(i int) {
		totalFuel += getFuel(i)
	}); err != nil {
		panic(err)
	}
	return totalFuel
}

func getFuel(mass int) int {
	return mass/3 - 2
}

func PartB() int {
	totalFuel := 0
	if err := input.ForEachNumInFile("day1.txt", func(i int) {
		for additionalFuel := getFuel(i); additionalFuel > 0; additionalFuel = getFuel(additionalFuel) {
			totalFuel += additionalFuel
		}
	}); err != nil {
		panic(err)
	}
	return totalFuel
}
