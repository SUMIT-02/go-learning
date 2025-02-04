package main

import "fmt"

//func printslice[T any](items []T) {
// func printslice[T int | string | bool](items []T) {
 func printslice[T comparable](items []T) {


// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	printslice(nums)
// }

//stacks are lifo last in first out

type stack[T any] struct {
	elements []
}

func main() {
	myStack := stack(string){
		elements: []int{1, 2, 3, 3},
	}
	fmt.Println(myStack)
}
