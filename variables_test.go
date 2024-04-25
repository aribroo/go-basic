package main

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {
	var a string = "Apple"
	var b = "Orange"
	c := "Melon"

	name, age := "ari", 21

	var (
		car   = "Avanza"
		brand = "Toyota"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(car)
	fmt.Println(brand)
}
