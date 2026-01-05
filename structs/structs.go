package main

import (
	"fmt"
	"time"
)

//order struct
// type order struct {
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time //nanoseconds precision
// }



//constructor
// func newOrder(id string, amount float32, status string) *order {
// 	//initial setup goes here
// 	myOrder := order {
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// //attaching a function to order struct
// func (o *order)changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) getAmount() float32 {
// 	return o.amount;
// }


// --------------------------------------------------------
//struct embedding - kinda like inheritance
type customer struct {
	name string
	phone string
}
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time //nanoseconds precision
	customer 
}


func main() {

	//if you dont initialize any value ,it is set to 0 value by default
	//zero values int -> 0, string -> "", bool -> false

	// var order order = order {
		// id : "1",
		// amount : 50.00,
		// status: "received",
	// }

	// myOrder := order{
	// 	id : "1",
	// 	amount : 50.00,
	// 	status: "received",
	// }

	// myOrder2 := order{
	// 	id : "2",
	// 	amount : 100.00,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.createdAt = time.Now()
	// fmt.Println("order struct -", myOrder)
	// fmt.Println("order status -", myOrder.status)
	// fmt.Println("order struct -", myOrder2)

	// myOrder := order{
	// 	id : "1",
	// 	amount : 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println("status changed in order struct -", myOrder)

	// // amount := myOrder.amount
	// fmt.Println("amount is -", myOrder.getAmount())

	//using constructor function
	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println( myOrder)
	// fmt.Println( myOrder.amount)

	//making a single use struct
	// language := struct {
	// 	name string
	// 	isGood bool
	// } {"golang", true}

	// fmt.Println(language)

	// ------------------struct embedding---------
	// newCustomer := customer {
	// 	name: "John",
	// 	phone: "0123456789",
	// }
	// newOrder := order {
	// 	id: "1",
	// 	amount: 30,
	// 	status: "received",
	// 	customer: newCustomer,
	// }

	newOrder := order {
		id: "1",
		amount: 30,
		status: "received",
		customer: customer{
			name: "John",
			phone: "0123456789",
		},
	}
	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)

}