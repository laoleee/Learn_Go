package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	i := &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	// 接口值i保存了一个具体底层类型的具体值(具体看哪个实现调用)
	fmt.Printf("(%v, %T)\n", i, i)
}
