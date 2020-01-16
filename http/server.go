package main

import (
	"log"
	"net/http"
	"os"
)

// 定义http请求的处理方法
func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http hello on golang\n"))
}

func main() {

	// 注册http请求的处理方法
	http.HandleFunc("/hello", handlerHello)

	// 在8086端口启动http服务，会一直阻塞执行
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}

	// http服务因故停止后 才会输出如下内容
	log.Println("Server on 8080 stopped")
	os.Exit(0)
}
