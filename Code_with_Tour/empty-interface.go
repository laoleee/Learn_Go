package main

import "fmt"

func main() {

	// 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
	var i interface{}
	describe(i) // (nil, nil)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
