package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//or we can use interface
// func printSlice[T interface](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//to allow only int or string 
// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//to make it more generic, we can use comparable keyword
// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//we can have more parameters also
// func printSlice[T comparable, V string](items []T, name V) {
// 	for _, item := range items {
// 		fmt.Println(item, name)
// 	}
// }


// we can also make a generic stack
// type stack struct {
// 	elements[]int
// }

//to make it generic
type stack[T any] struct {
	elements[]T
}

func main() {
	// nums := []int{1,2,3}

	//if i have to print another slice containing strings i cant use the same function as we used to print integers, even though the logic is same; to solve this problem, to create a generic function we use generics
	// printSlice(nums)

	//we created a slice to use the function on any datatype
	// nums := []int{1,2,3}
	// printSlice(nums)
	// names := []string {"happy", "world"}
	// printSlice(names)

	//using stack
	// myStack := stack{
	// 	elements: []int{1,2,3},
	// }
	// fmt.Println(myStack)

	//using generalized stack
	myStack := stack[int]{
		elements: []int{1,2,3},
	}
	fmt.Println(myStack)

	myStack2 := stack[string]{
		elements: []string{"hello","helu","hi"},
	}
	fmt.Println(myStack2)

}