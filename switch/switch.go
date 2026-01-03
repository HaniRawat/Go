package main

import (
	"fmt"
)

func main() {

	// switch time.Now().Weekday() {
	// case time.Sunday , time.Saturday: //multiple condition
	// 	fmt.Println("ITS WEEKEND")

	// default:
	// 	fmt.Println("ITS SHIT")
	// }

	whoAmI := func (i interface{})  {
		switch i.(type) {
		case int :
			fmt.Println("its an integer")
		case string :
			fmt.Println("its a string")
		case bool : 
			fmt.Println("its a boolean value")
		default :
			fmt.Println("other")
		}
	}
	whoAmI(27)
}