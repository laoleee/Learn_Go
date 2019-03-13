package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	t1 := time.Now().Nanosecond()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Now().Nanosecond() - t1)

	// 第二种方法
	t2 := time.Now().Nanosecond()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now().Nanosecond() - t2)

}
