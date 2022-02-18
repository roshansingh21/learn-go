package main

import "fmt"

func main() {
	a := (21 == 22)
	b := (22 <= 24)
	c := (30 >= 21)
	d := (21 != 22)
	e := (21 > 22)
	f := (23 < 30)

	fmt.Println(a, b, c, d, e, f)

}
