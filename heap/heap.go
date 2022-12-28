package main

import "fmt"

var res []int
var arr []int = []int{1, 3, 5, 7, 2}

func main() {
	for idx, val := range arr {
		parent := idx / 2
		ch1 := 2 * idx
		ch2 := 2*idx + 1
		fmt.Printf("Val: %d, Parent: %d, Ch1: %d, Ch2: %d\n", val, parent, ch1, ch2)
	}
}
