package main

import "fmt"

func sum(nums ...int) int {	//only accepts integer arguments
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

// func sum(nums ...interface{}) { //accepts any type of arguments passed into it, this gives error in this program because of addition
// 	total := 0

// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

func main() {
	// fmt.Println(1,2,3,4,5, "hello") //Println() is a variadic function as we can pass any number of arguments in it

	// result := sum(1,2,3,4,5,6,7,8,9)
	// result := sum(1,2,3,4,5,6,7,8,9, 1.1)
	// fmt.Println(result)

	nums := []int {1,2,3,4,5}

	result := sum(nums...)
	fmt.Println(result)

}