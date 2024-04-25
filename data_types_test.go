package main

import (
	"fmt"
	"testing"
)

func TestDataTypes(t *testing.T) {
	// string
	var x string = "Hello"
	y := "World"

	// int
	var number int = 10

	fmt.Println(x + y)
	fmt.Println(number)

	// float
	result := 1.1 + 2.2
	fmt.Println(result)

	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
