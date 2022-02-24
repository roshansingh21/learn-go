package main

import "fmt"

func main() {
	x := "Bond"

	if x == "Bond" {
		fmt.Println("Its James Bond")
	} else if x == "Spidy" {
		fmt.Println("Its Spiderman")
	} else {
		fmt.Println("Its regular human")
	}
}
