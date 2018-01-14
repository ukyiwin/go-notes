package main

import "fmt"

func main() {
	// name := "Aung"
	// switch name {
	// case "Mg":
	// 	fmt.Println("Name is Mg Mg")
	// case "Aung":
	// 	fmt.Println("Name is Aung")
	// }

	//fallthrough
	/*
	  fallthrough is that next condition is true without true;
	*/
	// a := 1
	// switch a {
	// case 1:
	// 	println("1")
	// 	fallthrough
	// case 2:
	// 	println("2")
	// 	fallthrough
	// case 3:
	// 	println("3")
	// 	fallthrough
	// case 4:
	// 	println("4")
	// case 5:
	// 	println("5")
	// }

	//multiple-evels(true on line)
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}

}
