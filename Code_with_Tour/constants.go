// 常量的定义与变量类似，只不过使用 const 关键字。

// 常量可以是字符、字符串、布尔或数字类型的值。

// 常量不能使用 := 语法定义。

package main

import "fmt"

func main(){
	const world = "laole"
	const trueth = true

	fmt.Println(world, trueth)
}