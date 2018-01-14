package main

import "fmt"

func main() {
	/*For basic*/
	// a := 0
	// for i := 0; i < 1000000; i++ {
	// 	a = a + i
	// }
	// fmt.Println(a)

	/*For Like while*/
	// b := 0
	// for b < 100 {
	// 	fmt.Println(b)
	// 	b++
	// }

	//Dont use this//never use
	/*For with no condition*/
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	/*For Break*/
	// j := 0
	// for {
	// 	if j < 10 {
	// 		break
	// 	}
	// 	j++
	// }

	/*For Continue*/
	k := 0
	for {
		k++
		if k > 100 {
			continue
		}
		fmt.Println(k)
	}
}
