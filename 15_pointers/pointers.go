package main

import "fmt"

//by value
// func changenum(num int) {
// 	num = 5
// 	fmt.Println("Inchangenum", num)
// }

// by refrence
func changeNum(num *int) {
	*num = 5
	fmt.Println("Inchange", *num)
}

func main() {
	num := 1
	//changenum(num)

	changeNum(&num)

	fmt.Println("Memoryaddress", &num)

	fmt.Println("AfterChange", num)
}
