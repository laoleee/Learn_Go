// 一个常规的Go版socket server模板看起来就长这样

package go_tcpsock

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:5000'")

	if err != nil {
		fmt.Println("Server listen error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Server accept error: ", err)
			break
		}

		go handleConn(conn) // 开Goroutine处理来自客户端的socket

	}
}

func handleConn(conn net.Conn) {

	defer conn.Close()

	for {
		// 这里就可以从连接的socket里读写数据
		// ...
	}
}
