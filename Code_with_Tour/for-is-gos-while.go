package main

import "fmt"


func main() {
	sum := 1
	for sum<100{
		sum += sum
	}

	// 死循环 for {}
	fmt.Println(sum)
}