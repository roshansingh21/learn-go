package main

import "fmt"

func main() {
	var favSport string = "footballl"
	switch favSport {
	case "cricket":
		fmt.Println("His fav sport is Cricket")
		
	case "football":
		fmt.Println("His fav sport is football")
		
	case "Basketball":
		fmt.Println("His fav sport is basketball")
		
	default:
		fmt.Println("His fav sport is not listed")
	}
}
