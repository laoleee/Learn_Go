// 可以为指针接收者声明方法
// 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 为结构体声明方法
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 为结构体声明方法 结果输出5
// func (v Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// 为(结构体)指针声明方法 结果输出50
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 以指针为接收者的方法被调用时，接收者既能为值又能为指针
// 以值为接收者的方法被调用时，接收者既能为值又能为指针
func main() {
	v := Vertex{3, 4}
	v.Scale(10) // Go 会将语句 v.Scale(10) 解释为 (&v).Scale(10)
	p := &v
	fmt.Println(p.Abs()) // Go会将语句p.Abs() 解释为 (*p).Abs()
}

// 使用指针接收者的原因有二：

// 首先，方法能够修改其接收者指向的值。

// 其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
