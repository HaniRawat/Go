package main

import (
	"fmt"
	"time"
)

//go-routines are light threads which are used for multithreading

// func task(id int) {
// 	fmt.Println("doing task with id", id)
// }

func main() {
	for i:=0; i <= 10; i++ {
		// go task(i)

		go func(i int) {
			fmt.Println(i)
		} (i)
	}


	time.Sleep(time.Second * 2)	//delaying the main function
}