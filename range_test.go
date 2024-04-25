package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, v := range nums {
		sum += v
	}

	fmt.Println(sum)

	for _, v := range nums {
		if v%2 == 0 {
			fmt.Println(v)
		}
	}

	fruits := map[string]string{"a": "Apple", "b": "Banana"}

	// key value in map
	for k, v := range fruits {
		fmt.Printf("%s - %s\n", k, v)
	}

	// key in map
	for k := range fruits {
		fmt.Println(k)
	}

	// iterate string with range
	for i, c := range "ari" {
		fmt.Printf("index %d = %s\n", i, string(c))
	}

}
