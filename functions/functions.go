package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a , b int) int {
// 	return a + b
// }

// func getManyReturns() (string, string, bool){
// 	return "Hanu", "Arpu", true
// }

// func processFn(fn func(a int)int)int {
// 	return fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	// fmt.Println(add(3,4))
	// fmt.Println(getManyReturns())
	// lang1, lang2, lang3 := getManyReturns()
	// fmt.Println(lang1, lang2, lang3)

	// fn := func (a int)int  {
	// 	return 2
	// }
	// fmt.Println(processFn(fn))

	// function returning another function 
	// val := processIt()
	// fmt.Println(val(1))
}