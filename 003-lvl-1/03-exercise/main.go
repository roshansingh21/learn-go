package main

import "fmt"

var x int = 42
var y string = "Roshan Singh"
var z bool = false

func main() {
	s := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	v := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	fmt.Println(v)
}
