package main

import "fmt"

func main() {
	/*Shorthand*/
	a := 1.1
	b := "MgMg"
	c := true
	fmt.Println(a, b, c)
	fmt.Printf("%T \n", a, b, c)

	/*Zero Value*/
	var d string
	var e int
	var f bool
	fmt.Println(d, e, f)
	fmt.Printf("%T \n", d, e, f)
	fmt.Println()

	/*Var*/
	var message string
	message = "My Name is Mg Mg"
	fmt.Println(message)

	/*int many at once time*/
	var g, h, i, j, k int = 2, 3, 4, 5, 6
	fmt.Println(g, h, i, j, k)

	/*mix at once time*/
	var l, n, m = true, "Aung Aung", 3.4
	fmt.Println(l, n, m)

	/*single code,double code,back tick*/
	o := "o"
	/*Please dont use Now(Next time fix)*/
	//p := 'a'
	q := `q`
	fmt.Println(o, q)

}
