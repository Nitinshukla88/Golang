package main

import "fmt"

func addnumbers(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	sumofnumbers := addnumbers(3, 4)
	fmt.Println(sumofnumbers)
}
