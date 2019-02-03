package main

import "fmt"

func main(){
	fmt.Println(add(1,1));
}


func add(x int, y int) int {  // 更常用这种写法 func add(x, y int) int{}
	return x+y
}