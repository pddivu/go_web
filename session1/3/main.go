package main

import "fmt"

type person struct { // struct declaration 1
	fname string
	lname string
}

func main() {
	p1 := person{"Divya", "pd"}
	p2 := person{"Eshaan", "Bhagavad"}

	fmt.Println(p1, p2)

	m := struct { // struct decalration 2
		name string
		age  int
	}{"Divya", 35}
	fmt.Println(m)

	//func declaration

	f := receiver()
	fmt.Println(f)

}

func receiver() int {
	return 5
}
