// 为非结构体类型申明方法
// 接收者的类型定义和方法声明必须在同一包内；
// 且不能为内建类型声明方法

package package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}


import "fmt"

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {

	f1 := MyFloat(23.5)
	f2 := MyFloat(-12.5)
	fmt.Println("f1:", f1.Abs(), "f2:", f2.Abs())

}
