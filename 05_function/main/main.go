package main

import "fmt"

func main() {
	say("Hi Mg Mg")
	say(sayHello("Nyein", "Chan"))
	say(functionOne("Mg ", "Mg"))
	say(functionOne(functionTwo("Mg", "Mg")))
}

/*param-args*/
func say(speak string) {
	fmt.Println(speak)
}

/*return*/
func sayHello(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

/*return naming*/
//retrun s--> different to readable
func functionOne(x, y string) (s string) {
	s = fmt.Sprint(x, y)
	return
}

/*multiple retrun*/
func functionTwo(x, y string) (string, string) {
	return x, y
}
