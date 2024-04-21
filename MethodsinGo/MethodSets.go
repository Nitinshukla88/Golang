package main

import "fmt"

type student struct {
	name   string
	grades []int
}

func (inst student) displayname() {
	fmt.Println(inst.name)
}

func (inst student) calcpercent() {
	sum := 0
	for _, value := range inst.grades {
		sum += value
	}
	percent := float64(sum*100) / float64(len(inst.grades)*100)
	fmt.Println(percent)
}

func main() {
	obj1 := student{
		name:   "Nitin",
		grades: []int{34, 23, 78},
	}
	obj1.displayname()
	obj1.calcpercent()
}
