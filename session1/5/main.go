package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretagent struct {
	person
	righttokill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says, 'Good morning'")
}

func (sa secretagent) speak() {
	fmt.Println(sa.fname, sa.lname, "You are Done")
}

type human interface {
	speak()
}

func info(h human) {
	h.speak()
}

func vomit(h human) {
	fmt.Printf("%T\n", h)
	fmt.Println(h)
	switch v := h.(type) {
	case person:
		fmt.Println(v.fname)
	case secretagent:
		fmt.Println(v.fname)
	default:
		fmt.Println("unknown")
	}
}

func main() {

	p1 := person{"Divya", "Divakaran"}
	sa1 := secretagent{person{"James", "Bond"}, true}

	p1.speak()
	sa1.speak()
	sa1.person.speak()

	info(p1)
	info(sa1)
	vomit(p1)
}
