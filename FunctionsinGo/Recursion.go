package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
func main() {
	factorialofnum := factorial(5)
	fmt.Println(factorialofnum)
}
