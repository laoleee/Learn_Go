//  与C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换

package main

import "fmt"


func main(){
	var i int = 64  // var i = 64 自动类型推导
	var f float64 = float64(i)
	var k uint = uint(f)

	// 更简洁的写法是这样
	// i:=64  // 自动类型推导
	// f:=float64(i)
	// k:=uint(f)

	fmt.Println(i, f, k)
}