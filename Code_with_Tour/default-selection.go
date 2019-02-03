package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // 为了在尝试发送或者接收时不发生阻塞，可使用 default 分支
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
