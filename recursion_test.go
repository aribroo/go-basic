package main

import (
	"fmt"
	"testing"
)

func TestRecursion(t *testing.T) {

	n := factorial(5)
	fmt.Println(n)

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	x := fib(7)
	fmt.Println(x)

}

func factorial(n int) int {

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)

}
