package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 监听8080端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// 循环接收客户端的连接，没有连接时会阻塞，出错则跳出循环
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("[server] accept new connection.")

		// 启动一个goroutine 处理连接
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	for {
		// 循环从连接中 读取请求内容，没有请求时会阻塞，出错则跳出循环
		request := make([]byte, 128)
		readLength, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if readLength == 0 {
			fmt.Println(err)
			break
		}

		// 控制台输出读取到的请求内容，并在请求内容前加上hello和时间后向客户端输出
		fmt.Println("[server] request from ", string(request))
		conn.Write([]byte("hello " + string(request) + ", time: " + time.Now().Format("2006-01-02 15:04:05")))
	}
}
