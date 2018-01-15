package main

import "fmt"

func main() {
	// var x [100]int64
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[90])
	// x[90] = 30
	// fmt.Println(x[90])

	var y [256]string
	fmt.Println(len(y))
	fmt.Println(y[0])
	for i := 0; i < 256; i++ {
		y[i] = string(i)
	}
	for _, v := range y {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}
