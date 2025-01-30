// package main

// import "fmt"

//	func main() {
//		nums := []int{6, 7, 8}
//		for i := 0; i < len(nums); i++ {
//			fmt.Println(nums[i])
//		}
//	}
package main

import "fmt"

func main() {
	//nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	//sum := 0
	// for i, num := range nums {
	// 	//	sum = sum + num
	// 	fmt.Println(num, i)
	// }
	// //fmt.Println(sum)

	// m := map[string]string{"fnmae": "john", "lname": "pawar"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }
	//unicode its rune
	// i is starting valur of rune
	for i, c := range "golang" {
		//fmt.Println(i, c)
		//to convert binary code in string we use
		fmt.Println(i, string(c))
	}

}
