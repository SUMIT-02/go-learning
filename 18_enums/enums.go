package main

import "fmt"

type OrderStatus string

// const (
// 	recived OrderStatus = iota
// 	changed
// 	deliverd
// 	prepared
// )

const (
	Recived  OrderStatus = "recived"
	Changed              = "changed"
	Deliverd             = "deliverd"
	Prepared             = "prepared"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("order recived", status)

}

func main() {
	changeOrderStatus("recived")
}
