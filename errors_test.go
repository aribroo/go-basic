package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {

	for i := range 5 {
		if res, err := getTea(i); err != nil {
			if errors.Is(err, errOutOfTea) {
				fmt.Println(err.Error())
			} else {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(res)
		}
	}

}

var errOutOfTea = errors.New("tea not available")
var ErrPower = errors.New("can't boil water")

func getTea(n int) (string, error) {
	if n == 2 {
		return "", errOutOfTea
	} else if n == 4 {
		return "", ErrPower
	} else {
		return "tea is ready to drink", nil
	}
}
