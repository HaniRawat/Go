package main

import (
	"fmt"
	"sync"
)

//go-routines are light threads which are used for multithreading

// func task(id int) {
// 	fmt.Println("doing task with id", id)
// }

// func main() {
// 	for i:=0; i <= 10; i++ {
// 		// go task(i)

// 		go func(i int) {
// 			fmt.Println(i)
// 		} (i)
// 	}

// 	time.Sleep(time.Second * 2)	//delaying the main function
// }

//we have a problem inside goroutines, in the above main function we delayed the main by poroviding time, but in real life scenarios we dont know how much time will be needed therefore we use wait-groups to synchronize the goroutines
//wait groups are a mechanism to synchronize goroutines

//only remember - add(), done(), wait()

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task with id", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

}
