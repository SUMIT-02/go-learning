package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	//gateWay stripe
	gateWay paymenter
}

func (p payment) makePayment(amount float32) {
	//razorPaymentGw := razorPay{}
	//stripePaymentGw := stripePay{}

	//razorPaymentGw.pay(amount)
	//stripePaymentGw.pay(amount)
	p.gateWay.pay(amount)
}

type razorPay struct{}

// func (r razorPay) pay(amount float32) {
// 	fmt.Println("making Payment using razorpay", amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making Payment using stripe", amount)
// }

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}
func main() {
	//razorPaymentGw := razorPay{}
	paypalGw := paypal{}
	newPayment := payment{
		gateWay: paypalGw,
	}
	newPayment.makePayment(100)
}
