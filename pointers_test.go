package main

import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {
	// pointers, allowing you to pass references to values and records within your program.

	n := 10

	fmt.Println("initial:", n)

	zeroVal(n)
	fmt.Println("zeroval:", n)

	zeroPtr(&n)
	fmt.Println("zeroptr:", n)

	// address of n
	fmt.Println("pointer:", &n)
}

func zeroVal(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}
