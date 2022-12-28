package main

import (
	"fmt"
)

func Fibonacci(n int) int {
	if n < 2 {
		return n
	} else {
		return Fibonacci(n-2) + Fibonacci(n-1)
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(Fibonacci(n))
}
