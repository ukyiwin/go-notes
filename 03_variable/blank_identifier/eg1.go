package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//with error checking
	// res, err := http.Get("http://www.geekwiseacademy.com/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// page, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", page)

	//You can use that...
	//without error checking
	res, _ := http.Get("https://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

}
