package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favfood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname, "is walking")
}

func main() {
	p1 := person{"Divya",
		"pd",
		[]string{"orange", "apple", "candy"}}

	//a := p1.walk()
	fmt.Println(p1.walk())
}
