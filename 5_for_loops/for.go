package main

import "fmt"

// for -> only construct in go for looping
func main(){

	// i :=1
	
	// while loop style

	// for i<=3{
	// 	fmt.Println(i)
	// 	i++
	// }


	// infinite loop
	// for{
	// 	println(i)
	// 	i++
	// }

	// classic for loop

	// for i:=0; i<3; i++ {
	// 	fmt.Println(i)
	// }


	// range based for loop
	for i := range 10 {
		fmt.Println(i)
	}
}