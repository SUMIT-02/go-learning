package main

import "fmt"

func main() {
	//Simple Switvh

	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("other")

	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("WeekDay")

	// }

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI(true)

}
