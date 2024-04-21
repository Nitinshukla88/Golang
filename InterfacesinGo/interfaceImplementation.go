package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side int
}

type rect struct {
	length int
	breath int
}

func (s square) area() float64 {
	return float64(s.side) * float64(s.side)
}

func (s square) perimeter() float64 {
	return float64(s.side) * 4
}

func (r rect) area() float64 {
	return float64(r.length) * float64(r.breath)
}

func (r rect) perimeter() float64 {
	return float64(r.length)*2 + float64(r.breath)*2
}

func printdata(sh shape) {
	fmt.Println(sh)
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())
}

func main() {
	shape1 := square{
		side: 4,
	}
	shape2 := rect{
		length: 3,
		breath: 5,
	}
	printdata(shape1)
	printdata(shape2)
}
