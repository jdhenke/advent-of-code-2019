package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestQuine(t *testing.T) {
	output := make(chan int, 1)
	const code = "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	go run(code, nil, output)
	var vals []string
	for val := range output {
		vals = append(vals, fmt.Sprint(val))
	}
	if got, want := strings.Join(vals, ","), code; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func Test16DigitNumber(t *testing.T) {
	output := make(chan int, 1)
	const code = "1102,34915192,34915192,7,4,7,99,0"
	go run(code, nil, output)
	val := <-output
	if got, want := len(fmt.Sprint(val)) == 16, true; got != want {
		t.Errorf("%d was not a 16 digit number", val)
	}
	fmt.Println(<-output)
}

func TestLargeNumber(t *testing.T) {
	output := make(chan int, 1)
	const code = "104,1125899906842624,99"
	go run(code, nil, output)
	if got, want := <-output, 1125899906842624; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPartA(t *testing.T) {
	if got, want := PartA(), 2932210790; got != want {
		t.Errorf("PartA got %d, want %d", got, want)
	}
}

func TestPartB(t *testing.T) {
	if got, want := PartB(), 73144; got != want {
		t.Errorf("PartB got %d, want %d", got, want)
	}
}
