package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	int_num0 := 1
	var float_num0 float32 = 1.2
	var complex_num0 complex128 = 1 + 1.2i

	fmt.Println(int_num0, float_num0, complex_num0)

	var int_num1 int = int_num0 + 1
	// float_num1 := int_num0 + 2.1 无法隐式转换
	float_num1 := float_num0 + 2.1
	fmt.Println(int_num1, float_num1, cmplx.Conj(complex_num0))
}
