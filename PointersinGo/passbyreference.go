package main

import "fmt"

// In Golang Slices and maps are passed into functions by default

func modify(ptr *string) {
	*ptr = "Rajat"
}
func modify1(slice []int) {
	slice[1] = 12
}
func modify2(s map[string]int) {
	s["Abhay"] = 23
}
func main() {

	// str := "Nitin"

	// fmt.Println(str)

	// modify(&str)

	// After modify function call, the value of string str has changed due to the fact that we have passed the address of variable str, which is the example of call by refernce.

	// fmt.Println(str)

	// Taking the example of slice..

	// slice := []int{2, 23, 43, 45}

	// modify1(slice)

	// fmt.Println(slice)

	// Taking the example of maps in Golang..

	sample := make(map[string]int)

	sample["Rajan"] = 11

	sample["Nitin"] = 56

	modify2(sample)

	fmt.Println(sample)

	// Here in all the cases excluding the string change where we passed the address of variable, we have used the concept of pass by reference.
}
