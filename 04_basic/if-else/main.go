package main

import "fmt"

func main() {
	/*Normal*/
	if true {
		fmt.Println("TRUE")
	}
	if false {
		fmt.Println("FALSE")
	}

	/*initialize statement*/
	b := true
	if name := "Mg Mg"; b {
		fmt.Println(name)
	}
}
