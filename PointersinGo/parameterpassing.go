package main

import "fmt"

func modify(value int) {
	value += 12
}

func main() {
	num := 19

	// fmt.Println(num)

	// // First we change the value of num using standard approach!

	// num = 23

	// Lets try to change the variable of num using a function and parameter passing method is pass by value

	fmt.Println(num)

	modify(num)

	fmt.Println(num)

}
