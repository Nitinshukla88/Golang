package main

import "fmt"

// Returning single value from function

// func addnumbers(num1 int, num2 int) int {
// 	sum := num1 + num2
// 	return sum
// }

// func main() {
// 	sumofnumbers := addnumbers(3, 4)
// 	fmt.Println(sumofnumbers)
// }

// Returning multiple values from function

// func multi_returner(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2
// 	diff := num1 - num2
// 	return sum, diff
// }
// func main() {
// 	var num1 int = 20
// 	var num2 int = 10
// 	sum, difference := multi_returner(num1, num2)
// 	fmt.Printf("Sum of %v and %v is %v along with difference %v", num1, num2, sum, difference)
// }

// Variadic functions in Golang - Those which accepts variable number of arguments

// func variadicingo(numbers ...int) int {
// 	sum := 0
// 	for _, element := range numbers {
// 		sum += element
// 	}
// 	return sum
// }

// func main() {
// 	sum := variadicingo(2, 3, 4, 5, 6)
// 	fmt.Println(sum)
// 	fmt.Println(variadicingo(23, 45))
// 	fmt.Println(variadicingo(3))
// }

// Suppose we want to print sum of marks of any student using variadic function

func variadicingo(student string, numbers ...int) {
	sumofmarks := 0
	for _, value := range numbers {
		sumofmarks += value
	}
	fmt.Printf("%s has scored %v marks", student, sumofmarks)
}

// func main() {
// 	variadicingo("Nitin", 43, 5, 2, 3, 5)
// }
