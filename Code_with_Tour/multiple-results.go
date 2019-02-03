package main

import "fmt"

func main(){
	a, b := swap("hello", "laole")
	fmt.Println(a, b)
}


func swap(x, y string) (string, string){
	return y, x  /// 多返回值
}