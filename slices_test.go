package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSlices(t *testing.T) {
	var names []string

	names = append(names, "ari")
	names = append(names, "joko")
	names = append(names, "dimas")

	fmt.Println(len(names))
	fmt.Println(cap(names))
	fmt.Println(names)

	animals := make([]string, 3)

	animals[0] = "Zebra"
	animals[1] = "Giraffe"

	fmt.Println(len(animals))
	fmt.Println(cap(animals))
	fmt.Println(animals)

	// copy slices
	animalsCopy := make([]string, len(animals))
	copy(animalsCopy, animals)

	fmt.Println(len(animals))
	fmt.Println(cap(animals))
	fmt.Println(animals)

	l := names[1:]
	h := names[2:4]
	j := names[:3]

	fmt.Println(l)
	fmt.Println(h)
	fmt.Println(j)

	s1 := []int{1, 2, 3}

	s2 := []int{1, 2, 3}
	if slices.Equal(s1, s2) {
		fmt.Println("s1 == s2")
	}

}
