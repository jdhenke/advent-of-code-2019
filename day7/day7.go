package main

import (
	"fmt"
	"strconv"
	"strings"
)

const realCode = `3,8,1001,8,10,8,105,1,0,0,21,34,51,68,89,98,179,260,341,422,99999,3,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,1002,9,2,9,4,9,99,3,9,1001,9,3,9,102,3,9,9,101,4,9,9,4,9,99,3,9,102,2,9,9,101,2,9,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,99`

func main() {
	fmt.Println(PartA())
}

func PartA() int {
	_, output := findMaxAmp(realCode)
	return output
}

func findMaxAmp(code string) (bestCombo []int, bestOutput int) {
	combos([]int{0, 1, 2, 3, 4}, func(ints []int) {
		output := runAmplifiers(code, ints)
		if output > bestOutput {
			bestOutput = output
			bestCombo = make([]int, len(ints))
			copy(bestCombo, ints)
		}
	}, 0)
	return bestCombo, bestOutput
}

func runAmplifiers(code string, settings []int) int {
	output := 0
	for _, setting := range settings {
		input := make(chan int, 2)
		input <- setting
		input <- output
		output = run(code, input)
	}
	return output
}

func combos(vals []int, fn func([]int), col int) {
	if col == len(vals)-1 {
		fn(vals)
		return
	}
	for i := col; i < len(vals); i++ {
		vals[col], vals[i] = vals[i], vals[col]
		combos(vals, fn, col+1)
		vals[col], vals[i] = vals[i], vals[col]
	}
}

func run(code string, input <-chan int) int {
	const (
		opAdd = 1
		opMul = 2
		opInp = 3
		opOut = 4
		opJTr = 5
		opJFa = 6
		opLes = 7
		opEqu = 8
		opEnd = 99
	)
	var ops []int
	for _, opStr := range strings.Split(code, ",") {
		op, err := strconv.Atoi(opStr)
		if err != nil {
			panic(err)
		}
		ops = append(ops, op)
	}
	i := 0
	output := 0
	for {
		op := ops[i]
		getParam := func(p int) int {
			val := ops[i+p]
			if (p == 1 && (op/100)%10 == 1) ||
				(p == 2 && (op/1000)%10 == 1) ||
				(p == 3 && (op/10000)%10 == 1) {
				return val
			}
			return ops[val]
		}
		switch op % 100 {
		case opAdd:
			ops[ops[i+3]] = getParam(1) + getParam(2)
			i += 4
		case opMul:
			ops[ops[i+3]] = getParam(1) * getParam(2)
			i += 4
		case opInp:
			ops[ops[i+1]] = <-input
			i += 2
		case opOut:
			output = getParam(1)
			i += 2
		case opJTr:
			if getParam(1) != 0 {
				i = getParam(2)
			} else {
				i += 3
			}
		case opJFa:
			if getParam(1) == 0 {
				i = getParam(2)
			} else {
				i += 3
			}
		case opLes:
			val := 0
			if getParam(1) < getParam(2) {
				val = 1
			}
			ops[ops[i+3]] = val
			i += 4
		case opEqu:
			val := 0
			if getParam(1) == getParam(2) {
				val = 1
			}
			ops[ops[i+3]] = val
			i += 4
		case opEnd:
			return output
		default:
			panic(fmt.Sprintf("Unknown op at position %d: %d", i, op))
		}
	}
}
