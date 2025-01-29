package main

import (
	"fmt"
)

//slices are dynamic arrays
//no length required automaticllly expand

func main() {
	//uninitilized slices are nil
	//var nums []int
	//fmt.Println(nums == nil)
	//fmt.Println(len(nums))

	//capacity
	//fmt.Println(cap(nums))
	//fmt.Println(nums == nil)
	//nums := []int{}
	//nums = append(nums, 1)
	//fmt.Println(nums)
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums1 = make([]int, len(nums))

	// //nums[0] = 3
	// copy(nums1, nums)
	// fmt.Println(nums, nums1)

	//var nums = []int{1, 2, 3}
	//var nums1 = []int{1, 2, 5}

	//fmt.Println(nums[0:1])

	// fmt.Println(nums[:1])
	// fmt.Println(nums[1:])
	//slices
	//fmt.Println(slices.Equal(nums, nums1))

	var nums = [][]int{{1, 2, 3, 4}, {1, 2, 3}}
	fmt.Println(nums)
}
