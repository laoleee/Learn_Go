package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// 实现接口
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

// 实现接口
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
