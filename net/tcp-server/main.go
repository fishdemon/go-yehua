package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Starting the server ...")

	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:7575")

	if err != nil {
		fmt.Println("Error starting the server", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting from the client", err.Error())
			return
		}

		go handle(conn)
	}

}

func handle(conn net.Conn)  {
	// 设定超时时间
	conn.SetDeadline(time.Now().Add(2 * time.Minute))
	fmt.Println("新连接: ", conn.RemoteAddr())
	for {
		buf := make([]byte, 521)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
		fmt.Println()

		conn.Write([]byte("i have received"))
		fmt.Println("回复成功")
	}

}