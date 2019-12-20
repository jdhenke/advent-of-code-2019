package main

import "fmt"

func main() {
	fmt.Println(PartA())
	fmt.Println(PartB())
}

func PartA() int {
	numMet := 0
	for i := 402328; i <= 864247; i++ {
		if meetCriteriaPartA(i) {
			numMet++
		}
	}
	return numMet
}

func meetCriteriaPartA(x int) bool {
	digits := []int{
		(x / 100000) % 10,
		(x / 10000) % 10,
		(x / 1000) % 10,
		(x / 100) % 10,
		(x / 10) % 10,
		(x / 1) % 10,
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

func PartB() int {
	numMet := 0
	for i := 402328; i <= 864247; i++ {
		if meetCriteriaPartB(i) {
			numMet++
		}
	}
	return numMet
}

func meetCriteriaPartB(x int) bool {
	digits := []int{
		(x / 100000) % 10,
		(x / 10000) % 10,
		(x / 1000) % 10,
		(x / 100) % 10,
		(x / 10) % 10,
		(x / 1) % 10,
	}
	hasExactDouble := false
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] &&
			(i == 0 || digits[i-1] != digits[i]) &&
			(i+2 >= len(digits) || digits[i] != digits[i+2]) {
			hasExactDouble = true
		}
		if digits[i] > digits[i+1] {
			return false
		}
	}
	return hasExactDouble
}
