package main

import "fmt"

func main() {
	x := [5]int{}
	x[0] = 1
	x[1] = 11
	x[2] = 111
	x[3] = 1111
	x[4] = 11111

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(v, " is at index", i)
	}

	fmt.Printf("%T", x)
}
