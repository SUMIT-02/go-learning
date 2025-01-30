package main

import "fmt"

func sum(nums ...int) int {
	//any value use interface
	//func sum(nums ...interface) int {

	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {
	nums := []int{1, 2, 3, 4}
	result := sum(nums...)
	fmt.Println(result)
}
