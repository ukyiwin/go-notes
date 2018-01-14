package packagetest

import "fmt"

/*Can call Outside of the other Package
need to import
need to first letter of the function name to Capitical */
func Hello() {
	fmt.Println("Hello World")
	/*call from same package name*/
	name()
}
