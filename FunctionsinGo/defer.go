package main

import "fmt"

//defer keyword is used to delay the function execution of any function.

func printname(str string){
	fmt.Println(str)
}
func printage(age int){
	fmt.Println(age)
}
func printaddress(add string){
	fmt.Println(add)
}
func main(){
	defer printname("Nitin")
	printage(22)
	printaddress("Kanpur")
}