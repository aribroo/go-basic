package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitchCase(t *testing.T) {
	n := 1

	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")

	}

	checkType := func(v any) string {
		switch t := v.(type) {
		case bool:
			return "boolean"
		case int:
			return "integer"
		case string:
			return "string"
		default:
			return fmt.Sprintf("unknown type: %T", t)
		}
	}

	fmt.Println(checkType(true))
	fmt.Println(checkType(10))
	fmt.Println(checkType("Hello"))
	fmt.Println(checkType(10.1))

}
