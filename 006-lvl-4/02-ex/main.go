package main

import "fmt"

func main(){
	x:=[]int{1,2,3,0,9,8,4,5,6,8,7}

	for _,v:=range x{
		fmt.Println(v)
	}

	fmt.Printf("%T",x)
}