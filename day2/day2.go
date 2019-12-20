package main

import "fmt"

var defaultNums = []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,10,23,2,10,23,27,1,27,6,31,1,13,31,35,1,13,35,39,1,39,10,43,2,43,13,47,1,47,9,51,2,51,13,55,1,5,55,59,2,59,9,63,1,13,63,67,2,13,67,71,1,71,5,75,2,75,13,79,1,79,6,83,1,83,5,87,2,87,6,91,1,5,91,95,1,95,13,99,2,99,6,103,1,5,103,107,1,107,9,111,2,6,111,115,1,5,115,119,1,119,2,123,1,6,123,0,99,2,14,0,0}

func main() {
	fmt.Println(PartA())
	fmt.Println(PartB())
}

func PartA() int {
	return run(12, 2)
}

func PartB() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if run(noun, verb) == 19690720 {
				return noun * 100 + verb
			}
		}
	}
	panic("No valid inputs found")
}

func run(noun, verb int) int {
	nums := make([]int, len(defaultNums))
	copy(nums, defaultNums)
	nums[1], nums[2] = noun, verb
	for i := 0; ; i+=4 {
		switch nums[i] {
		case 1:
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		case 2:
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		case 99:
			return nums[0]
		default:
			panic(fmt.Sprintf("Unknown op at position %d: %d", i, nums[i]))
		}
	}
}