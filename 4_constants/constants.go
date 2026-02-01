package main

import "fmt"

func main() {
	const name string = "Golang"

	// name = "yuvraj"

	println(name)

	// constant grouping

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

}