package main

import "fmt"

/*Iota can be use within the const*/
func main() {
	const (
		_ = iota //assign 0
		a = iota //assign 1
		b = iota //assign 2
		c        //assign 3
		d        //assign 4
	)
	fmt.Println(a, b, c, d)

	////////////////////////
	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB = 1 << (iota * 10) // 1 << (2 * 10)
		GB = 1 << (iota * 10) // 1 << (3 * 10)
		TB = 1 << (iota * 10) // 1 << (4 * 10)
	)

	fmt.Println(KB, MB, GB, TB)
}
