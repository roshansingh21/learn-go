package main

import "fmt"

func main() {
	year := 1994

	for {
		if year > 2022 {
			break
		}
		fmt.Println(year)
		year++
	}
}
