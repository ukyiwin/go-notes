package main

import "fmt"

//need to import other package (packagetest)
import "github.com/golangmy/go-notes/02_Package/packagetest"

func main() {
	packagetest.Hello()
	fmt.Println(packagetest.V1)
}
