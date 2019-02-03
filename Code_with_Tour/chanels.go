/*
chanel是带有类型的管道，可以通过chanel操作符<-
来发送或接收值：
ch  <- v  // 将v发送至信道ch
v := <- ch  // 将ch接收值并赋予v
箭头就是数据流的方向

创建chanel:
ch := make(chan int)

默认情况下，发送和接收操作在另一端准备好之前都会阻塞
这使得Go程可以在没有显式的锁或竞态变量的情况下进行同步
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v

	}

	c <- sum
}

func main() {
	s := []int{7, 2, 7, 8, 8, 1}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c种接收数据
	fmt.Println(x, y, x+y)
}
