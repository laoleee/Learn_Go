// Go的基本数据类型

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 的别名

// rune // int32 的别名
//      // 代表一个Unicode码

// float32 float64

// complex64 complex128

// 变量在定义时没有明确的初始化时会赋值为_零值_。

// 零值是：

// 数值类型为 `0`，
// 布尔类型为 `false`，
// 字符串为 `""`（空字符串）。

package main

import (
	"fmt"
	"math/cmplx"
) // 同时与导入语句一样，变量的定义“打包”在一个语法块中

var (
	a int16      = 16
	b bool       = true
	c complex128 = 3 - 5i
	d complex64  // 0—-0i
	f float32    // 0
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(cmplx.Sqrt(c))
	fmt.Println(d, f)

}
