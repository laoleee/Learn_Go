package main

import "fmt"

func main() {

	// 字符串不允许使用单引号
	var s1 string = "Single quotes work well for string literals."
	var s2 string = "Double quotes work just as well."
	s3 := "It's easy to escape the string delimiter."
	s4 := "It's even easier to use the other delimiter."

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
