package main

import (
	"fmt"
	"maps"
	"testing"
)

func TestMaps(t *testing.T) {
	m := make(map[string]int)

	m["k1"] = 5
	m["k2"] = 10

	fmt.Printf("Map : %v\n", m)
	fmt.Printf("Len : %v\n", len(m))

	k1 := m["k1"]
	k2 := m["k2"]
	k3 := m["k3"]

	fmt.Println(k1)
	fmt.Println(k2)
	fmt.Println(k3)

	// deletes the element with the specified key
	delete(m, "k1")
	fmt.Println(m)

	// clear deletes all entries, resulting in an empty map
	clear(m)
	fmt.Println(m)

	// check if value exist, return bool
	_, ok := m["k2"]
	fmt.Println("result:", ok)

	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(m1, m2) {
		fmt.Println("m1 == m2")
	}

}
