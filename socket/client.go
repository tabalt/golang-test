package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	// 从命令行中读取第二个参数作为名字，如果不存在第二个参数则报错退出
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s name ", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	// 连接到服务端的8080端口
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkError(err)

	for {
		// 循环往连接中 写入名字
		_, err = conn.Write([]byte(name))
		checkError(err)

		// 循环从连接中 读取响应内容，没有响应时会阻塞
		response := make([]byte, 256)
		readLength, err := conn.Read(response)
		checkError(err)

		// 将读取响应内容输出到控制台，并sleep一秒
		if readLength > 0 {
			fmt.Println("[client] server response:", string(response))
			time.Sleep(1 * time.Second)
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("fatal error: " + err.Error())
	}
}
