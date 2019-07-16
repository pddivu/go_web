package main

import "fmt"

type gator int
type flemingo bool

func (g gator) greeting() {
	fmt.Println("I am a gator")
}

func (f flemingo) greeting() {
	fmt.Println("I am pink n beautiful")

}

type swamp interface {
	greeting()
}

func buyo(s swamp) {
	s.greeting()
}

func main() {

	var g1 gator
	g1.greeting()

	var f1 flemingo
	f1.greeting()
	buyo(f1)
}
