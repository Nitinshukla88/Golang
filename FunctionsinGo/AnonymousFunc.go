package main

import "fmt"

func main() {
	x := func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Println(x(2, 3))
	fmt.Printf("The data-type of x is %T", x)
}
