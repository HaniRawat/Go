package main

import "fmt"

type paymenter interface{
	//we cant use concrete implementation of any payment gateway
	pay(amount float32)	//we dont have to tell which interface are we implementing as GO matches the function signatures
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	//if we want to add more payment gateways like stripe
	// this breaks Open Close Principle of SOLID
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

func (p payment) makeRefund(amount float32, account string) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	//if we want to add more payment gateways like stripe
	// this breaks Open Close Principle of SOLID
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.refund(amount, account)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

//if we want to add more payment gateways like stripe

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println(("making payment using fake gateway for testing purposes..."))
}

type payPal struct {}

func (p payPal) pay(amount float32) {
	fmt.Println("making payment using PayPal", amount)
}

func (p payPal) refund(amount float32, account string) {
	fmt.Println("refund of", amount ,"done for account ", account, "by Paypal")
}

func main() {
	// razorpayPaymentGw := razorpay{}

	// newPayment := payment {
	// 	gateway : razorpayPaymentGw,
	// }
	// newPayment.makePayment(100)

	// fakeGw := fakePayment{}
	// newPayment := payment {
	// 	gateway: fakeGw,
	// }l

	payPalGw := payPal{}

	newPayment := payment{
		gateway : payPalGw,
	}

	newPayment.makePayment(100)
	newPayment.makeRefund(100, "xxxxxx156")
}