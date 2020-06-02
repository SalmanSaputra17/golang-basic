package main

import(
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type Square struct {
	l, w float64
}

func (s *Square) Area() float64 {
	return s.l * s.w
}

func totalArea(shapes ...Shape) float64 {
	var total float64

	for _, row := range shapes {
		total += row.Area()
	}

	return total
}

func main() {
	c := Circle{10, 10, 15}
	fmt.Println(c.Area())

	s := Square{14, 25}
	fmt.Println(s.Area())

	fmt.Println(totalArea(&c, &s))
}