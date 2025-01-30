package main

import "fmt"

//func add(a, b int) int {
//	return a + b
//}

func getLanguages() (string, string, bool) {
	return "java", "cpp", false
}
func main() {

	//result := add(5, 5)
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
}
