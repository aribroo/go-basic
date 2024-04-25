package main

import (
	"fmt"
	"testing"
)

func TestVariadicFunction(t *testing.T) {

	r := sumNumbers(1, 2, 3, 4, 5)
	fmt.Println(r)

	x := []int{5, 1, 5, 3, 9}
	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...)
	y := sumNumbers(x...)
	fmt.Println(y)

}

func sumNumbers(nums ...int) int {
	fmt.Println(nums)
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}
