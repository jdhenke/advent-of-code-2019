package main

import "fmt"

func main() {
	fmt.Println(PartA())
}

func PartA() int {
	numMet := 0
	for i := 402328; i <= 864247; i++ {
		if meetCriteria(i) {
			numMet++
		}
	}
	return numMet
}

func meetCriteria(x int) bool {
	digits := []int{
		(x/100000) % 10,
		(x/10000) % 10,
		(x/1000) % 10,
		(x/100) % 10,
		(x/10) % 10,
		(x/1) % 10,
	}
	hasDouble := false
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			hasDouble = true
		}
		if digits[i] > digits[i+1] {
			return false
		}
	}
	return hasDouble
}