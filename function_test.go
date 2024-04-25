package main

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {

	result := sum(1, 2, 3)
	fmt.Println(result)

	hello := sayHello("Goku")
	fmt.Println(hello)

}

func sum(a, b, c int) int {
	return a + b + c
}

func sayHello(name string) string {
	return "Hello " + name
}
