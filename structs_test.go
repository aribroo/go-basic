package main

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {

	budi := person{name: "Budi", age: 21}
	ari := newPerson("ari", 21)

	fmt.Println(budi)
	fmt.Println(*ari)

	s := &budi
	fmt.Println(s.name)

	cat := struct {
		name   string
		isGood bool
	}{"Momo", true}

	fmt.Println(cat)
}

type person struct {
	name string
	age  uint8
}

func newPerson(name string, age uint8) *person {
	return &person{name, age}
}
