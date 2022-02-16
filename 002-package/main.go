package main

import "fmt"

func printSomething() {
	n, _ := fmt.Println("Greeting for the day", 0730, false)
	fmt.Println(n)
}

func main() {
	printSomething()
}
