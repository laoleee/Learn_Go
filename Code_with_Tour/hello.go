package main  // 每个Go程序由包组成

import (
	"fmt"
	"time"
)

func main() {  // 程序入口是包的main()函数
	fmt.Println("Hello World!")

	fmt.Println("The time is", time.Now())
}

