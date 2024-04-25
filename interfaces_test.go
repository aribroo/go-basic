package main

import (
	"fmt"
	"testing"
)

func TestInterfaces(t *testing.T) {

	c := cat{}
	d := dog{}

	fmt.Println(makeAnimalSpeak(&c))
	fmt.Println(makeAnimalSpeak(&d))

}

type animal interface {
	speak() string
}

type cat struct{}

func (a *cat) speak() string {
	return "meoww"
}

type dog struct{}

func (a *dog) speak() string {
	return "gug gug gug"
}

func makeAnimalSpeak(a animal) string {
	return a.speak()
}
