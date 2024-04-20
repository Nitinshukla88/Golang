package main

import "fmt"

func main() {

	var value int = 45
	var ptr *int = &value
	fmt.Printf("%T is the datatype of ptr and value at ptr is %d", ptr, *ptr)

}
