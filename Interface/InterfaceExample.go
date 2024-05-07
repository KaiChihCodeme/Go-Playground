package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

type Square struct {
	length float64
}

func (r Square) Area() float64 {
	return r.length * r.length
}

func PrintArea(s Shape) {
	// type 1
	fmt.Println(s.Area())
}

func PrintWho(s Shape) {
	// type 2
	switch s.(type) {
	case Circle:
		fmt.Println("I am circle")
	case Shape:
		fmt.Println("I am Sqaure")
	}
}

func main() {
	circle := Circle {
		radius: 5,
	}

	square := Square {
		length: 3,
	}

	PrintArea(circle)
	PrintArea(square)

	PrintWho(circle)
	
}

