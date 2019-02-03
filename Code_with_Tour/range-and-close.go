/*
*  只有发送者可通过close关闭一个channel来表示没有需要
*  发送的值了（接收者不能），接收者可以通过为接收表达式分配第二个参数
*  来测试chanel 是否被关闭，没有值接收且chanel已被关闭
*  那么在执行完 v, ok := <- ch后ok被设置为false
*
*  信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再
*  有需要发送的值时才有必要关闭，例如终止一个 range 循环。
*
*  循环 for i := range c 会不断从信道接收值，直到它被关闭
 */

package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
