package main

import "fmt"

func main() {
	// var name string = "Golang" -> explicit type declaration

	// var name = "Golang" // -> implicit type declaration

	var age int = 100



	// shorthand syntax
	// name := "golang"

	var name string
	name = `goland${age}`



	

	fmt.Println(name)
	fmt.Println(age)
}