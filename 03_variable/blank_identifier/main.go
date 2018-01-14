package main

import "fmt"

func main() {
	//Golang cannot declare without using........
	/*
		var a = 12
		var b = 24

		fmt.Println(a)
	*/

	//you can use that..........
	_ = 2
	var c = 3
	fmt.Println(c)
}
