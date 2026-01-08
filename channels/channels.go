package main

import (
	"fmt"
)

//sending data
// func processNum( numChan chan int ) {

// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// 	// fmt.Println("processing number", <-numChan)
// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func task(done chan bool) {
// 	defer func ()  {
// 		done <- true
// 	} ()

// 	fmt.Println("processing...")
// }

// func emailSender (emailChan <-chan string, done chan bool) { // <-chan means its receive only channel, its used to type-safety , same for chan<- bool
// 	defer func ()  {
// 		done <- true
// 	} ()

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// channels are used for communication between goroutines
func main() {

	// time.Sleep(time.Second * 2)

	// messageChan := make(chan string)

	// messageChan <- "ping" //blocking

	// msg := <-messageChan
	// fmt.Println(msg)
	//---------------------------------

	// numChan := make(chan int)

	// go processNum(numChan)

	// // numChan <- 5
	// for{
	// 	numChan <- rand.Intn(100)
	// }
	//-------------------------------------

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result	//blocking
	// fmt.Println(res)
	//-------------------------------------

	// done := make(chan bool)
	// go task(done)

	// <- done //blocking
	//-------------------------------------

	//buffered channels-> we can send an amount of data without blocking

	// emailChan := make(chan string, 100) //buffered channel
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i:=0; i<5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending emails")
	// close(emailChan) //important to close channel
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	//-------------------------------------

	// listening to multiple channels at once
	// chan1 := make(chan int)
	// chan2 := make(chan string)

	// go func() {
	// 	chan1 <- 10
	// }()

	// go func() {
	// 	chan2 <- "pong"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1Val := <-chan1:
	// 		fmt.Println("received data from chan1", chan1Val)

	// 	case chan2Val := <-chan2:
	// 		fmt.Println("received data from chan2", chan2Val)
	// 	}
	// }
}
