package main

import(
	"fmt"
	"math"
)

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

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hello, my name is", p.Name)
}

type Android struct {
	// embedded type
	Person
	Name string
}

func main() {
	fmt.Println("Learn how to use struct in Go.")

	// set zero value to struct and return a pointer
	// c := new(Circle)
	// c.x = 10
	// c.y = 10
	// c.r = 5

	// c := Circle{x: 5, y: 5, r: 2}

	c := Circle{5, 5, 5}

	fmt.Println(c.Area())

	s := Square{94.10, 201.4}

	fmt.Println(s.Area())

	p := new(Person)
	p.Name = "salman"
	p.Talk()

	// a := new(Android)

	// a.Person.Talk()
	// a.Talk()
}