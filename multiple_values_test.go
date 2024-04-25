package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestMultipleValues(t *testing.T) {

	a, b := getValues()

	fmt.Println(a)
	fmt.Println(b)

	_, d := getValues()

	fmt.Println(d)

	result1, err := checkNumber(10)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result1)
	}

	if result2, err := checkNumber(1); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result2)
	}

}

func getValues() (int, int) {
	return 10, 50
}

func checkNumber(n int) (int, error) {
	if n < 5 {
		return 0, errors.New("number must greater than 5")
	} else {
		return n, nil
	}
}
