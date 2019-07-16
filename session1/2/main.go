package main

import "fmt"

func main() {

	//map literal

	m := map[string]int{ // MAp decalration 1
		"Divya":  35,
		"Eshaan": 7,
	}
	m2 := make(map[string]string) // mpa decalration 2
	m2["Divya"] = "PD"
	fmt.Println(m2["Divya"])

	for i, v := range m {
		fmt.Println(i, "-", v) // v or m[i]
	}

	fmt.Printf("%T\n", m)
	fmt.Println(m["Divya"])
}
