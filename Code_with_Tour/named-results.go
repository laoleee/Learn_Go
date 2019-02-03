package main

import "fmt"

func main(){
	a, b:=split(10)
	fmt.Println(a, b)
}


func split(sum int) (x, y int){  // Go 的返回值可以被命名，并且像变量那样使用
	x = sum -10
	y = sum -1
	return
}