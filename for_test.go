package main

import (
	"fmt"
	"testing"
)

func TestLooping(t *testing.T) {

	i := 0

	for i < 5 {
		fmt.Print(i)
		i++
	}

	fmt.Println()

	for n := 0; n < 5; n++ {
		fmt.Print(n)
	}

	fmt.Println()

	for {
		fmt.Println("looping")
		break
	}

	for v := range 5 {
		fmt.Println(v)
	}

}
