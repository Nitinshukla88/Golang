package main

import "fmt"

type name struct {
	yourname string
}

// n*name as receiver changed the yourname field inside the "inst" instance of structure name

// func (n *name) namechanger() {
// 	n.yourname = "rajat"
// }

// while n name as receiver couldn't because here only instance has passed not the address of it.

func (n name) namechanger() {
	n.yourname = "rajat"
}

func main() {
	inst := name{
		yourname: "Nitin",
	}
	inst.namechanger()
	fmt.Printf("%+v", inst)
}
