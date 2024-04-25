package main

import (
	"fmt"
	"testing"
)

func TestArrays(t *testing.T) {
	var names [3]string

	names[0] = "ari"
	names[1] = "dimas"
	names[2] = "wahyu"

	fmt.Println(len(names))
	fmt.Println(names)

	numbers := [...]int{1, 2, 3, 4, 5}

	fmt.Println(len(numbers))
	fmt.Println(numbers)

}
