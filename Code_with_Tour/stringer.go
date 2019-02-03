// fmt 包中定义的 Stringer 是最普遍的接口之一。
/*
type interface Stringer{
	String() string
}
*/

package main

import "fmt"

// Person 定义一个人结构体
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"laole", 23}
	b := Person{"rui", 22}
	fmt.Println(a, b)
}
