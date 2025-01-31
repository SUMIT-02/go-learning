package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func main() {
	myOrder := order{
		id:     "1",
		amount: 322.00,
		status: "acceped",
	}

	myOrder.createdAt = time.Now()
	fmt.Println(myOrder.status)

	myOrder2 := order{
		id:     "2",
		amount: 23.43,
		status: "rejected",
	}

	fmt.Println("order status", myOrder)
	fmt.Println("order status", myOrder2)

}
