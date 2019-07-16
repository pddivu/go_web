package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourwheel bool
}

func (s sedan) transportdevice() string {
	return fmt.Sprintln("Luxury car")
}

func (t truck) transportdevice() string {
	return fmt.Sprintln("Luggage truck")
}

type transportation interface {
	transportdevice() string
}

func report(h transportation) {
	fmt.Println(h.transportdevice())
}

func main() {

	s1 := sedan{vehicle{4, "black"}, true}
	fmt.Println(s1)

	t1 := truck{vehicle{5, "blue"}, true}
	fmt.Println(t1.transportdevice())

	report(s1)
}
