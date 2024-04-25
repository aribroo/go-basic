package main

import (
	"fmt"
	"testing"
)

func TestClosures(t *testing.T) {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())

}

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}

}
