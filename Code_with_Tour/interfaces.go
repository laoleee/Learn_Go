// 接口类型 是由一组方法签名定义的集合。

// 接口类型的变量可以保存任何实现了这些方法的值

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat实现了Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex实现了 Abser
	fmt.Println(a.Abs())

	// a = v // 没有实现Abser，v是一个Vertex，而不是*Vertex
	// fmt.Println(a.Abs())
	//  Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser
}
