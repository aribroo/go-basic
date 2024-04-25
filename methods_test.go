package main

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {

	c := circle{r: 5}
	fmt.Println(c.area())

	r := rectangle{width: 10, height: 5}
	fmt.Println(r.area())

}

type circle struct {
	r float64
}

func (c *circle) area() float64 {
	return 3.14 * c.r * c.r
}

type rectangle struct {
	width  float64
	height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}
