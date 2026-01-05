package main

import "fmt"

func changeNum(num int) {
	num = 5 //this is a copy of the num argument
	fmt.Println("changed num", num);
}

func changeByReference(num *int) {
	*num = 5
	fmt.Println("changed reference", *num)
}
func main() {
	num := 1

	//changeNum(num) //we only pass by value so no change happens to the original num

	// fmt.Println(num);

	//we have to pass the number by reference
	// fmt.Println("memory address of num", &num) 
	changeByReference(&num)
	fmt.Println("after change", num);
}