package main

import (
	"fmt"
)

func main() {
	p := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("p[1:4] = ", p[1:4])
	fmt.Println("p[1:1] = ", p[1:1]) // 空的
	fmt.Println("p[1:] = ", p[1:])   // 1到结束
	fmt.Println("p[:4] = ", p[:4])   // 开始至4
	fmt.Println("p[1:2] = ", p[1:2]) // 1个元素

}
