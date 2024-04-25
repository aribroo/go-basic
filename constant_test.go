package main

import (
	"fmt"
	"testing"
)

const s string = "constant"

func TestConstant(t *testing.T) {
	fmt.Println(s)

	const PI float32 = 3.14

	const (
		x = 10
		y = 20
	)

	fmt.Println(PI)
	fmt.Println(PI * x)
	fmt.Println(PI * y)

}
