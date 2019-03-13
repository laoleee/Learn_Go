package main

import "fmt"

func main() {

	var s []int // 申明

	printSlice(s)

	s = append(s, 0) // 添加一个元素
	printSlice(s)

	s = append(s, 1) // 切片会按需增长
	printSlice(s)

	s = append(s, 2, 3, 4) // 添加多个元素
	printSlice(s)
}

func printSlice(x []int) {
	fmt.Printf("len(x)=%d, cap(x)=%d, v=%v\n", len(x), cap(x), x)
}
