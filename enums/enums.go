package main

import "fmt"

//enums - enumaerated types, implemented using "type" in go

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received OrderStatus = "received"
	Confirmed			 = "confirmed"
	Prepared			 = "prepared"
	Delivered  			 = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
func main() {
	changeOrderStatus(Received)
}