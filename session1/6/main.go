package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	for i, v := range []byte(s) {
		fmt.Println(i, v)
	}

	fmt.Println([]byte(s)[2])
	fmt.Println(string([]byte(s)[2]))
	fmt.Println(string([]byte(s)[2:]))

	d := " I am sorry i cant do that"
	fmt.Println(d)
	fmt.Println([]byte(d))
	fmt.Println(string([]byte(d)))
	fmt.Println(d[:11])
	fmt.Println(d[11:])
	for _, v := range d {
		fmt.Println(string(v))
	}
}
