package main

import "fmt"

type Curency int

const ( // 指定一个索引和对应值列表进行初始化数组
	USD Curency = iota // 美元
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int                // 定义一个int型的数组 0
	var b [3]int = [3]int{1, 3} // 定义且初始化, b[2]为0
	c := [...]int{3, 4, 5}      // 或者这样定义和初始化

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1]) // len()函数返`回数组中元素个数

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range b {
		fmt.Printf("%d\n", v)
	}

	for _, v := range c {
		fmt.Printf("%d\n", v)
	}

}
