package main

import (
	"fmt"
	"testing"
)

func TestConditionalStatement(t *testing.T) {

	number := 10

	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if n := 1; n < 10 {
		fmt.Println("Less than 10")
	} else if n > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Equal 10")
	}

}
