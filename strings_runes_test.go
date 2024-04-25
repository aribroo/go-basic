package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestStringsAndRunes(t *testing.T) {
	s := "hello"

	fmt.Printf("Len : %d\n", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count = ", utf8.RuneCountInString(s))
}
