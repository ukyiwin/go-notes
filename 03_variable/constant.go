package main

import "fmt"

func main() {
	const mgmg = "This is Mg Mg"

	//cannot assign mgmg bz mgmg is const
	mgmg = "This is Aung Aung"

	fmt.Println(mgmg)
}
