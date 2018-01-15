package main

import "fmt"

func main() {
	var a = average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(a)
}

func average(x ...float64) float64 {
	var total float64
	for _, v := range x {
		total += v
	}
	total = total / float64(len(x))
	return total
}
