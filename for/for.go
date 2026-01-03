package main

import "fmt"

func main() {
	//there is no while loop in Go
	//while loop in Go implemented by for loop
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//standard for loop
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	//using range
	for i:= range(3) {
		fmt.Println(i)
	}
}