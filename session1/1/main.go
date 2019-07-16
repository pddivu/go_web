package main

import "fmt"

// composite literal
//Slice

func main() {

	x := []int{7, 9, 10}
	fmt.Println(x)

	for i := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 1, 100)
	y = append(y, 7)
	y = append(y, 100)
	fmt.Println(y)
	fmt.Println(y[1])

	for j, v := range y { //---- j is the key and v is the value
		fmt.Println(j, "-", v)
	}
}

/*
output
[7 9 10] -----> x
[0 7 100] -----> y

*/
