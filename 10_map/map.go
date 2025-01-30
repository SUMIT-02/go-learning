package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]string)

	// m["name"] = "golang"
	// m["area"] = "backend"
	//fmt.Println(m["name"], m["area"])
	//if key is empty its returns empyt value

	//m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 1000
	// // m1 := make(map[string]bool)
	// // m2 := make(map[string]string)

	// //fmt.Println(m["phone"])
	// // fmt.Println(m1["phone"])
	// // fmt.Println(m2["phone"])
	// // fmt.Println(len(m))
	// //delete(m, "age")
	// clear(m)
	// fmt.Println(m)

	//m := map[string]int{"price": 1000, "age": 30}
	//fmt.Println(m)
	// m := map[string]int{"price": 1000, "age": 30}
	// v, ok := m["price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	m1 := map[string]int{"price": 1000, "age": 30}
	m2 := map[string]int{"price": 1000, "age": 10}
	fmt.Println(maps.Equal(m1, m2))

}
