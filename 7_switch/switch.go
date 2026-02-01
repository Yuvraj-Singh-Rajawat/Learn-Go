package main

import (
	"fmt"
)

func main() {
	// simple switch

	// i := 2

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2: 
	// 	fmt.Println("two")
	// default:
	// 	fmt.Print("other")
	// }



	// multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 		fmt.Println("It's weekend")
	// default: 
	// 	fmt.Println("It's a weekday")
	// }


	whoAmI := func(i interface{}){
		switch t:= i.(type){
		case int: 
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown", t)
		}

	}


	whoAmI("goland")
}